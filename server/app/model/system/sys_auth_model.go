package system

import "time"

type SysAuth struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	RealName  string    `json:"realName"`
	Phone     string    `json:"phone"`
	RoleId    int       `json:"roleId"`
	IsRoot    bool      `json:"isRoot"`
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
