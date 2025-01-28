package middleware

import (
	"main/pkg/common"
	"main/pkg/config"
	"main/pkg/model"
	"net/http"
	"strings"

	"golang.org/x/net/context"
)

func FirebaseJWTAuth(next http.Handler, firebaseClient *config.FirebaseClient, db *model.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		tokenStr := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
		if tokenStr == "" {
			common.HTTPError(w, "Authorization token is required", http.StatusBadRequest)

			return
		}

		verifiedToken, err := firebaseClient.AuthClient.VerifyIDToken(ctx, tokenStr)
		if err != nil {
			common.HTTPError(w, err.Error(), http.StatusUnauthorized)

			return
		}

		email, ok := verifiedToken.Claims["email"].(string)
		if !ok {
			common.HTTPError(w, "Email not found in token", http.StatusUnauthorized)

			return
		}

		userIdentity, err := db.FetchUserIdentity(email)
		if err != nil {
			common.HTTPError(w, err.Error(), http.StatusUnauthorized)

			return
		}

		r = r.WithContext(context.WithValue(ctx, "userIdentity", userIdentity))

		next.ServeHTTP(w, r)
	})
}
