package services

import (
	"github.com/tavomartinez88/marvel/internal/dao"
	error2 "github.com/tavomartinez88/marvel/internal/error"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/internal/utils/helper"
	"github.com/tavomartinez88/marvel/pkg/models"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ICharacterService interface {
	GetCharacters(name string) (models.CharacterResponse, error)
}

type CharacterService struct {
	Client utils.MarvelClient
	Dao dao.MarvelDao
}

func (cs *CharacterService) GetCharacters(name string) (models.CharacterResponse, error) {
	if !regexp.MustCompile(`^[ a-zA-Z]+$`).MatchString(name) {
		return models.CharacterResponse{}, error2.ClientError{
			HttpStatus: http.StatusBadRequest,
			Message: "Name is invalid",
		}
	}

	// add check with db
	result, errDb := cs.Dao.GetCharacters(name)

	if errDb != nil && errDb.Error() != "record not found"{ //there's a problem with db
		return models.CharacterResponse{}, error2.ClientError{
			HttpStatus: http.StatusBadRequest,
			Message: errDb.Error(),
		}
	}

	if errDb != nil && errDb.Error() == "record not found"{// not found on db and i should go to marvel
		return cs.GetCharactersFromMarvelAndSaveOnDb(name)
	}

	t, errParser := time.Parse("02-01-2006 15:04:05", result.LastSync) //i found collabprators andcheck last update
	if errParser != nil {
		return models.CharacterResponse{}, error2.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message: errParser.Error(),
		}
	}

	lastSync, _ := strconv.ParseFloat(os.Getenv(utils.LimitSync), 10)
	if time.Now().Sub(t).Hours() > lastSync{ //hay mas de x de la ultima actualizacion hay que sincronizar la bd
		return cs.GetCharactersFromMarvelAndSaveOnDb(name)
	}
	return helper.BuildCharacters(result), nil
}

func (cs *CharacterService) GetCharactersFromMarvelAndSaveOnDb(name string) (models.CharacterResponse, error) {
	var response models.CharacterResponse
	id, err := cs.Client.GetHeroId(strings.ReplaceAll(name, " ", "%20"))
	if err != nil {
		return models.CharacterResponse{}, err
	}

	characters, err := cs.Client.GetCaractersByHeroId(id)

	if err != nil {
		return models.CharacterResponse{}, err
	}

	response.LastSync = time.Now().Format("02-01-2006 15:04:05")
	response.Characters = characters
	err = cs.Dao.SaveOrUpdateCharacters(dao.CharacterDao{
		CharID:     name,
		LastSync:  response.LastSync,
		Characters: characters,
	})

	if err != nil {
		return models.CharacterResponse{}, error2.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	return response, nil
}