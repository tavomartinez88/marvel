package services

import (
	"github.com/stretchr/testify/mock"
	"github.com/tavomartinez88/marvel/internal/dao"
	"github.com/tavomartinez88/marvel/internal/utils"
	"github.com/tavomartinez88/marvel/pkg/models"
)

type ClientMock struct {
	mock.Mock
}

type DaoMock struct {
	mock.Mock
}


//client Marvel

func (c *ClientMock) GetHeroId(name string) (int64, error) {
	args := c.Called()
	return args.Get(0).(int64), args.Error(1)
}

func (c *ClientMock) GetCollaboratorsByHeroId(id int64) ([]utils.Collaborator, error) {
	args := c.Called()
	return args.Get(0).([]utils.Collaborator), args.Error(1)
}

func (c *ClientMock) GetCaractersByHeroId(id int64) ([]models.Character, error) {
	args := c.Called()
	return args.Get(0).([]models.Character), args.Error(1)
}

//dao Marvel

func (d *DaoMock) GetCollaborators(name string) (dao.CollaboratorsDao, error) {
	args := d.Called()
	return args.Get(0).(dao.CollaboratorsDao), args.Error(1)
}

func (d *DaoMock) GetCharacters(name string) (dao.CharacterDao, error) {
	args := d.Called()
	return args.Get(0).(dao.CharacterDao), args.Error(1)
}

func (d *DaoMock) SaveOrUpdateCharacters(c dao.CharacterDao) error {
	args := d.Called()
	return args.Error(0)
}

func (d *DaoMock) SaveOrUpdateCollaborators(c dao.CollaboratorsDao) error {
	args := d.Called()
	return args.Error(0)
}

