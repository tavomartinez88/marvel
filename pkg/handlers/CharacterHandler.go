package handlers

import (
	"github.com/tavomartinez88/marvel/internal/dao"
	"github.com/tavomartinez88/marvel/internal/services"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/pkg/models"
)

type ICharacterHandler interface {
	GetCharacters(name string) (*models.CharacterResponse, error)
}

type characterHandler struct {
	Service services.ICharacterService
}

func (cs *characterHandler) GetCharacters(name string) (models.CharacterResponse, error) {
	return cs.Service.GetCharacters(name)
}

func NewCharacterHandler() *characterHandler {
	return &characterHandler{
		Service: &services.CharacterService{
			Client: &utils.MarvelClientImpl{},
			Dao: &dao.MarvelDaoImpl{},
		},
	}
}
