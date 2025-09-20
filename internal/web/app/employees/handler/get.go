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

	// タイトルを設定
	title := "Employees"

	// 従業員データを渡す
	props := template.Props{
		Employees: employees,
	}

	// テンプレートに従業員データを渡してレンダリング
	layouts.Base(title, props.EmployeesPage()).Render(ctx, w)
}




