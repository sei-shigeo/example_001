package handler

import (
	"net/http"
	"project/internal/web/app/vehicles/models"

	"github.com/starfederation/datastar-go/datastar"
)

func (h *Handler) ValidateVehicle(w http.ResponseWriter, r *http.Request) {

	mode := r.URL.Query().Get("mode")
	if mode == "" {
		http.Error(w, "バリデーションに失敗しました", http.StatusBadRequest)
		return
	}

	if mode == "create" {
		var sig models.VehicleCreateSignals
		if err := datastar.ReadSignals(r, &sig); err != nil {
			http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
			return
		}

		sig.Create.Validate()

		sig.Create.Disabled = sig.Create.IsValid()

		sse := datastar.NewSSE(w, r)
		sse.MarshalAndPatchSignals(sig)
	}

	if mode == "edit" {
		var sig models.VehicleEditSignals
		if err := datastar.ReadSignals(r, &sig); err != nil {
			http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
			return
		}

		sig.Edit.Validate()

		sig.Edit.Disabled = sig.Edit.IsValid()

		sse := datastar.NewSSE(w, r)
		sse.MarshalAndPatchSignals(sig)
	}

}
