package handlers

import (
	"github.com/tavomartinez88/marvel/internal/dao"
	"github.com/tavomartinez88/marvel/internal/services"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/pkg/models"
)

type ICharacterHandler interface {
	GetCharacters(name string) (models.CharacterResponse, error)
}

type CharacterHandler struct {
	Service services.ICharacterService
}

func (cs *CharacterHandler) GetCharacters(name string) (models.CharacterResponse, error) {
	return cs.Service.GetCharacters(name)
}

func NewCharacterHandler() *CharacterHandler {
	return &CharacterHandler{
		Service: &services.CharacterService{
			Client: &utils.MarvelClientImpl{},
			Dao: &dao.MarvelDaoImpl{},
		},
	}
}
