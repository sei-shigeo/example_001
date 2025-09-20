package handler

import (
	"fmt"
	"net/http"
	"project/internal/db"
	"project/internal/web/app/employees/models"
	"project/internal/web/app/employees/template"
	"strconv"
	"strings"

	"github.com/starfederation/datastar-go/datastar"
)

// Edit employeeform content
func (h *Handler) EditEmployeeForm(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "文字を数値に変換できませんでした", http.StatusBadRequest)
		return
	}
	var sig models.EmployeeEditSignals
	if err = datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	emp, err := h.DB.Queries.GetEmployeeByID(ctx, int32(idInt))
	if err != nil {
		http.Error(w, "従業員データの取得に失敗しました", http.StatusInternalServerError)
		return
	}

	sig.Edit = models.Employee{
		Name:  emp.Name,
		Email: emp.Email,
	}

	sig.Edit.NewErrs()

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.EditEmployeeForm(emp),
		datastar.WithSelectorID("editEmployeeForm"),
		datastar.WithModeInner(),
	)
	sse.MarshalAndPatchSignals(sig)
	// 取得出来たら編集フォームを表示
	sse.PatchSignals([]byte("{editEmployeeForm: true}"))
}

// 更新
func (h *Handler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := strings.TrimSpace(r.PathValue("id"))
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "IDの取得に失敗しました", http.StatusBadRequest)
		return
	}

	var sig models.EmployeeEditSignals
	if err = datastar.ReadSignals(r, &sig); err != nil {
		http.Error(w, "signalsの読み込みに失敗しました", http.StatusInternalServerError)
		return
	}

	// 更新パラメータを作成
	updateParams := &db.UpdateEmployeeParams{
		ID:    int32(idInt),
		Name:  sig.Edit.Name,
		Email: sig.Edit.Email,
	}

	// 更新
	emp, err := h.DB.Queries.UpdateEmployee(ctx, updateParams)
	if err != nil {
		http.Error(w, "従業員の更新に失敗しました", http.StatusInternalServerError)
		return
	}

	sse := datastar.NewSSE(w, r)
	sse.PatchElementTempl(
		template.EmployeeOne(emp),
		datastar.WithSelectorID(fmt.Sprintf("employee-%d", emp.ID)),
		datastar.WithModeReplace(),
	)
}
