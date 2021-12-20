package handlers

import (
	"github.com/tavomartinez88/marvel/internal/dao"
	"github.com/tavomartinez88/marvel/internal/services"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/pkg/models"
)

type ICollaboratorHandler interface {
	GetCollaborators(name string) (models.CollaboratorsResponse, error)
}

type collaboratorHandler struct {
	Service services.ICollaboratorService
}

func (cs *collaboratorHandler) GetCollaborators(name string) (models.CollaboratorsResponse, error) {
	return cs.Service.GetCollaborators(name)
}

func NewCollaboratorHandler() *collaboratorHandler {
	return &collaboratorHandler{
		Service: &services.CollaboratorService{
			Client: &utils.MarvelClientImpl{},
			Dao: &dao.MarvelDaoImpl{},
		},
	}
}

