package router

import (
	"database/sql"
	"github.com/TechBowl-japan/go-stations/handler"
	"net/http"
)

func NewRouter(todoDB *sql.DB) *http.ServeMux {
	// register routes
	mux := http.NewServeMux()
	mux.Handle("/healthz", &handler.HealthzHandler{})
	return mux
}
