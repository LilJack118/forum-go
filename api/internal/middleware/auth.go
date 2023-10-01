package middleware

import (
	"forum/api/pkg/httpErrors"
	"forum/api/pkg/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func AuthJWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")

		if bearer == "" {
			httpErrors.JSONError(w, "no bearer token provided", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(bearer, " ")
		if len(tokenParts) != 2 {
			msg := "invalid authorization header format. it should be 'bearer token'"
			httpErrors.JSONError(w, msg, http.StatusUnauthorized)
			return
		}

		tokenString := tokenParts[1]

		auth_jwt, err := utils.AuthJWT()
		if err != nil {
			log.Print(err)
			httpErrors.JSONError(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		id, err := auth_jwt.GetUserID(tokenString)
		if err != nil {
			httpErrors.JSONError(w, err.Error(), http.StatusUnauthorized)
			return
		}

		// set user id
		mux.Vars(r)["uid"] = id

		next.ServeHTTP(w, r)
	})
}
