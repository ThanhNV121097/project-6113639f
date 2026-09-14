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
