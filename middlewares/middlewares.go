package middlewares

import (
	"go-project/response"
	"go-project/security"
	"net/http"
)

func AuthValidate(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := security.TokenValidate(r); err != nil {
			response.ErrorJSON(w, http.StatusUnauthorized, err)
			return
		}

		next(w, r)
	}
}
