package model

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(50);not null" json:"username"` //用户名
	Password  string    `gorm:"type:varchar(50); " json:"-"`               //密码
	Nickname  string    `gorm:"type:varchar(50); " json:"nickname"`        //昵称
	Salt      string    `gorm:"type:varchar(100);" json:"-"`               //密码盐值
	Phone     string    `gorm:"type:varchar(20); " json:"phone"`           //手机号
	Email     string    `gorm:"type:varchar(100); " json:"email"`          //邮箱
	Avatar    string    `gorm:"type:varchar(255); " json:"avatar"`         //头像URL
	Status    int       `gorm:"type:int ;default:0" json:"status"`         //0:保存, 1:启用, 9:禁用
	RoleID    int64     `gorm:"type:bigint; " json:"role_id"`              //角色ID
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"create_at"`
	CreatedBy int64     `gorm:"type:bigint; " json:"create_by"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at;" json:"update_at"`
	UpdatedBy int64     `gorm:"type:bigint; " json:"update_by"`
}

func (User) TableName() string { return "user" }
