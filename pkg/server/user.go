package server

import (
	"main/pkg/common"
	"net/http"
)

func (app *App) UserIdentity(w http.ResponseWriter, r *http.Request) {
	// Get user from context
	user := r.Context().Value("userIdentity")

	// Return user identity
	common.HTTPOK[any](w, user, nil)
}
