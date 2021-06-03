package entities

type WithHeld struct {
	Scope        string   `json:"scope"`
	CountryCodes []string `json:"country_codes"`
}
