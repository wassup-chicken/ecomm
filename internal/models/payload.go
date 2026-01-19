package models

type JSONPayload struct {
	Error bool `json:"error"`
	Data  any  `json:"data,omitempty"`
}
