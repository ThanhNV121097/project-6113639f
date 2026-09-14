# Service contracts — Greeting API

Backend routes receive paths after proxy removes `/api`. All JSON responses use `Content-Type: application/json; charset=utf-8`.

## Error envelope

All API errors use:

```json
{"error":{"code":"invalid_request","message":"Greeting must not be blank."}}
```

`code` is stable machine value. `message` is safe plain text. Invalid JSON or invalid greeting returns `400`; database failure returns `500` with `internal_error` and generic message.

## Endpoints

### `GET /v1/greeting`

Returns current singleton greeting. Missing singleton is initialized by migration.

Response `200`:

```json
{"greeting":"Hello, World!"}
```

### `PUT /v1/greeting`

Replaces shared greeting. Server trims whitespace before validation and persistence. No authentication.

Request:

```json
{"greeting":"Hello again"}
```

Response `200`:

```json
{"greeting":"Hello again"}
```

Errors: `400` for malformed JSON, missing greeting, non-string greeting, or blank trimmed greeting; `500` for unavailable persistence.

## Story extension — Editable persisted greeting

UI PR #6 mock returns `{ "greeting": string }` for reads and saves. Existing endpoint response shape matches it; backend integration replaces mock functions and deletes `code/frontend/lib/mock/editable-persisted-greeting.ts`. No component contract change.

### Endpoint execution details

| Endpoint | Auth | Success | Error codes |
|---|---|---|---|
| `GET /v1/greeting` | none | `200` with `{ "greeting": string }` | `500 internal_error` when persistence is unavailable or query fails |
| `PUT /v1/greeting` | none | `200` with trimmed persisted `{ "greeting": string }` | `400 invalid_request` for malformed JSON, missing/non-string `greeting`, or blank after trim; `500 internal_error` when persistence is unavailable or update fails |

`PUT` must use parameterized SQL, update singleton `id = 1`, and set `updated_at = now()`. Latest completed successful update wins. No pagination applies: service exposes one singleton resource. Missing singleton is migration invariant; if breached, return `500 internal_error`, never manufacture an unpersisted response.
