package helper

import (
	"github.com/tavomartinez88/marvel/internal/dao"
	"github.com/tavomartinez88/marvel/pkg/models"
)

func BuildCollaborator(result dao.CollaboratorsDao) models.CollaboratorsResponse {
	var response models.CollaboratorsResponse
	response.LastSync = result.LastSync
	response.Colorists = result.Colorists
	response.Editors = result.Editors
	response.Writers = result.Writers
	return response
}

func BuildCharacters(result dao.CharacterDao) models.CharacterResponse {
	return models.CharacterResponse{
		LastSync:   result.LastSync,
		Characters: result.Characters,
	}
}