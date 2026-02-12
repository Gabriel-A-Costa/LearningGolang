package middleware

import (
	"fmt"
	"net/http"
)

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Printf(
				"Método: %s, Rota: %s\n",
				r.Method,
				r.URL.Path,
			)
			next.ServeHTTP(w, r)
		},
	)
}
