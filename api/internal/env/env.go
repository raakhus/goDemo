package env

import "os"

type Config struct {
	APIPort      string
	JWTSecret    string
	CookieName   string
	CookieSecure bool
}

func Load() Config {
	return Config{
		APIPort:      get("API_PORT", "8080"),
		JWTSecret:    get("JWT_SECRET", "dev-secret"),
		CookieName:   get("COOKIE_NAME", "app_session"),
		CookieSecure: get("COOKIE_SECURE", "false") == "true",
	}
}

func get(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}
