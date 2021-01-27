// Package service implements video-api servicec
package service

// LoginService provides interface for loggin a new user
type LoginService interface {
	Login(username string, passowrd string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// NewLoginService inits a new loginService
func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "akshit",
		authorizedPassword: "akshit",
	}
}

func (l *loginService) Login(username string, passowrd string) bool {
	return l.authorizedUsername == username && 
			l.authorizedPassword == passowrd
}