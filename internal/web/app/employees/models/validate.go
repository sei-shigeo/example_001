package models

import "unicode"

func (f *Employee) NewErrs() {
	f.Errs = map[string]string{
		"name":  "",
		"email": "",
	}
}

// Validate はフォームのバリデーションを行います
func (f *Employee) Validate() {
	f.NewErrs()

	if f.Name == "" {
		f.Errs["name"] = "名前は必須です"
	} else if containsSpace(f.Name) {
		f.Errs["name"] = "空白は使用できません"
	} else if !isJapaneseOnly(f.Name) {
		f.Errs["name"] = "日本語のみ入力してください"
	}

	if f.Email == "" {
		f.Errs["email"] = "メールアドレスは必須です"
	} else if !isValidEmail(f.Email) {
		f.Errs["email"] = "有効なメールアドレスを入力してください"
	}
}

// errsが空かどうかをチェックします
func (f *Employee) IsValid() bool {
	for _, err := range f.Errs {
		if err != "" {
			return false
		}
	}
	return true
}

// isJapaneseOnly は日本語のみかどうかをチェックします
func isJapaneseOnly(s string) bool {
	// ひらがな、カタカナ、漢字のみを許可
	for _, r := range s {
		if !((r >= 0x3040 && r <= 0x309F) || // ひらがな
			(r >= 0x30A0 && r <= 0x30FF) || // カタカナ
			(r >= 0x4E00 && r <= 0x9FAF)) { // 漢字
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

// isValidEmail はメールアドレスの形式をチェックします
func isValidEmail(email string) bool {
	if len(email) < 5 || len(email) > 254 {
		return false
	}

	// @の位置と数をチェック
	atIndex := -1
	atCount := 0
	for i, r := range email {
		if r == '@' {
			atCount++
			atIndex = i
		}
	}

	if atCount != 1 || atIndex == 0 || atIndex == len(email)-1 {
		return false
	}

	localPart := email[:atIndex]
	domainPart := email[atIndex+1:]

	// ローカル部の検証
	if len(localPart) == 0 || len(localPart) > 64 {
		return false
	}

	// ドメイン部の検証
	if len(domainPart) == 0 || len(domainPart) > 253 {
		return false
	}

	// ドットの存在チェック
	hasDot := false
	for _, r := range domainPart {
		if r == '.' {
			hasDot = true
			break
		}
	}

	return hasDot
}
