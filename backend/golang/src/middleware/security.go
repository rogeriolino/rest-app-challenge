package middleware

import (
  "app"
  "model"
  "net/http"
)

type Security struct {
  App app.App
}

func (s *Security) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    allowed := r.URL.Path == "/token"

    if !allowed {
      token := model.Token{}
      row := s.App.DB.QueryRow("SELECT t.username, t.token FROM `tokens` t WHERE t.token = ?", r.Header.Get("X-Token"))
      err := row.Scan(&token.Username, &token.Token)
      allowed = err == nil
    }

		if allowed {
			next.ServeHTTP(w, r)
		} else {
      s.App.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
