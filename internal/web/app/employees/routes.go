package employees

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/employees/handler"
)

// RegisterRoutes ユーザー関連のルートを登録
func RegisterRoutes(mux *http.ServeMux, database *db.DB) {
	handler := handler.NewHandler(database)

	// ユーザー関連のルート
	mux.HandleFunc("GET /employees", handler.Employees)
	mux.HandleFunc("GET /employees/create", handler.CreateEmployeeForm)
	mux.HandleFunc("GET /employees/edit/{id}", handler.EditEmployeeForm)
	// POST, PATCH, DELETE
	mux.HandleFunc("POST /employees/create", handler.CreateEmployee)
	mux.HandleFunc("PATCH /employees/edit/{id}", handler.UpdateEmployee)
	mux.HandleFunc("DELETE /employees/delete/{id}", handler.DeleteEmployee)
	mux.HandleFunc("POST /employees/validate", handler.ValidateEmployee)
	// 将来のユーザー関連ルート
	// mux.HandleFunc("/users", userHandler.GetUsers)
	// mux.HandleFunc("/users/{id}", userHandler.GetUser)
	// mux.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
}
