package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

// 削除
func (h *Handler) DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "文字を数値に変換できませんでした", http.StatusBadRequest)
		return
	}

	err = h.DB.Queries.DeleteVehicle(ctx, int32(idInt))
	if err != nil {
		http.Error(w, "車両の削除に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.RemoveElementByID(fmt.Sprintf("vehicle-%d", idInt))
}