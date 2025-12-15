package dto

import "go-webmvc/internal/repository/model"

type UserListResponse struct {
	BaseResponse
	Count int           `json:"count"`
	Data  []*model.User `json:"data"`
}

// UserInfoRequest get 请求参数
type UserInfoRequest struct {
	UserID int64 `form:"user_id" binding:"required"`
}

// UserInfoRequest2 post 请求体
type UserInfoRequest2 struct {
	UserID int64 `json:"user_id" binding:"required"`
}
