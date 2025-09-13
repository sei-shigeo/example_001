package company

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/company/handler"
)

// RegisterRoutes 会社関連のルートを登録
func RegisterRoutes(mux *http.ServeMux, database *db.DB) {
	companyHandler := handler.NewHandler(database)

	// 会社関連のルート
	mux.HandleFunc("/company", companyHandler.Company)

	// 将来の会社関連ルート
	// mux.HandleFunc("/companies", companyHandler.GetCompanies)
	// mux.HandleFunc("/companies/{id}", companyHandler.GetCompany)
	// mux.HandleFunc("/companies", companyHandler.CreateCompany).Methods("POST")
}
