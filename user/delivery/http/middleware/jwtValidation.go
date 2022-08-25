package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/ihsanbudiman/notes_app/helpers"
)

func MyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get token from header
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "token not found", http.StatusUnauthorized)
			return
		}

		// split token
		splittedToken := strings.Split(token, " ")
		if len(splittedToken) != 2 {
			http.Error(w, "token not found", http.StatusUnauthorized)
			return
		}

		// get token
		token = splittedToken[1]
		if token == "" {
			http.Error(w, "token not found", http.StatusUnauthorized)
			return
		}

		// call usecase
		tokenClaims, err := helpers.ValidateJwt(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// set context
		ctx := context.WithValue(r.Context(), "credentials", tokenClaims)
		r = r.WithContext(ctx)

		// call next handler
		next.ServeHTTP(w, r)

	})
}
