package services

import (
	"github.com/tavomartinez88/marvel/internal/dao"
	error2 "github.com/tavomartinez88/marvel/internal/error"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/pkg/models"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	WRITER = "writer"
	COLORIST = "colorist"
	EDITORS = "editor"
)

type ICollaboratorService interface {
	GetCollaborators(name string) (models.CollaboratorsResponse, error)
}

type CollaboratorService struct {
	Client utils.MarvelClient
	Dao dao.MarvelDao
}

func (cs *CollaboratorService) GetCollaborators(name string) (models.CollaboratorsResponse, error) {
	if !regexp.MustCompile(`^[ a-zA-Z]+$`).MatchString(name) {
		return models.CollaboratorsResponse{}, error2.ClientError{
			HttpStatus: http.StatusBadRequest,
			Message: "Name is invalid",
		}
	}

	// add check with db
	result, errDb := cs.Dao.GetCollaborators(name)

	if errDb != nil && errDb.Error() != "record not found"{ //there's a problem with db
		return models.CollaboratorsResponse{}, error2.ClientError{
			HttpStatus: http.StatusBadRequest,
			Message: errDb.Error(),
		}
	}

	if errDb != nil && errDb.Error() == "record not found"{// not found on db and i should go to marvel
		return cs.GetCollaboratorsFromMarvelAndSaveOnDb(name)
	}

	t, errParser := time.Parse("02-01-2006 15:04:05", result.LastSync) //i found collabprators andcheck last update
	if errParser != nil {
		return models.CollaboratorsResponse{}, error2.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message: errParser.Error(),
		}
	}

	lastSync, _ := strconv.ParseFloat(os.Getenv(utils.LimitSync), 10)
	if time.Now().Sub(t).Hours() > lastSync{
		return cs.GetCollaboratorsFromMarvelAndSaveOnDb(name)
	}
	return buildCollaborator(result), nil
}

func (cs *CollaboratorService) GetCollaboratorsFromMarvelAndSaveOnDb(name string) (models.CollaboratorsResponse, error) {
	var response models.CollaboratorsResponse
	id, err := cs.Client.GetHeroId(strings.ReplaceAll(name, " ", "%20"))
	if err != nil {
		return models.CollaboratorsResponse{}, err
	}

	collaborators, err := cs.Client.GetCollaboratorsByHeroId(id)

	if err != nil {
		return models.CollaboratorsResponse{}, err
	}

	response.Writers, response.Editors, response.Colorists = filterCollaboratorsByRole(collaborators)
	response.LastSync = time.Now().Format("02-01-2006 15:04:05")
	err = cs.Dao.SaveOrUpdateCollaborators(dao.CollaboratorsDao{
		ColID:     name,
		LastSync:  response.LastSync,
		Editors:   response.Editors,
		Writers:   response.Writers,
		Colorists: response.Colorists,
	})

	if err != nil {
		return models.CollaboratorsResponse{}, error2.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response, nil
}

func buildCollaborator(result dao.CollaboratorsDao) models.CollaboratorsResponse {
	var response models.CollaboratorsResponse
	response.LastSync = result.LastSync
	response.Colorists = result.Colorists
	response.Editors = result.Editors
	response.Writers = result.Writers
	return response
}

func filterCollaboratorsByRole(collaborators []utils.Collaborator) ([]string, []string, []string) {
	var writers, editors, colorists []string
	for _, collaborator := range collaborators {
		if collaborator.Role == WRITER {
			writers = append(writers, collaborator.Name)
		}

		if collaborator.Role == COLORIST {
			colorists = append(colorists, collaborator.Name)
		}

		if collaborator.Role == EDITORS {
			editors = append(editors, collaborator.Name)
		}
	}

	return writers, editors, colorists
}
