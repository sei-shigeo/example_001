package handlers

import (
	"net/http"
)

func Vehicles(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/vehicles", http.StatusFound)
}