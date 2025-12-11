package service

type LoginI interface {
	Login(username string, password string) (token *string, err error)
}

type loginService struct {
}

func (*loginService) Login(username string, password string) (token *string, err error) {

	return nil, nil
}
