package services

import (
	"github.com/atulanand206/find/server/core/models"
	net "github.com/atulanand206/go-network"
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthService struct{}

func (service AuthService) GenerateTokens(user models.Player) (models.AuthenticationResponse, error) {
	var token models.AuthenticationResponse
	// Generate a new access token for the username
	accessToken, err := net.CreateAccessToken(service.AccessTokenClaims(user))
	if err != nil {
		return token, err
	}
	// Generate a new refresh token for the username
	refreshToken, err := net.CreateRefreshToken(service.RefreshTokenClaims(user))
	if err != nil {
		return token, err
	}
	// Create the tokens object.
	token.AccessToken = accessToken
	token.RefreshToken = refreshToken
	return token, nil
}

// Generates claims to sign with the access token.
func (service AuthService) AccessTokenClaims(user models.Player) (claims jwt.MapClaims) {
	claims = jwt.MapClaims{}
	claims["access"] = true
	claims["email"] = user.Email
	claims["userId"] = user.Id
	claims["name"] = user.Name
	return
}

// Generates claims to sign with the refresh token.
func (service AuthService) RefreshTokenClaims(user models.Player) (claims jwt.MapClaims) {
	claims = jwt.MapClaims{}
	claims["refresh"] = true
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["userId"] = user.Id
	return
}
