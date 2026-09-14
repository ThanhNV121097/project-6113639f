# Architecture — Hello World Acceptance 3

## Stack

| Part | Choice | Reason |
|---|---|---|
| Frontend | Next.js 15 App Router, TypeScript, Tailwind v3 | Required UI stack; standalone output fits fixed image |
| Backend | Go 1.22, `net/http`, `database/sql`, pgx stdlib driver | Small API; no web framework needed |
| Database | PostgreSQL 16 | Required persisted shared greeting |

## Layout

- `code/backend/cmd/api`: only executable; boots migrations, health endpoint, later API routes.
- `code/backend/migrations`: embedded ordered SQL migration pairs.
- `code/frontend/app`: App Router shell. `page.tsx` remains composition root.
- `code/frontend/components`: one default-exported story component per story.
- `code/frontend/lib/mock`: UI-phase fixtures, deleted when API lands.

## Contracts and conventions

- Backend reads `DATABASE_URL`, applies pending migrations, then listens on `PORT`, `APP_PORT`, or `8080`.
- `/healthz` returns 200 only after migration and database ping succeed.
- SQL uses parameterized queries. API validation trims greeting and rejects blank input.
- Database has one shared greeting. Latest completed write wins.
- API contract paths omit proxy-only `/api` prefix; see `services.md`.
- Go uses `cmd/api`, lower-case package names, timestamped migration pairs. React files default-export PascalCase components. `page.tsx` only composes components.
- Client-interactive React files start with literal first line `"use client"`; composition root stays server-side.
- All shared visual values use tokens from `app/globals.css`. No CSS motion.

## Environment

| Service | Variables |
|---|---|
| Backend | `DATABASE_URL`, `PORT`, optional `APP_PORT` fallback |
| Frontend | `NEXT_PUBLIC_API_URL`, `API_ORIGIN` |
| Compose database | `POSTGRES_USER`, `POSTGRES_PASSWORD`, `POSTGRES_DB` |

Examples list keys and comments, never secrets. Compose supplies local defaults.

## Run

1. Copy root `.env.example` to `.env` if local values differ.
2. Run `docker compose --profile local up --build` from repository root.
3. Open `http://localhost:3000`; backend health is `http://localhost:8080/healthz`.

## Decisions

| Decision | Rejected | Tradeoff |
|---|---|---|
| Self-migrate at backend boot | Manual migration command | Safe empty runtime DB; startup owns migration failure |
| `net/http` instead of framework | Router dependency | Less abstraction for two routes |
| One-row table with fixed ID | Per-visitor records | Matches one shared greeting; no identity model |
| Next standalone output | Custom Node server | Required by fixed frontend image |

## Rollout and unknowns

Migration is additive and tracked in `schema_migrations`; rollback uses paired down SQL only during controlled recovery. No external credentials, sign-in, or provider setup. Greeting length limit is not specified; contract accepts non-blank trimmed text.
