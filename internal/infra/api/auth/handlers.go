package auth

import (
	"net/http"

	"github.com/CValier/PruebaGO/internal/pkg/entity"
	"github.com/CValier/PruebaGO/internal/pkg/ports"
	"github.com/CValier/PruebaGO/internal/pkg/utils"
	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service ports.AuthService
}

func newHandler(service ports.AuthService) *authHandler {
	return &authHandler{
		service: service,
	}
}

func (a *authHandler) signInUser(c *gin.Context) {

	cred := &entity.Credentials{}

	if err := c.Bind(cred); err != nil {
		c.JSON(http.StatusBadRequest, "Failed to bind credential information: "+err.Error())
		return
	}

	if cred.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "'Email' field is missing."})
	} else if cred.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "'Password' field is missing."})
	}

	token, err := a.service.LoginUser(cred.Email, cred.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func (a *authHandler) registerUser(c *gin.Context) {

	user := &entity.User{}

	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to bind user information: " + err.Error()})
		return
	}

	if isValidate, message := utils.ValidateFields(user); !isValidate {
		c.JSON(http.StatusBadRequest, gin.H{"message": message})
		return
	}

	if !utils.ValidatePassword(user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The password does not meet the minimum requirements."})
		return
	}

	if !utils.ValidateEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The email format is not valid."})
		return
	}

	if !utils.ValidatePhoneNumber(user.PhoneNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "The phone number format is not valid."})
		return
	}

	err := a.service.RegisterUser(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User successfully registered."})
}
