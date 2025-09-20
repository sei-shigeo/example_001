package models

type EmployeeEditSignals struct {
	Edit Employee `json:"employeeEdit"`
}

type EmployeeCreateSignals struct {
	Create Employee `json:"employeeCreate"`
}

type EmployeeValidationSignals struct {
	Errs map[string]string `json:"errs"`
}

// employee signals
type Employee struct {
	Name     string            `json:"name"`
	Email    string            `json:"email"`
	Errs     map[string]string `json:"errs"`
	Disabled bool              `json:"disabled"`
	Required bool              `json:"required"`
}
