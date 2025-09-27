package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/customer/template"
	"project/internal/web/app/layouts"
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

func (h *Handler) Customer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	customers, err := h.DB.Queries.GetCustomers(ctx)
	if err != nil {
		http.Error(w, "顧客データの取得に失敗しました", http.StatusInternalServerError)
		return
	}
	d := template.Props{
		Customers: customers,
		Title:     "顧客一覧",
	}

	layouts.Base(d.Title, d.CustomerPage()).Render(ctx, w)
}
