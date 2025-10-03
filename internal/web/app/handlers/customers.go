package handlers

import (
	"net/http"
	"project/internal/web/app/views/layouts"
	"project/internal/web/app/views/pages"
	"strconv"
)

func (app *AppHandlers) CustomersRoute() error {

	// 顧客一覧
	app.mux.HandleFunc("GET /customers", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		cusmers, err := app.db.Queries.GetCustomers(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		layouts.Base("顧客一覧", pages.Customers(cusmers)).Render(ctx, w)
	})

	// 顧客詳細
	app.mux.HandleFunc("GET /customers/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		customer, err := app.db.Queries.GetCustomerByID(ctx, int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		pages.CustomerEdit(customer).Render(ctx, w)
	})

	// 顧客作成
	app.mux.HandleFunc("POST /customers/create", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("顧客作成"))
	})

	// 顧客更新
	app.mux.HandleFunc("PATCH /customers/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("顧客更新"))
	})

	return nil
}
