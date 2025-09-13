package user

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/user/handler"
)

// RegisterRoutes ユーザー関連のルートを登録
func RegisterRoutes(mux *http.ServeMux, database *db.DB) {
	userHandler := handler.NewHandler(database)

	// ユーザー関連のルート
	mux.HandleFunc("/employees", userHandler.Employees)

	// 将来のユーザー関連ルート
	// mux.HandleFunc("/users", userHandler.GetUsers)
	// mux.HandleFunc("/users/{id}", userHandler.GetUser)
	// mux.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
}
