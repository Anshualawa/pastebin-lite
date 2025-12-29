# Pastebin-Lite
a lightweight Pastebin-like service built using Golang gin framework and Redis

User can create text pastes, receive shareable links, and 
view them until they expire or exceed view limits.

---

## Tech Stack
- Golang
- Gin Web Framework
- Redis (Presistence)

---

## Features
- Create a paste with optional TTL and view limits
- Fetch paste via API (counts as a view)
- View paste via browser (HTML)
- Time-based expiry (TTL)
- View-count limits
- Deternministic time support for automated tests
- Redis-backed persistence

---

## Persistence Layer
Redis is used to store pastes and metadata.
This ensures data survives across requests and works correctly in 
serverless or distributed environments

---

## Run Locally

### Requirements
- Go 1.24+
- Redis

### Steps
```bash
go mod tidy
```
```bash
export REDIS_ADDR=localhost:6379
export REDIS_PASSWORD=
```
```bash
go run  cmd/main.go
```

#### Health Check 
```bash
GET /api/healthz
```
Response
```json
{ "ok" : true }
```

## Create Paste

### Endpoint
POST /api/pastes

### Request Body
```base
curl  -X  POST  http://localhost:8080/api/pastes \
  -H "Content-Type: application/json" \
  -d '{"content":"Hello Pastebin","ttl_seconds":60,"max_views":3}'
```
Response
```json
{
    "id" : "string",
    "url" : "http://localhost:8080/p/{id}"
}
```

## Fetch Paste (API)

### Endpoint
```base
curl -X GET http://localhost:8080/api/pastes/{id}
```
### Success Response
```json
{
  "content" : "string",
  "remaining_views" : 4,
  "expires_at" : "2025-12-29T00:00:00.000Z"
}
```

## View Paste (HTML)

### Endpoitn
```bash
http://loclahost:8080/p/{id}
```
- Returns an HTML page displaying the paste content
- Paste content is rendered safely (HTML escaped)
- Expired or unavailable pastes return HTTP 404