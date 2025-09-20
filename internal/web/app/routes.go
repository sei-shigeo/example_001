package app

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/company"
	"project/internal/web/app/login"
	"project/internal/web/app/employees"
	"project/internal/web/app/vehicles"
)

// Routes アプリケーション全体のルートを設定
func Routes(database *db.DB) func(*http.ServeMux) {
	return func(mux *http.ServeMux) {
		// 各モジュールのルートを登録
		employees.RegisterRoutes(mux, database)
		company.RegisterRoutes(mux, database)
		login.RegisterRoutes(mux, database)
		vehicles.RegisterRoutes(mux, database)

		// グローバルなルート（ヘルスチェックなど）
		mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK"))
		})
	}
}
