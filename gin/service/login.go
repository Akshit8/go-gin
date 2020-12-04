package service

// LoginService ...
type LoginService interface {
	Login(username string, password string) bool
}

type loginService struct {
	authorizedUsername string
	authorizedPassword string
}

// NewLoginService ...
func NewLoginService() LoginService {
	return &loginService{
		authorizedUsername: "Akshit",
		authorizedPassword: "test123",
	}
}

func (loginSrv *loginService) Login(username string, password string) bool {
	return loginSrv.authorizedUsername == username && loginSrv.authorizedPassword == password
}
