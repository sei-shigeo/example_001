package handlers

import (
	"fmt"
	"net/http"
	"project/internal/db"
	"project/internal/web/app/views/layouts"
	"project/internal/web/app/views/pages"
	"strconv"

	"github.com/starfederation/datastar-go/datastar"
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

		sig := &pages.EmployeesSignal{}
		if err := datastar.ReadSignals(r, sig); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		employee, err := app.db.Queries.GetEmployeeByID(ctx, int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sig.Edit.Name = employee.Name
		sig.Edit.Email = employee.Email
		sig.Edit.Phone = employee.Phone
		sig.Edit.IsShow = true

		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(
			pages.EmployeeEdit(employee),
			datastar.WithSelectorID("employeeEditForm"),
			datastar.WithModeInner(),
		)

		sse.MarshalAndPatchSignals(sig)

	})

	// 従業員作成
	app.mux.HandleFunc("POST /employees/create", func(w http.ResponseWriter, r *http.Request) {

		sig := &pages.EmployeesSignal{}
		if err := datastar.ReadSignals(r, sig); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		create := &db.CreateEmployeeParams{
			Name:  sig.Create.Name,
			Email: sig.Create.Email,
			Phone: sig.Create.Phone,
		}

		employees, err := app.db.Queries.CreateEmployee(r.Context(), create)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sse := datastar.NewSSE(w, r)
		sse.PatchElementTempl(
			pages.EmployeesList(employees), 
			datastar.WithModeAppend(), 
			datastar.WithSelectorID("employeeList"))
	})

	// 従業員更新
	app.mux.HandleFunc("PATCH /employees/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sig := &pages.EmployeesSignal{}
		if err := datastar.ReadSignals(r, sig); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		update := &db.UpdateEmployeeParams{
			ID:    int32(id),
			Name:  sig.Edit.Name,
			Email: sig.Edit.Email,
			Phone: sig.Edit.Phone,
		}

		employee, err := app.db.Queries.UpdateEmployee(ctx, update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		sse := datastar.NewSSE(w, r)

		sse.PatchElementTempl(
			pages.EmployeesList(employee),
			datastar.WithSelectorID(fmt.Sprintf("employee-%d", employee.ID)),
			datastar.WithModeReplace(),
		)
	})

	// 従業員削除
	app.mux.HandleFunc("DELETE /employees/delete/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := app.db.Queries.DeleteEmployee(ctx, int32(id)); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		selectorID := fmt.Sprintf("employee-%d", id)
		sse := datastar.NewSSE(w, r)
		sse.RemoveElementByID(selectorID)

	})

	return nil
}
