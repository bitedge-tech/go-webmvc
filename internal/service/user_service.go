package service

import (
	"errors"
	"go-webmvc/internal/repository/model"
	"go-webmvc/internal/repository/query"

	"gorm.io/gorm"
)

type UserI interface {
	UserInfo(userID int64) (user *model.User, err error)
}

type userService struct {
}

func (*userService) UserInfo(userID int64) (user *model.User, err error) {

	//从数据库查询用户信息
	q := query.User
	user, err = q.WithContext(nil).Where(q.ID.Eq(userID)).First()

	// 如果查询出错且不是记录未找到错误，则返回错误
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return user, nil

}
