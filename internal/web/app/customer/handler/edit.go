package handler

import (
	"fmt"
	"net/http"
	"project/internal/db"
	"project/internal/web/app/customer/models"
	"project/internal/web/app/customer/template"
	"strconv"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

func (h *Handler) EditCustomerForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "文字を数値に変換できませんでした", http.StatusBadRequest)
		return
	}
	var sig models.CustomerEditSignals
	if err = datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}
	customer, err := h.DB.Queries.GetCustomerByID(ctx, int32(idInt))
	if err != nil {
		http.Error(w, "顧客データの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	sig.Edit = models.Customer{
		Name:  customer.Name,
		Email: customer.Email,
	}

	sig.Edit.NewErrs()

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.EditCustomerForm(customer),
		datastar.WithSelectorID("editCustomerForm"),
		datastar.WithModeInner(),
	)
	sse.MarshalAndPatchSignals(sig)
	// 取得出来たら編集フォームを表示
	sse.PatchSignals([]byte("{editCustomerForm: true}"))
}

func (h *Handler) UpdateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "文字を数値に変換できませんでした", http.StatusBadRequest)
		return
	}

	var sig models.CustomerEditSignals
	if err = datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	updateParams := &db.UpdateCustomerParams{
		ID:    int32(idInt),
		Name:  sig.Edit.Name,
		Email: sig.Edit.Email,
	}

	customer, err := h.DB.Queries.UpdateCustomer(ctx, updateParams)
	if err != nil {
		http.Error(w, "顧客の更新に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.CustomerOne(customer),
		datastar.WithSelectorID(fmt.Sprintf("customer-%d", customer.ID)),
		datastar.WithModeReplace(),
	)
}
