package service

type UserI interface {
	UserInfo(userID int64) (username *string, err error)
}

type userService struct {
}

func (*userService) UserInfo(userID int64) (username *string, err error) {

	return nil, nil
}
