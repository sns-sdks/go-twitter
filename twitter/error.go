package twitter

import "fmt"

// APIError Errors for twitter response or library error
type APIError struct {
	ClientID           string `json:"client_id,omitempty"`
	RequiredEnrollment string `json:"required_enrollment,omitempty"`
	RegistrationUrl    string `json:"registration_url,omitempty"`
	Title              string `json:"title"`
	Detail             string `json:"detail"`
	Reason             string `json:"reason,omitempty"`
	Type               string `json:"type,omitempty"`
	Status             int    `json:"status,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Errors: %s %s", e.Title, e.Detail)
}
