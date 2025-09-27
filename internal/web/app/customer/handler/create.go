package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/customer/models"
	"project/internal/web/app/customer/template"

	"github.com/starfederation/datastar-go/datastar"
)

func (h *Handler) CreateCustomerForm(w http.ResponseWriter, r *http.Request) {
	var sig models.CustomerCreateSignals
	if err := datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	sig.Create.NewErrs()

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.CreateCustomer(),
		datastar.WithSelectorID("createCustomerForm"),
		datastar.WithModeInner(),
	)
	sse.MarshalAndPatchSignals(sig)
}

func (h *Handler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var sig models.CustomerCreateSignals
	if err := datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	if !sig.Create.IsValid() {
		http.Error(w, "バリデーションエラーがあります", http.StatusBadRequest)
		return
	}

	customer, err := h.DB.Queries.CreateCustomer(ctx, &db.CreateCustomerParams{
		Name:  sig.Create.Name,
		Email: sig.Create.Email,
	})
	if err != nil {
		http.Error(w, "顧客の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.CustomerOne(customer),
		datastar.WithSelectorID("customersList"),
		datastar.WithModeAppend(),
	)

	sig.Create = models.Customer{}
	sig.Create.NewErrs()
	sse.MarshalAndPatchSignals(sig)
}
