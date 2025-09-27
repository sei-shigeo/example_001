package customer

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/customer/handler"
)

// RegisterRoutes 会社関連のルートを登録
func RegisterRoutes(mux *http.ServeMux, database *db.DB) {
	customerHandler := handler.NewHandler(database)


	// 会社関連のルート
	mux.HandleFunc("GET /customers", customerHandler.Customer)
	mux.HandleFunc("GET /customers/create", customerHandler.CreateCustomerForm)
	mux.HandleFunc("GET /customers/edit/{id}", customerHandler.EditCustomerForm)
	// POST, PATCH, DELETE
	mux.HandleFunc("POST /customers/create", customerHandler.CreateCustomer)
	mux.HandleFunc("POST /customers/validate", customerHandler.ValidateCustomer)
	mux.HandleFunc("PATCH /customers/edit/{id}", customerHandler.UpdateCustomer)
	mux.HandleFunc("DELETE /customers/delete/{id}", customerHandler.DeleteCustomer)

	// 将来の会社関連ルート
	// mux.HandleFunc("/companies", companyHandler.GetCompanies)
	// mux.HandleFunc("/companies/{id}", companyHandler.GetCompany)
	// mux.HandleFunc("/companies", companyHandler.CreateCompany).Methods("POST")
}
