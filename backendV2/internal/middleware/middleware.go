package middleware

import (
	"log/slog"
	"net/http"
)

type MiddleWare struct {
	log *slog.Logger
}

func New(log *slog.Logger) *MiddleWare {
	return &MiddleWare{log: log}
}

func (mw *MiddleWare) UseHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Header().Add("Access-Control-Allow-Headers", "*")
		next.ServeHTTP(w, r)
	})
}
