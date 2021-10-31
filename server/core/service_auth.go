package core

import (
	net "github.com/atulanand206/go-network"
	"github.com/dgrijalva/jwt-go/v4"
)

type AuthService struct{}

func (service AuthService) GenerateTokens(user Player) (AuthenticationResponse, error) {
	var token AuthenticationResponse
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
func (service AuthService) AccessTokenClaims(user Player) (claims jwt.MapClaims) {
	claims = jwt.MapClaims{}
	claims["access"] = true
	claims["email"] = user.Email
	claims["userId"] = user.Id
	claims["name"] = user.Name
	return
}

// Generates claims to sign with the refresh token.
func (service AuthService) RefreshTokenClaims(user Player) (claims jwt.MapClaims) {
	claims = jwt.MapClaims{}
	claims["refresh"] = true
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["userId"] = user.Id
	return
}
