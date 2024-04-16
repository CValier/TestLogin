package ports

import "github.com/CValier/PruebaGO/internal/pkg/entity"

type AuthRepository interface {
	RegisterUser(usuario *entity.User) error
	LoginUser(email, password string) (*entity.User, error)
}

type AuthService interface {
	RegisterUser(usuario *entity.User) error
	LoginUser(email, password string) (string, error)
}
