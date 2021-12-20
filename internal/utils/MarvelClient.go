package utils

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/sirupsen/logrus"
	errors "github.com/tavomartinez88/marvel/internal/error"
	"github.com/tavomartinez88/marvel/pkg/models"
	"net/http"
	"os"
	"strconv"
	"time"
)

type MarvelClient interface {
	GetHeroId(name string) (int64, error)
	GetCollaboratorsByHeroId(id int64) ([]Collaborator, error)
	GetCaractersByHeroId(id int64) ([]models.Character, error)
}

const characters = "characters/"

type MarvelClientImpl struct{}

func (cdi *MarvelClientImpl) GetCaractersByHeroId(id int64) ([]models.Character, error) {
	log := logrus.Entry{}
	var comicsByCharacter = new(Comics)
	var comicsByCharacterList []models.Character

	response,err := resty.New().R().Get(os.Getenv(BaseUrl)+ characters + strconv.FormatInt(id, 10) +"/comics?apikey="+os.Getenv(ApiKey)+"&hash="+ getHash()+"&ts=1")

	if err != nil {
		return nil, errors.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = json.Unmarshal(response.Body(), comicsByCharacter)

	if err != nil {
		log.Errorf("Error unmarshalled, %v", err)
	}

	if comicsByCharacter.Data.Count == 0 {
		return nil, errors.ClientError{
			HttpStatus: http.StatusNotFound,
			Message:    "comics not found",
		}
	}

	for _, comic := range comicsByCharacter.Data.Results {
		for _, hero := range comic.Characters.Items {
			indexHero := getIndexHero(comicsByCharacterList, hero.Name)
			if indexHero < 0 { //not found
				c := models.Character{}
				c.Character = hero.Name
				c.Comics = []string{comic.Title}
				comicsByCharacterList = append(comicsByCharacterList, c)
			}else {
				comicsByCharacterList[indexHero].Comics = append(comicsByCharacterList[indexHero].Comics, comic.Title)
			}
		}
	}

	return comicsByCharacterList, nil
}

func (cdi *MarvelClientImpl) GetCollaboratorsByHeroId(id int64) ([]Collaborator, error){
	log := logrus.Entry{}
	var comicsByCharacter = new(Comics)
	var collaborators []Collaborator
	response,err := resty.New().R().Get(os.Getenv(BaseUrl)+ characters + strconv.FormatInt(id, 10) +"/comics?apikey="+os.Getenv(ApiKey)+"&hash="+ getHash()+"&ts=1")

	if err != nil {
		return nil, errors.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = json.Unmarshal(response.Body(), comicsByCharacter)

	if err != nil {
		log.Errorf("Error unmarshalled, %v", err)
	}

	if comicsByCharacter.Data.Count == 0 {
		return nil, errors.ClientError{
			HttpStatus: http.StatusNotFound,
			Message:    "comics not found",
		}
	}

	for _, comic := range comicsByCharacter.Data.Results {
		for _, item := range comic.Collaborators.Items {
			collaborators = append(collaborators, item)
		}
	}

	return collaborators, nil
}

func (cdi *MarvelClientImpl) GetHeroId(name string) (int64, error) {
	log := logrus.Entry{}
	var hero = new(Hero)
	response,err := resty.New().R().Get(os.Getenv(BaseUrl)+"characters?apikey="+os.Getenv(ApiKey)+"&hash="+ getHash()+"&name="+name+"&ts=1")

	if err != nil {
		return 0, errors.ClientError{
			HttpStatus: http.StatusInternalServerError,
			Message:    err.Error(),
		}
	}

	err = json.Unmarshal(response.Body(), hero)

	if err != nil {
		log.Warn("Error unmarshalled, %v", err)
	}

	if hero.Data.Count == 0 {
		return 0, errors.ClientError{
			HttpStatus: http.StatusNotFound,
			Message:    "Character not found",
		}
	}

	return hero.Data.Results[0].Id, nil
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getHash() string {
	_ = time.Now().Unix()
	return getMD5Hash(fmt.Sprintf("%d%s%s", 1, os.Getenv(PrivateApiKey), os.Getenv(ApiKey)))
}

func getIndexHero(heroes []models.Character, heroName string) int {
	for index, hero := range heroes {
		if hero.Character == heroName {
			return index
		}
	}

	return -1 //value if i couldn't element
}