package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/layouts"
	"project/internal/web/app/vehicles/template"
)

type Handler struct {
	DB *db.DB
}

func NewHandler(database *db.DB) *Handler {
	return &Handler{
		DB: database,
	}
}

func (h *Handler) GetVehicles(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vehicles, err := h.DB.Queries.GetVehicles(ctx)
	if err != nil {
		http.Error(w, "車両データの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	title := "Vehicles"

	props := template.Props{
		Vehicles: vehicles,
	}

	layouts.Base(title, props.VehiclesPage()).Render(ctx, w)
}
