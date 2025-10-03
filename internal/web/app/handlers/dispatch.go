package handlers

import (
	"net/http"
)

func Dispatch(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/dispatch", http.StatusFound)
}