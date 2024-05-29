package handlers

import (
	"log/slog"
	"net/http"
)

type HttpHandler func(w http.ResponseWriter, r *http.Request) error

func Func(handler HttpHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err != nil {
			slog.Error("handler error", err)
			w.WriteHeader(500)
			w.Write([]byte("internal error"))
		}

	}

}
