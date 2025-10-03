package handlers

import (
	"project/internal/db"
	"net/http"
)


type AppHandlers struct {
	db *db.DB
	mux *http.ServeMux
}

func NewAppHandlers(database *db.DB, mux *http.ServeMux) *AppHandlers {
	return &AppHandlers{database, mux}
}
