package helper

import (
	"github.com/golang-jwt/jwt"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"go-wishlist-api-2/entities"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	viper.Set("SECRET_TOKEN", "secret")
	user := &entities.User{
		Id:    1,
		Email: "admin@gmail.com",
	}
	token, err := GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	claims := JWTClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	assert.NoError(t, err)
	assert.True(t, parsedToken.Valid)
	assert.Equal(t, user.Id, claims.Id)
	assert.Equal(t, user.Email, claims.Email)
	assert.True(t, claims.ExpiresAt > time.Now().Unix())
}
