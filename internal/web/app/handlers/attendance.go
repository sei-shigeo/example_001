package handlers

import (
	"net/http"
)

func Attendance(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/attendance", http.StatusFound)
}