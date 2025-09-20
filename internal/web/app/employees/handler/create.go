package handler

import (
	"net/http"
	"project/internal/db"
	"project/internal/web/app/employees/models"
	"project/internal/web/app/employees/template"

	"github.com/starfederation/datastar-go/datastar"
)

func (h *Handler) CreateEmployeeForm(w http.ResponseWriter, r *http.Request) {
	var empSignals models.EmployeeCreateSignals
	if err := datastar.ReadSignals(r, &empSignals); err != nil {
		http.Error(w, "従業員の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	empSignals.Create.NewErrs()

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.CreateEmployee(),
		datastar.WithSelectorID("createEmployeeForm"),
		datastar.WithModeInner(),
	)
	sse.MarshalAndPatchSignals(empSignals)
}

// 作成
func (h *Handler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var empSignals models.EmployeeCreateSignals
	if err := datastar.ReadSignals(r, &empSignals); err != nil {
		http.Error(w, "従業員の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	// バリデーションエラーがある時は送信できないようにする
	if !empSignals.Create.IsValid() {
		http.Error(w, "バリデーションエラーがあります", http.StatusBadRequest)
		return
	}

	emp, err := h.DB.Queries.CreateEmployee(ctx, &db.CreateEmployeeParams{
		Name:  empSignals.Create.Name,
		Email: empSignals.Create.Email,
	})
	if err != nil {
		http.Error(w, "従業員の作成に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)

	sse.PatchElementTempl(
		template.EmployeeOne(emp),
		datastar.WithSelectorID("employeesList"),
		datastar.WithModeAppend(),
	)

	var sig models.EmployeeCreateSignals
	sig.Create = models.Employee{}
	sig.Create.NewErrs()
	sse.MarshalAndPatchSignals(sig)
}
