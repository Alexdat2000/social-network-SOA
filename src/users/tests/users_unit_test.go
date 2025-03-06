package tests

import (
	"github.com/stretchr/testify/assert"
	api "soa/users/go"
	"strings"
	"testing"
)

func TestJWTToken(t *testing.T) {
	api.InitAuthHandler()

	token, err := api.CreateToken("test_user")
	assert.NoError(t, err)
	name, err := api.ValidateToken(token)
	assert.NoError(t, err)
	assert.Equal(t, "test_user", name)
	name, err = api.ValidateToken(token[1:])
	assert.Error(t, err)
}

func TestPasswordStrengthChecker(t *testing.T) {
	strong, errMsg := api.CheckPasswordStrength("a")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must be at least 8 characters", errMsg)
	strong, errMsg = api.CheckPasswordStrength(strings.Repeat("a", 33))
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must be less than 32 characters", errMsg)
	strong, errMsg = api.CheckPasswordStrength("aAaA@@@@")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must contain at least 1 digit", errMsg)
	strong, errMsg = api.CheckPasswordStrength("ab1b@@@@")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must contain at least 1 uppercase letter", errMsg)
	strong, errMsg = api.CheckPasswordStrength("aAaAbBbB1")
	assert.Equal(t, false, strong)
	assert.Equal(t, "Password must contain at least 1 special character", errMsg)
	strong, errMsg = api.CheckPasswordStrength("aAaAbBb@1")
	assert.Equal(t, true, strong)
	assert.Equal(t, "", errMsg)
}

func TestPasswordHasher(t *testing.T) {
	hash1 := api.HashPassword("alex", "password")
	hash2 := api.HashPassword("alex", "wordpass")
	assert.NotEqual(t, hash1, hash2)
	hash3 := api.HashPassword("alex", "password")
	assert.Equal(t, hash1, hash3)
}

func TestPhoneSanitizer(t *testing.T) {
	ok, phone := api.SanitizePhone("+7 (800) 555 - 35 - 35")
	assert.True(t, ok)
	assert.Equal(t, "78005553535", phone)
	ok, phone = api.SanitizePhone("+7(800)555-35-35")
	assert.True(t, ok)
	assert.Equal(t, "78005553535", phone)
	ok, phone = api.SanitizePhone("+7(800)555-35-ab")
	assert.False(t, ok)
	ok, phone = api.SanitizePhone("+7(8000)555-35-35-35")
	assert.False(t, ok)
}
