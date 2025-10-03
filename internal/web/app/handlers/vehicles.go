package handlers

import (
	"net/http"
	"project/internal/web/app/views/layouts"
	"project/internal/web/app/views/pages"
	"strconv"
)

func (app *AppHandlers) VehiclesRoute() error {
	// 車両一覧
	app.mux.HandleFunc("GET /vehicles", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		vehicles, err := app.db.Queries.GetVehicles(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		layouts.Base("車両一覧", pages.Vehicles(vehicles)).Render(ctx, w)
	})
	// 車両詳細
	app.mux.HandleFunc("GET /vehicles/{id}", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		vehicle, err := app.db.Queries.GetVehicleByID(ctx, int32(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages.VehicleEdit(vehicle).Render(ctx, w)

	})

	// 車両作成
	app.mux.HandleFunc("POST /vehicles/create", func(w http.ResponseWriter, r *http.Request) {
		vehicles, err := app.db.Queries.GetVehicles(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages.Vehicles(vehicles).Render(r.Context(), w)
	})

	// 車両更新
	app.mux.HandleFunc("PATCH /vehicles/update/{id}", func(w http.ResponseWriter, r *http.Request) {
		vehicles, err := app.db.Queries.GetVehicles(r.Context())
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		pages.Vehicles(vehicles).Render(r.Context(), w)
	})

	return nil
}
