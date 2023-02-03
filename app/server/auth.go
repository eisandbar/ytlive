package server

import (
	"log"
	"net/http"

	"github.com/eisandbar/ytlive/app/internal"
)

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key, err := internal.LoadToken("keys/server_key.txt")
		if err != nil {
			log.Println(err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		token := r.Header.Get("Authorization")

		if token == key {
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Forbidden", http.StatusForbidden)
		}
	})
}
