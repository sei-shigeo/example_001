package company

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/company/handler"
)

// Routes 会社関連のルートを設定
func Routes(mux *http.ServeMux, database *db.DB) {
	h := handler.NewHandler(database)

	mux.HandleFunc("/company", h.Company)
}
