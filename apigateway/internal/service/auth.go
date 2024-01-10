package service

import (
	"apigateway/internal/container"
	"apigateway/internal/model"
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type IAuthService interface {
	VerifyToken(ctx context.Context, token string) (*model.AuthClaims, error)
	GenerateToken(ctx context.Context, userClaims *model.AuthClaims) (string, error)
}

type authService struct {
	c container.IContainer
}

func NewAuthService() *authService {
	return &authService{}
}

func (s *authService) VerifyToken(ctx context.Context, token string) (*model.AuthClaims, error) {
	tkn, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return s.c.GetConfig().JWT.SecretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !tkn.Valid {
		return nil, errors.New("invalid token: not valid")
	}

	if claims, ok := tkn.Claims.(jwt.MapClaims); ok {
		userId, err := uuid.Parse(claims["user_id"].(string))
		if err != nil {
			return nil, err
		}
		return &model.AuthClaims{
			UserId:   userId,
			Username: claims["username"].(string),
			Email:    claims["email"].(string),
		}, nil
	}

	return nil, errors.New("invalid token: claims")
}

func (s *authService) GenerateToken(ctx context.Context, userClaims *model.AuthClaims) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour)
	claims["user_id"] = userClaims.UserId.String()
	claims["username"] = userClaims.Username
	claims["email"] = userClaims.Email
	tokenString, err := token.SignedString(s.c.GetConfig().JWT.SecretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
