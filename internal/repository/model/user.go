package model

import (
	"time"
)

type SysUser struct {
	ID        int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(50);not null" json:"username"` //用户名
	Password  string    `gorm:"type:varchar(50); " json:"-"`               //密码
	Nickname  string    `gorm:"type:varchar(50); " json:"nickname"`        //昵称
	Salt      string    `gorm:"type:varchar(100);" json:"-"`               //密码盐值
	Phone     string    `gorm:"type:varchar(20); " json:"phone"`
	Email     string    `gorm:"type:varchar(100); " json:"email"`
	Avatar    string    `gorm:"type:varchar(255); " json:"avatar"`
	Status    int       `gorm:"type:int ;default:0" json:"status"` //0:保存, 1:启用, 9:禁用
	RoleID    int64     `gorm:"type:bigint; " json:"role_id"`      //角色ID
	CreatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"create_at"`
	CreatedBy int64     `gorm:"type:bigint; " json:"create_by"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at;" json:"update_at"`
	UpdatedBy int64     `gorm:"type:bigint; " json:"update_by"`
}

type SysRole struct {
	ID        int64      `gorm:"primaryKey;autoIncrement" json:"id"`
	RoleName  string     `gorm:"type:varchar(50);not null" json:"role_name"`
	RoleDesc  string     `gorm:"type:varchar(255); " json:"role_desc"`
	CreatedAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"create_at"`
	CreatedBy int64      `gorm:"type:bigint; " json:"create_by"`
	UpdatedAt time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at;" json:"update_at"`
	UpdatedBy int64      `gorm:"type:bigint; " json:"update_by"`
	Status    int        `gorm:"type:int;default:0;" json:"status"` //0:保存, 1:启用, 9:禁用
	SysMenu   []*SysMenu `gorm:"many2many:sys_role_menus;" json:"menus"`
}

type SysMenu struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	MenuName   string    `gorm:"type:varchar(50);not null" json:"menu_name"`
	Icon       string    `gorm:"type:varchar(100); " json:"icon"`
	Router     string    `gorm:"type:varchar(255); " json:"router"`
	MenuType   int       `gorm:"type:int; " json:"menu_type"` //
	ReqAction  string    `gorm:"type:varchar(20); " json:"req_action"`
	ParentId   int64     `gorm:"type:bigint; " json:"parent_id"`
	Path       string    `gorm:"type:varchar(255); " json:"path"` // "/0/id/id/id" //层级路径
	Permission string    `gorm:"type:varchar(255); " json:"permission"`
	Sort       int       `gorm:"type:int;default:0 " json:"sort"`
	CreatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP;column:created_at;" json:"create_at"`
	CreatedBy  int64     `gorm:"type:bigint; " json:"create_by"`
	UpdatedAt  time.Time `gorm:"type:timestamp;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;column:updated_at;" json:"update_at"`
	UpdatedBy  int64     `gorm:"type:bigint; " json:"update_by"`
	Status     int       `gorm:"type:int;default:0 " json:"status"` //0:保存, 1:启用, 9:禁用

	SysRole []*SysRole `gorm:"many2many:sys_role_menus;" json:"roles"`
}
