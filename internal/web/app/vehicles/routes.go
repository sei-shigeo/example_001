package vehicles

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/vehicles/handler"
)

// RegisterRoutes ユーザー関連のルートを登録
func RegisterRoutes(mux *http.ServeMux, database *db.DB) {
	handler := handler.NewHandler(database)

	// ユーザー関連のルート
	mux.HandleFunc("GET /vehicles", handler.GetVehicles)
	mux.HandleFunc("GET /vehicles/create", handler.CreateVehicleForm)
	mux.HandleFunc("GET /vehicles/edit/{id}", handler.EditVehicleForm)
	// POST, PATCH, DELETE
	mux.HandleFunc("POST /vehicles/create", handler.CreateVehicle)
	mux.HandleFunc("POST /vehicles/validate", handler.ValidateVehicle)
	mux.HandleFunc("PATCH /vehicles/edit/{id}", handler.UpdateVehicle)
	mux.HandleFunc("DELETE /vehicles/delete/{id}", handler.DeleteVehicle)
}
