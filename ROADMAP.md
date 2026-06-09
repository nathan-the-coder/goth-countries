# Backend Learning Roadmap

A collection of backend improvement ideas to level up your Go skills.
Each item introduces a new stdlib package or Go concept you can practice on this project.

---

## 1. Middleware Chain

**Goal:** Add logging, request-ID, panic recovery, and timeout middleware.

**Go stdlib:** `net/http` (middleware pattern), `log/slog`, `context`

**Why learn it:** The `func(next http.Handler) http.Handler` pattern is the
universal middleware idiom in Go. You'll see it in every production Go codebase.

**Sketch:**
```go
func RequestID(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        id := r.Header.Get("X-Request-Id")
        if id == "" {
            id = uuid.New().String()
        }
        ctx := context.WithValue(r.Context(), "req_id", id)
        w.Header().Set("X-Request-Id", id)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

---

## 2. Structured Logging

**Goal:** Replace `log.Println` with `slog` (Go 1.21+), and include request-scoped
attributes (method, path, duration).

**Go stdlib:** `log/slog`

**Why learn it:** Structured logs are machine-parseable and essential for any
non-trivial service. `slog` is now the standard and integrates with most
observability tools.

**Sketch:**
```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("request", "method", r.Method, "path", r.URL.Path, "duration", dur)
```

---

## 3. Context Propagation

**Goal:** Pass request-scoped values (request ID, auth info) through the handler
chain via `context.Context`.

**Go stdlib:** `context`

**Why learn it:** Context is Go's canonical way to carry deadlines,
cancellation signals, and request-scoped values. Every database call, HTTP
client request, and goroutine you spawn will use it.

**Sketch:**
```go
type key string
const ReqIDKey key = "request_id"

func GetReqID(ctx context.Context) string {
    v, _ := ctx.Value(ReqIDKey).(string)
    return v
}
```

---

## 4. Handler Tests

**Goal:** Write `httptest.NewRecorder` + `httptest.NewRequest` tests for
`CountryList`, `CountryDetail`, and `SearchCountries`.

**Go stdlib:** `net/http/httptest`, `testing`

**Why learn it:** Testing HTTP handlers without spinning up a real server is a
superpower. This is the single highest-leverage skill you can add — it forces
you to write testable, decoupled code.

**Sketch:**
```go
func TestCountryList(t *testing.T) {
    req := httptest.NewRequest("GET", "/", nil)
    rec := httptest.NewRecorder()
    handlers.CountryList(rec, req)
    if rec.Code != http.StatusOK {
        t.Fatalf("expected 200, got %d", rec.Code)
    }
}
```

---

## 5. JSON Data Loading

**Goal:** Move country data from a hardcoded Go slice to a `countries.json` file
loaded at startup.

**Go stdlib:** `encoding/json`, `os`, `embed`

**Why learn it:** Separating config/data from code is a core software engineering
principle. You'll also learn `embed` (Go 1.16+) to bundle the JSON into the
binary for production deployments.

**Sketch:**
```go
//go:embed countries.json
var data []byte

func LoadCountries() ([]Country, error) {
    var countries []Country
    err := json.Unmarshal(data, &countries)
    return countries, err
}
```

---

## 6. Concurrent Access with `sync.RWMutex`

**Goal:** Add a write endpoint (e.g. admin add/remove country) protected by a
`sync.RWMutex` so reads aren't blocked by each other.

**Go stdlib:** `sync`

**Why learn it:** Concurrency safety is a Go hallmark. `RWMutex` lets you have
multiple concurrent readers while still protecting writes — a pattern you'll
use constantly.

**Sketch:**
```go
type Store struct {
    mu        sync.RWMutex
    Countries []Country
}

func (s *Store) GetAll() []Country {
    s.mu.RLock()
    defer s.mu.RUnlock()
    return s.Countries
}
```

---

## 7. Efficient Routing (Map / Trie)

**Goal:** Replace the linear slice loop in handlers with a `map[string]Country`
lookup.

**Go stdlib:** built-in `map` type

**Why learn it:** Maps are Go's fundamental data structure for O(1) lookups.
Understanding when to use a map vs a slice is a basic but critical design
decision.

**Sketch:**
```go
var countryMap map[string]Country // built once at startup

func init() {
    countryMap = make(map[string]Country, len(Countries))
    for _, c := range Countries {
        countryMap[c.Name] = c
    }
}
```

---

## Suggested Order

| # | Topic               | Difficulty | Why first?                          |
|---|---------------------|------------|-------------------------------------|
| 4 | Handler tests       | Easy       | Unlocks safe refactoring            |
| 1 | Middleware          | Easy-Med   | Ubiquitous pattern                  |
| 2 | Structured logging  | Easy       | Quick win, big payoff               |
| 5 | JSON data           | Easy       | Separates config from code          |
| 3 | Context             | Medium     | Builds on middleware work           |
| 7 | Map lookup          | Easy       | Simple performance fix              |
| 6 | RWMutex             | Medium     | Adds concurrency to the mix         |

Start with **handler tests** — they'll make every subsequent change safer and
more fun.
