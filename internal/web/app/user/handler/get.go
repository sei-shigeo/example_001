package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/user/template"
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
	template.LoginPage().Render(r.Context(), w)
}

func (h *Handler) Employees(w http.ResponseWriter, r *http.Request) {
	// データベースから従業員データを取得
	ctx := r.Context()
	employees, err := h.DB.Queries.GetEmployees(ctx)
	if err != nil {
		http.Error(w, "従業員データの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// テンプレートに従業員データを渡してレンダリング
	template.EmployeesPage(employees).Render(ctx, w)
}
