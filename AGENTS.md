# goth-countries

Go + Templ + HTMX + Tailwind CSS v4 demo app.

## Quick start

```sh
# Build CSS, generate templ code, build Go binary (all three required):
./tailwindcss -i src/input.css -o public/css/tailwind.css
templ generate
go build -o ./tmp/main .

# Dev server with hot reload (uses air):
air
# App at http://localhost:8080, air proxy at http://localhost:9090
```

## Architecture

- `main.go` — entrypoint, registers routes on `http.ServeMux`, serves static `/public/`
- `handlers/` — HTTP handlers: `CountryList`, `CountryDetail`, `SearchCountries`
- `components/` — `.templ` files (type-safe HTML templates) with generated `_templ.go` counterparts
- `models/` — `Country` struct with a hardcoded in-memory slice
- `src/input.css` — Tailwind entry point (`@import "tailwindcss"`)
- `public/css/tailwind.css` — compiled Tailwind output (committed artifact)

## Key conventions

- Edit only `.templ` files, never `_templ.go` — those are generated. After changing a `.templ` file, run `templ generate`.
- Tailwind v4 uses `@import "tailwindcss"` (not `@tailwind` directives).
- The `tailwindcss` binary is checked in (standalone CLI).
- All country data is in-memory in `models/models.go` — no database.

## Routes

| Method | Path                 | Handler                |
|--------|----------------------|------------------------|
| GET    | `/`                  | `CountryList`          |
| GET    | `/country/{name}`    | `CountryDetail`        |
| GET    | `/search?search=...` | `SearchCountries`      |

## Dependencies

- Go 1.26+
- `github.com/a-h/templ` — CLI needed: `go install github.com/a-h/templ/cmd/templ@latest`
- Tailwind CSS — standalone CLI (prebuilt binary `./tailwindcss` in repo root)
- HTMX — loaded from CDN in `base.templ`
