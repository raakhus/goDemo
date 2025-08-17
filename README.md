# Go Demo (Dockerized)

Go + React minimal auth demo. Log in, set an HttpOnly JWT cookie, then access a protected page that says **“You made it.”**

## Quick start

```bash
cp .env.example .env
install go 1.22
run go mod tidy
docker compose up --build
```

Then open http://localhost:5173

**Demo creds:** `test@example.com` / `Passw0rd!`

### Notes
- The React dev server proxies `/api/*` to the `api` service, so cookies work without CORS hassle.
- Cookie is `HttpOnly`, `SameSite=Lax`. `COOKIE_SECURE` is `false` in dev; set to `true` in prod (HTTPS).
