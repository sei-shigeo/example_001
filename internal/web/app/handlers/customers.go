package handlers

import (
	"net/http"
)

func Customers(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/customers", http.StatusFound)
}