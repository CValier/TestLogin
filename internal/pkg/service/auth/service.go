package auth

import (
	"time"

	"github.com/CValier/PruebaGO/internal/pkg/entity"
	"github.com/CValier/PruebaGO/internal/pkg/ports"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type authSvc struct {
	repo ports.AuthRepository
}

func NewAuthService(authRepo ports.AuthRepository) *authSvc {
	return &authSvc{
		repo: authRepo,
	}
}

func (a *authSvc) RegisterUser(user *entity.User) error {
	// Hashing the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	errRegister := a.repo.RegisterUser(user)
	if errRegister != nil {
		return errRegister
	}

	return nil
}

func (a *authSvc) LoginUser(email, password string) (string, error) {
	// Get user with email and password
	user, err := a.repo.LoginUser(email, password)
	if err != nil {
		return "", err
	}

	// Generar token JWT
	token, err := generateTokenJWT(user)
	if err != nil {
		return "", err
	}

	return token, nil
}

func generateTokenJWT(user *entity.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	// Set claims to JWT
	claims["email"] = user.Email
	claims["password"] = user.Password
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token válido por 24 horas

	// Generate signed token
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
