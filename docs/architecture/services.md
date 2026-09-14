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
