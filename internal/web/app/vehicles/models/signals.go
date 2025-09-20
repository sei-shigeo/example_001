package models

type VehicleEditSignals struct {
	Edit Vehicle `json:"vehiclesEdit"`
}

type VehicleCreateSignals struct {
	Create Vehicle `json:"vehiclesCreate"`
}

// vehicle signals
type Vehicle struct {
	Number   string            `json:"number"`
	Errs     map[string]string `json:"errs"`
	Disabled bool              `json:"disabled"`
	Required bool              `json:"required"`
}
