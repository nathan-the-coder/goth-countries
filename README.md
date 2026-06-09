# 🏰 Goth Countries

A dark-themed demo app built with **Go**, **Templ**, **HTMX**, and **Tailwind CSS v4**.
Browse countries of the world with instant search — no JavaScript framework required.

## Stack

| Layer      | Technology                        |
|------------|-----------------------------------|
| Backend    | Go 1.26+ (`net/http` ServeMux)   |
| Templates  | Templ (type-safe Go HTML)         |
| Frontend   | HTMX 2.x (hypermedia-driven)      |
| Styling    | Tailwind CSS v4                   |

## Quick Start

```sh
# 1. Build CSS
./tailwindcss -i src/input.css -o public/css/tailwind.css

# 2. Generate templ code
templ generate

# 3. Build Go binary
go build -o ./tmp/main .

# 4. Run (or use `air` for hot reload)
./tmp/main
# App at http://localhost:8080
```

### Dev server with hot reload

```sh
air
# App at http://localhost:8080, air proxy at http://localhost:9090
```

## Routes

| Method | Path                 | Description                  |
|--------|----------------------|------------------------------|
| GET    | `/`                  | Country list (all countries) |
| GET    | `/country/{name}`    | Country detail page          |
| GET    | `/search?search=...` | Search countries (HTMX)      |

## Project Structure

```
.
├── main.go                  # Entrypoint & route registration
├── handlers/                # HTTP handlers
│   └── handlers.go
├── components/              # Templ templates (*.templ + generated *_templ.go)
│   ├── base.templ           # Base layout (nav, HTML shell)
│   ├── countries.templ      # Country list & search
│   └── country.templ        # Country detail
├── models/
│   └── models.go            # Country struct & in-memory data
├── src/
│   └── input.css            # Tailwind entry point (@import "tailwindcss")
├── public/
│   └── css/tailwind.css     # Compiled CSS (committed)
├── tailwindcss              # Standalone Tailwind CLI (committed)
└── .air.toml                # Air hot-reload config
```

## Dependencies

- [Go](https://go.dev) 1.26+
- [Templ](https://github.com/a-h/templ) — `go install github.com/a-h/templ/cmd/templ@latest`
- [Tailwind CSS v4](https://tailwindcss.com) — standalone CLI (prebuilt binary in repo root)
- [HTMX](https://htmx.org) — loaded from CDN

## License

Apache 2.0 — see [LICENSE](LICENSE).
