package models

type CollaboratorsResponse struct {
	LastSync string `json:"last_sync"`
	Editors []string `json:"editors,omitempty"`
	Writers []string `json:"writers,omitempty"`
	Colorists []string `json:"colorists,omitempty"`
}
