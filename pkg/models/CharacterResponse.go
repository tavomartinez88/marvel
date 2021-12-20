package models

type CharacterResponse struct {
	LastSync string        `json:"last_sync"`
	Characters []Character `json:"characters"`
}
