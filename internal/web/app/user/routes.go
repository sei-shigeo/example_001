package user

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/user/handler"
)

// Routes ユーザー関連のルートを設定
func Routes(mux *http.ServeMux, database *db.DB) {
	h := handler.NewHandler(database)

	mux.HandleFunc("/", h.Login)
	mux.HandleFunc("/employees", h.Employees)
}
