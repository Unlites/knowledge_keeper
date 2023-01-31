package auth

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenManager interface {
	NewAccessToken(userId string) (string, error)
	NewRefreshToken() (*RefreshToken, error)
	ParseUserIdFromAccessToken(accessToken string) (string, error)
}

type TokenManagerSettings struct {
	SigningKey      string
	AccessTokenTTL  time.Duration
	RefreshTokenTTL time.Duration
}

type RefreshToken struct {
	Token     string
	ExpiresAt int64
}

type manager struct {
	settings *TokenManagerSettings
}

var _ TokenManager = (*manager)(nil)

func NewTokenManager(settings *TokenManagerSettings) *manager {
	return &manager{settings: settings}
}

func (m *manager) NewAccessToken(userId string) (string, error) {
	expiresAt := time.Now().Add(m.settings.AccessTokenTTL).Unix()
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Subject:   userId,
		ExpiresAt: expiresAt,
	})

	accessTokenString, err := accessToken.SignedString([]byte(m.settings.SigningKey))
	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func (m *manager) NewRefreshToken() (*RefreshToken, error) {
	bytes := make([]byte, 32)
	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	_, err := r.Read(bytes)
	if err != nil {
		return nil, err
	}

	return &RefreshToken{
		Token:     fmt.Sprintf("%x", bytes),
		ExpiresAt: time.Now().Add(m.settings.RefreshTokenTTL).Unix(),
	}, nil
}

func (m *manager) ParseUserIdFromAccessToken(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %w", token.Header["alg"])
		}

		return []byte(m.settings.SigningKey), nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("failed to get user from claims")
	}

	return claims["sub"].(string), nil
}
