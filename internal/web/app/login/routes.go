package login

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/login/handler"
)

// RegisterRoutes 認証関連のルートを登録
func RegisterRoutes(mux *http.ServeMux, database *db.DB) {
	loginHandler := handler.NewHandler(database)

	// 認証関連のルート
	mux.HandleFunc("/login", loginHandler.Login)

	// 将来の認証関連ルート
	// mux.HandleFunc("/logout", loginHandler.Logout)
	// mux.HandleFunc("/register", loginHandler.Register)
	// mux.HandleFunc("/forgot-password", loginHandler.ForgotPassword)
}
