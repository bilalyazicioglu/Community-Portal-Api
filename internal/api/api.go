package api

import (
	"database/sql"
	"net/http"

	"github.com/MACSEC-Proje-Gelistirme/Community-Portal-Api/internal/middleware"
	"github.com/gorilla/mux"
)

type Router struct {
	router *mux.Router
	db     *sql.DB
}

func NewRouter(db *sql.DB) *Router {
	r := mux.NewRouter()
	ro := &Router{
		router: r,
		db:     db,
	}

	// Rotaları burada tanımlayın
	ro.router.HandleFunc("/api/user", ro.GetUser).Methods("GET")
	ro.router.HandleFunc("/api/user", ro.CreateUser).Methods("POST")
	return &Router{
		db: db,
	}
}

func (r *Router) NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.Use(middleware.CorsMiddleware)

	router.HandleFunc("/user", r.CreateUser).Methods(http.MethodPost, http.MethodOptions)

	protected := router.PathPrefix("/api").Subrouter()
	protected.Use(middleware.EnsureValidToken)

	return router
}
