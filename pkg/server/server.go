package server

import (
	"main/pkg/config"
	"main/pkg/middleware"
	"main/pkg/model"
	"net/http"
)

type App struct {
	firebaseClient *config.FirebaseClient
	minioClient    *config.MinioClient
	db             *model.DB
	router         http.Handler
}

func Router(firebaseClient *config.FirebaseClient, minioClient *config.MinioClient, db *model.DB) *App {
	app := &App{
		firebaseClient: firebaseClient,
		minioClient:    minioClient,
		db:             db,
	}

	api := http.NewServeMux()
	apiV1JWTAuth := http.NewServeMux()
	apiV1Public := http.NewServeMux()

	// VERSION 1
	/* ----- JWT Firebase Auth ----- */
	// Add more here (No need to add '/api')
	apiV1JWTAuth.HandleFunc("GET /whoami", app.UserIdentity)

	/* ----- Public Route ----- */
	// Add more here (No need to add '/api/public')

	// Handle middleware
	apiWithV1JWTAuth := middleware.FirebaseJWTAuth(apiV1JWTAuth, app.firebaseClient, app.db)

	// Handle API routes
	api.Handle("/api/v1/", http.StripPrefix("/api/v1", apiWithV1JWTAuth))
	api.Handle("/api/v1/public/", http.StripPrefix("/api/v1/public", apiV1Public))

	app.router = api

	return app
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.router.ServeHTTP(w, r)
}
