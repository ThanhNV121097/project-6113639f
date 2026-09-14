package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/ThanhNV121097/project-6113639f/backend/migrations"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL is required")
	}
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := migrate(ctx, db); err != nil {
		log.Fatal(err)
	}
	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		if err := db.PingContext(r.Context()); err != nil {
			http.Error(w, "database unavailable", http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	http.HandleFunc("/v1/greeting", greetingHandler(db))

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("APP_PORT")
	}
	if port == "" {
		port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func migrate(ctx context.Context, db *sql.DB) error {
	if _, err := db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS schema_migrations (filename text PRIMARY KEY)`); err != nil {
		return fmt.Errorf("create migration table: %w", err)
	}
	entries, err := fs.Glob(migrations.Files, "*.up.sql")
	if err != nil {
		return fmt.Errorf("list migrations: %w", err)
	}
	sort.Strings(entries)
	for _, name := range entries {
		var applied bool
		if err := db.QueryRowContext(ctx, `SELECT EXISTS (SELECT 1 FROM schema_migrations WHERE filename = $1)`, name).Scan(&applied); err != nil {
			return fmt.Errorf("check migration %s: %w", name, err)
		}
		if applied {
			continue
		}
		sqlBytes, err := migrations.Files.ReadFile(name)
		if err != nil {
			return fmt.Errorf("read migration %s: %w", name, err)
		}
		tx, err := db.BeginTx(ctx, nil)
		if err != nil {
			return err
		}
		if _, err = tx.ExecContext(ctx, string(sqlBytes)); err == nil {
			_, err = tx.ExecContext(ctx, `INSERT INTO schema_migrations (filename) VALUES ($1)`, name)
		}
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("apply migration %s: %w", name, err)
		}
		if err := tx.Commit(); err != nil {
			return fmt.Errorf("commit migration %s: %w", name, err)
		}
	}
	return nil
}

func greetingHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getGreeting(w, r, db)
		case http.MethodPut:
			putGreeting(w, r, db)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

func getGreeting(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	var greeting string
	if err := db.QueryRowContext(r.Context(), `SELECT text FROM greetings WHERE id = 1`).Scan(&greeting); err != nil {
		writeAPIError(w, http.StatusInternalServerError, "internal_error", "Internal server error.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"greeting": greeting})
}

func putGreeting(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	greeting, ok := parseGreetingRequest(w, r)
	if !ok {
		return
	}

	var saved string
	err := db.QueryRowContext(
		r.Context(),
		`UPDATE greetings SET text = $1, updated_at = now() WHERE id = 1 RETURNING text`,
		greeting,
	).Scan(&saved)
	if err != nil {
		writeAPIError(w, http.StatusInternalServerError, "internal_error", "Internal server error.")
		return
	}
	writeJSON(w, http.StatusOK, map[string]string{"greeting": saved})
}

func parseGreetingRequest(w http.ResponseWriter, r *http.Request) (string, bool) {
	decoder := json.NewDecoder(r.Body)

	var body struct {
		Greeting *string `json:"greeting"`
	}
	if err := decoder.Decode(&body); err != nil {
		writeAPIError(w, http.StatusBadRequest, "invalid_request", "Invalid request.")
		return "", false
	}
	if err := decoder.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		writeAPIError(w, http.StatusBadRequest, "invalid_request", "Invalid request.")
		return "", false
	}
	if body.Greeting == nil {
		writeAPIError(w, http.StatusBadRequest, "invalid_request", "Invalid request.")
		return "", false
	}
	greeting := strings.TrimSpace(*body.Greeting)
	if greeting == "" {
		writeAPIError(w, http.StatusBadRequest, "invalid_request", "Greeting must not be blank.")
		return "", false
	}
	return greeting, true
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func writeAPIError(w http.ResponseWriter, status int, code string, message string) {
	writeJSON(w, status, map[string]any{"error": map[string]string{"code": code, "message": message}})
}
