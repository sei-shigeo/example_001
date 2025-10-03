package handlers

import (
	"net/http"
	"project/internal/web/app/views/layouts"
	"project/internal/web/app/views/pages"
	"strconv"
)

func (app *AppHandlers) EmployeesRoute() error {

	// 従業員一覧
	app.mux.HandleFunc("GET /employees", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		employees, err := app.db.Queries.GetEmployees(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		layouts.Base("従業員一覧", pages.Employees(employees)).Render(ctx, w)
	})

	// 従業員詳細
	app.mux.HandleFunc("GET /employees/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		employee, err := app.db.Queries.GetEmployeeByID(ctx, int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages.EmployeeByID(employee).Render(ctx, w)

	})

	// 従業員作成
	app.mux.HandleFunc("POST /employees/create", func(w http.ResponseWriter, r *http.Request) {
		employees, err := app.db.Queries.GetEmployees(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages.Employees(employees).Render(r.Context(), w)
	})

	// 従業員更新
	app.mux.HandleFunc("PATCH /employees/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		employees, err := app.db.Queries.GetEmployees(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages.Employees(employees).Render(r.Context(), w)
	})

	return nil
}
