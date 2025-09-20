package models

import "unicode"

func (f *Vehicle) NewErrs() {
	f.Errs = map[string]string{
		"number":  "",
	}
}

// Validate はフォームのバリデーションを行います
func (f *Vehicle) Validate() {
	f.NewErrs()

	if f.Number == "" {
		f.Errs["number"] = "車両番号は必須です"
	} else if containsSpace(f.Number) {
		f.Errs["number"] = "空白は使用できません"
	}

}

// errsが空かどうかをチェックします
func (f *Vehicle) IsValid() bool {
	for _, err := range f.Errs {
		if err != "" {
			return false
		}
	}
	return true
}

// containsSpace は空白文字が含まれているかチェックします
func containsSpace(s string) bool {
	for _, r := range s {
		if unicode.IsSpace(r) {
			return true
		}
	}
	return false
}
