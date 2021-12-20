package dao

import (
	"github.com/stretchr/testify/assert"
	"github.com/tavomartinez88/marvel/internal/dao"
	"testing"
)

func TestGetHeroId(t *testing.T) {
	data := &dao.CharacterDaoImpl{}

	id, err := data.GetHeroId("iron+man")
	assert.Nil(t, err)
	assert.True(t, true, id > 0)
}
