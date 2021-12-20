package dao

import "github.com/tavomartinez88/marvel/pkg/models"

type CharacterDao struct {
	CharID string `json:"charid"`
	LastSync string        `json:"last_sync"`
	Characters []models.Character `json:"characters"`
}

func (c CharacterDao) ID() (jsonField string, value interface{}) {
	value=c.CharID
	jsonField="charid"
	return
}