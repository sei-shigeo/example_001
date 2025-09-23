package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/employees/template"
	"project/internal/web/app/layouts"
)

type Handler struct {
	DB *db.DB
}

func NewHandler(database *db.DB) *Handler {
	return &Handler{
		DB: database,
	}
}

func (h *Handler) Employees(w http.ResponseWriter, r *http.Request) {
	// データベースから従業員データを取得
	ctx := r.Context()
	employees, err := h.DB.Queries.GetEmployees(ctx)
	if err != nil {
		http.Error(w, "従業員データの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	// 従業員データを渡す
	d := template.Props{
		Employees: employees,
		Title:     "従業員一覧",
	}

	// テンプレートに従業員データを渡してレンダリング
	layouts.Base(d.Title, d.EmployeesPage()).Render(ctx, w)
}
