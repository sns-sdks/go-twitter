package twitter

import "fmt"

// APIError Errors for twitter response or library error
type APIError struct {
	ClientID           string      `json:"client_id,omitempty"`
	RequiredEnrollment string      `json:"required_enrollment,omitempty"`
	RegistrationUrl    string      `json:"registration_url,omitempty"`
	Title              string      `json:"title"`
	Detail             string      `json:"detail"`
	Reason             string      `json:"reason,omitempty"`
	Type               string      `json:"type,omitempty"`
	Status             int         `json:"status,omitempty"`
	Errors             interface{} `json:"errors,omitempty"`
}

func (e APIError) Error() string {
	return fmt.Sprintf("Twitter Error, Title: %s Detail: %s Errors: %s", e.Title, e.Detail, e.Errors)
}

func (e APIError) String() string {
	return Stringify(e)
}
