package service

type Container struct {
	Login LoginI
	User  UserI
}

var Services Container

func InitService() {
	Services = Container{
		Login: &loginService{},
		User:  &userService{},
	}
}
