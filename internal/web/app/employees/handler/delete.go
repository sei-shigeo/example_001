package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

// 削除
func (h *Handler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "従業員の削除に失敗しました", http.StatusBadRequest)
		return
	}

	err = h.DB.Queries.DeleteEmployee(ctx, int32(idInt))
	if err != nil {
		http.Error(w, "従業員の削除に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.RemoveElementByID(fmt.Sprintf("employee-%d", idInt))
}