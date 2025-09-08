package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/company/template"
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

func (h *Handler) Company(w http.ResponseWriter, r *http.Request) {
	template.CompanyPage().Render(r.Context(), w)
}
