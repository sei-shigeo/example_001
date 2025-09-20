package handler

import (
	"fmt"
	"net/http"
	"project/internal/db"
	"project/internal/web/app/vehicles/models"
	"project/internal/web/app/vehicles/template"
	"strconv"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

// Edit employeeform content
func (h *Handler) EditVehicleForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "文字を数値に変換できませんでした", http.StatusBadRequest)
		return
	}
	var sig models.VehicleEditSignals
	if err = datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	emp, err := h.DB.Queries.GetVehicleByID(ctx, int32(idInt))
	if err != nil {
		http.Error(w, "車両データの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	sig.Edit = models.Vehicle{
		Number: emp.Number,
	}

	sig.Edit.NewErrs()

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.EditVehicleForm(emp),
		datastar.WithSelectorID("editVehicleForm"),
		datastar.WithModeInner(),
	)
	sse.MarshalAndPatchSignals(sig)
	// 取得出来たら編集フォームを表示
	sse.PatchSignals([]byte("{editVehicleForm: true}"))
}

// 更新
func (h *Handler) UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "IDの取得に失敗しました", http.StatusBadRequest)
		return
	}

	var sig models.VehicleEditSignals
	if err = datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	// 更新パラメータを作成
	updateParams := &db.UpdateVehicleParams{
		ID:     int32(idInt),
		Number: sig.Edit.Number,
	}

	// 更新
	vehicle, err := h.DB.Queries.UpdateVehicle(ctx, updateParams)
	if err != nil {
		http.Error(w, "車両の更新に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.VehicleOne(vehicle),
		datastar.WithSelectorID(fmt.Sprintf("vehicle-%d", vehicle.ID)),
		datastar.WithModeReplace(),
	)
}
