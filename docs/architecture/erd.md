# ERD — Greeting

## `greetings`

One row represents app-wide greeting. Seed migration creates fixed row `id = 1` with `Hello, World!`.

| Column | PostgreSQL type | Constraints | Purpose |
|---|---|---|---|
| `id` | `smallint` | primary key, check (`id = 1`) | Enforces one shared record |
| `text` | `text` | not null, check (`btrim(text) <> ''`) | Trimmed greeting shown to visitors |
| `updated_at` | `timestamptz` | not null, default `now()` | Most recent completed save time |

Relationships: none. `greetings` is singleton table.

## Migration rules

- Migration filenames use UTC timestamp prefix and paired `.up.sql` / `.down.sql` files.
- Backend records applied filenames in `schema_migrations`.
- Initial up migration creates table and inserts singleton default. Down migration drops `greetings`.

## Story extension — Editable persisted greeting

This story uses existing `greetings` singleton unchanged. No new entity, column, relationship, foreign key, or index is needed. `greetings.id = 1` supports singleton lookup and primary-key update; no secondary index is warranted.

### Mock shape review

Reviewed UI PR #6 mock: `GreetingResponse` is `{ "greeting": string }`; `readGreeting()` and `saveGreeting()` return that shape. It is sound and matches existing service contract response bodies exactly. Mock `localStorage` persistence is UI-only and must be deleted when backend integration replaces it; no frontend response-shape change is needed.

### Migration plan

- Forward: create `greetings` with `id smallint primary key check (id = 1)`, non-blank `text`, and `updated_at`; insert `(1, 'Hello, World!')`. Safe on populated databases because it creates isolated table and singleton row.
- Backward: drop `greetings`. Safe structurally, but destructive: it permanently removes latest shared greeting. Use only before production data matters; restore from backup if rollback is required after data is saved.
- Existing migration rules apply: timestamped paired SQL files and `schema_migrations` tracking.
