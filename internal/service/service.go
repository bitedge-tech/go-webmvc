package service

type Container struct {
	login LoginI
	user  UserI
}

var Services Container

func InitService() {
	Services = Container{
		login: &loginService{},
		user:  &userService{},
	}
}
