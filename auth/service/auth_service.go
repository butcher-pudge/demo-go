package service

import (
	"context"
	"crypto/sha1"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/spf13/viper"
	authErrors "go-learning-demo/auth/error"
	"go-learning-demo/auth/repository"
	"go-learning-demo/models"
	"time"
)

type AuthClaims struct {
	jwt.StandardClaims
	User *models.User `json:"user"`
}

type AuthService struct {
	userRepository repository.IUserRepository
}

func NewAuthService(userRepository repository.IUserRepository) *AuthService {
	return &AuthService{
		userRepository: userRepository,
	}
}

func (service *AuthService) SignUp(ctx context.Context, username, password string) error {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))

	user := &models.User{
		UserName: username,
		Password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}

	return service.userRepository.CreateUser(ctx, user)
}

func (service *AuthService) SignIn(ctx context.Context, username, password string) (string, error) {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(viper.GetString("auth.hash_salt")))
	password = fmt.Sprintf("%x", pwd.Sum(nil))

	user, err := service.userRepository.GetUser(ctx, username, password)
	if err != nil {
		return "", authErrors.ErrUserNotFound
	}

	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Second * viper.GetDuration("auth.token_ttl"))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(viper.GetString("auth.signing_key")))
}

func (service *AuthService) ParseToken(_ context.Context, accessToken string) (*models.User, error) {
	token, err := jwt.ParseWithClaims(accessToken, &AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(viper.GetString("auth.signing_key")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*AuthClaims); ok && token.Valid {
		return claims.User, nil
	}

	return nil, authErrors.ErrInvalidAccessToken
}
