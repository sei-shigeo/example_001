package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/layouts"
	"project/internal/web/app/login/template"
)

// Handler ハンドラー構造体
type Handler struct {
	DB *db.DB
}

// NewHandler 新しいハンドラーを作成
func NewHandler(database *db.DB) *Handler {
	return &Handler{
		DB: database,
	}
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	layouts.Base("Login", template.LoginPage()).Render(r.Context(), w)
}
