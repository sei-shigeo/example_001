package app

import (
	"net/http"
	"project/internal/db"
	"project/internal/logger"
	"project/internal/web/app/handlers"
)

func Routes(database *db.DB) func(*http.ServeMux) {
	return func(mux *http.ServeMux) {
		handlers := handlers.NewAppHandlers(database, mux)

		// 従業員関連
		if err := handlers.EmployeesRoute(); err != nil {
			logger.Error("Failed to set up employees route", "error", err)
		}
		// 顧客関連
		if err := handlers.CustomersRoute(); err != nil {
			logger.Error("Failed to set up customers route", "error", err)
		}
		// 車両関連
		if err := handlers.VehiclesRoute(); err != nil {
			logger.Error("Failed to set up vehicles route", "error", err)
		}

	}
}
