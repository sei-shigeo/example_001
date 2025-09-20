package handler

import (
	"net/http"
	"project/internal/web/app/vehicles/models"
	"project/internal/web/app/vehicles/template"

	"github.com/starfederation/datastar-go/datastar"
)

func (h *Handler) CreateVehicleForm(w http.ResponseWriter, r *http.Request) {
	var sig models.VehicleCreateSignals
	if err := datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	sig.Create.NewErrs()

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.CreateVehicleForm(),
		datastar.WithSelectorID("createVehicleForm"),
		datastar.WithModeInner(),
	)
	sse.MarshalAndPatchSignals(sig)
}

// 作成
func (h *Handler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var sig models.VehicleCreateSignals
	if err := datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	// バリデーションエラーがある時は送信できないようにする
	if !sig.Create.IsValid() {
		http.Error(w, "バリデーションエラーがあります", http.StatusBadRequest)
		return
	}

	vehicle, err := h.DB.Queries.CreateVehicle(ctx, sig.Create.Number)
	if err != nil {
		http.Error(w, "車両の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)

	sse.PatchElementTempl(
		template.VehicleOne(vehicle),
		datastar.WithSelectorID("vehiclesList"),
		datastar.WithModeAppend(),
	)

	sig.Create = models.Vehicle{}
	sig.Create.NewErrs()
	sse.MarshalAndPatchSignals(sig)
}
