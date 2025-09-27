package models

type CustomerEditSignals struct {
	Edit Customer `json:"customerEdit"`
}

type CustomerCreateSignals struct {
	Create Customer `json:"customerCreate"`
}

type CustomerValidationSignals struct {
	Errs map[string]string `json:"errs"`
}

// customer signals
type Customer struct {
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Errs     map[string]string `json:"errs"`
	Disabled bool              `json:"disabled"`
	Required bool              `json:"required"`
}