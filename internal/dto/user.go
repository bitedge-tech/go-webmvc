package dto

import "go-webmvc/internal/repository/model"

type UserListResponse struct {
	BaseResponse
	Count int           `json:"count"`
	Data  []*model.User `json:"data"`
}

type UserInfoRequest struct {
	UserID int64 `form:"user_id" binding:"required"`
}
