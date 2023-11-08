package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	authErrors "go-learning-demo/auth/error"
	"go-learning-demo/auth/service"
	"net/http"
)

type AuthController struct {
	authService service.IAuthService
}

func NewAuthController(authService service.IAuthService) *AuthController {
	return &AuthController{
		authService: authService,
	}
}

type userRequest struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func (controller *AuthController) SignUp(c *gin.Context) {
	user := new(userRequest)

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if err := controller.authService.SignUp(c.Request.Context(), user.UserName, user.Password); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.Status(http.StatusOK)
}

func (controller *AuthController) SignIn(c *gin.Context) {
	user := new(userRequest)

	if err := c.BindJSON(user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	token, err := controller.authService.SignIn(c.Request.Context(), user.UserName, user.Password)
	if err != nil {
		if errors.Is(err, authErrors.ErrUserNotFound) {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: token})
}
