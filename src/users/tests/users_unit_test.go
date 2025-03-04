package tests

import (
	"github.com/stretchr/testify/assert"
	sw "soa/users/go"
	"strings"
	"testing"
)

func TestJWTToken(t *testing.T) {
	sw.InitAuthHandler()

	token, err := sw.CreateToken("test_user")
	assert.NoError(t, err)
	name, err := sw.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "test_user", name)
	name, err = sw.ValidateToken(token[1:])
	assert.Error(t, err)
}

func TestPasswordStrengthChecker(t *testing.T) {
	strong, errMsg := sw.CheckPasswordStrength("a")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must be at least 8 characters", errMsg)
	strong, errMsg = sw.CheckPasswordStrength(strings.Repeat("a", 33))
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must be less than 32 characters", errMsg)
	strong, errMsg = sw.CheckPasswordStrength("aAaA@@@@")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must contain at least 1 digit", errMsg)
	strong, errMsg = sw.CheckPasswordStrength("ab1b@@@@")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must contain at least 1 uppercase letter", errMsg)
	strong, errMsg = sw.CheckPasswordStrength("aAaAbBbB1")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must contain at least 1 special character", errMsg)
	strong, errMsg = sw.CheckPasswordStrength("aAaAbBb@1")
	assert.Equal(t, true, strong)
	assert.Equal(t, "", errMsg)
}
