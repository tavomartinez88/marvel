package helper

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/marvel/internal/dao"
	"github.com/tavomartinez88/marvel/internal/utils/helper"
	"github.com/tavomartinez88/marvel/pkg/models"
	"testing"
)

func TestBuildCollaborator(t *testing.T) {
	c := helper.BuildCollaborator(dao.CollaboratorsDao{
		ColID: "iron man",
		LastSync: "any-date",
		Colorists: []string{"hames"},
		Writers: []string{"tom"},
		Editors: []string{"Jack"},
	})

	assert.EqualValues(t, "any-date", c.LastSync)
	assert.EqualValues(t, true, len(c.Colorists)>0)
	assert.EqualValues(t, true, len(c.Writers)>0)
	assert.EqualValues(t, true, len(c.Editors)>0)
}

func TestBuildCharacters(t *testing.T) {
	c := helper.BuildCharacters(dao.CharacterDao{
		CharID: "iron man",
		LastSync: "any-date",
		Characters: []models.Character{
			{
				Character: "spider man",
				Comics: []string{"avengers"},
			},
		},
	})

	assert.EqualValues(t, "any-date", c.LastSync)
	assert.EqualValues(t, true, len(c.Characters)>0)
}

