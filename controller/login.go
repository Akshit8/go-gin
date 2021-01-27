// Package controller implements api endpoint controllers
package controller

import (
	"github.com/Akshit8/go-gin/entity"
	"github.com/Akshit8/go-gin/service"
	"github.com/gin-gonic/gin"
)

// LoginController defines method for login controller
type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

// NewLoginController inits a loginController
func NewLoginController(
	loginService service.LoginService,
	jwtService service.JWTService,
) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (lc *loginController) Login(ctx *gin.Context) string {
	var credentials entity.Credentials
	err := ctx.ShouldBind(&credentials)
	if err != nil {
		return err.Error()
	}
	isAuthenticated := lc.loginService.Login(credentials.Username, credentials.Password)
	if isAuthenticated {
		return lc.jwtService.GenerateToken(credentials.Username, true)
	}
	return ""
}
