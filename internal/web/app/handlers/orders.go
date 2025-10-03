package handlers

import (
	"net/http"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/orders", http.StatusFound)
}