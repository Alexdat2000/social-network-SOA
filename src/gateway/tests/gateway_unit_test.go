package tests

import (
	"github.com/stretchr/testify/assert"
	"soa/gateway/utils"
	"testing"
)

func TestParsePostId(t *testing.T) {
	id, err := utils.ParsePostId("10")
	assert.Equal(t, uint32(10), id)
	assert.NoError(t, err)
	id, err = utils.ParsePostId("-10")
	assert.Error(t, err)
	id, err = utils.ParsePostId("post1")
	assert.Error(t, err)
	id, err = utils.ParsePostId("?**")
	assert.Error(t, err)
	id, err = utils.ParsePostId("0")
	assert.Equal(t, uint32(0), id)
	assert.NoError(t, err)
}

func TestParsePostPrivate(t *testing.T) {
	assert.True(t, utils.ParsePostPrivate("true"))
	assert.True(t, utils.ParsePostPrivate("tRue"))
	assert.True(t, utils.ParsePostPrivate("private"))
	assert.True(t, utils.ParsePostPrivate("TRUE"))
	assert.True(t, utils.ParsePostPrivate("PriVate"))
	assert.False(t, utils.ParsePostPrivate(""))
	assert.False(t, utils.ParsePostPrivate("no"))
	assert.False(t, utils.ParsePostPrivate("false"))
	assert.False(t, utils.ParsePostPrivate("kikf9e3k90i34m9"))
}
