package dao

import (
	db "github.com/sonyarouje/simdb"
	"sync"
)

type MarvelDao interface {
	GetCollaborators(name string) (CollaboratorsDao, error)
	GetCharacters(name string) (CharacterDao, error)
	SaveOrUpdateCharacters(c CharacterDao) error
	SaveOrUpdateCollaborators(c CollaboratorsDao) error
}

type MarvelDaoImpl struct{}

var driver *db.Driver
var once sync.Once
func init() {
	once.Do(func() {
		driver, _ = db.New("marvel-db")
	})
}

func (mdi *MarvelDaoImpl) GetCollaborators(name string) (CollaboratorsDao, error) {
	var collaborators []CollaboratorsDao
	err := driver.Open(CollaboratorsDao{}).Where("colid","=",name).Get().AsEntity(&collaborators)

	if len(collaborators)>0 {
		return collaborators[0], err
	}

	return CollaboratorsDao{}, err //access to fisrt result because always i'll one items on db
}

func (mdi *MarvelDaoImpl) GetCharacters(name string) (CharacterDao, error) {
	var characters []CharacterDao
	err := driver.Open(CharacterDao{}).Where("charid","=",name).Get().AsEntity(&characters)

	if len(characters)>0 {
		return characters[0], err
	}

	return CharacterDao{}, err //access to fisrt result because always i'll one items on db
}

func (mdi *MarvelDaoImpl) SaveOrUpdateCharacters(c CharacterDao) error {
	return driver.Upsert(c)
}

func (mdi *MarvelDaoImpl) SaveOrUpdateCollaborators(c CollaboratorsDao) error {
	return driver.Upsert(c)
}

