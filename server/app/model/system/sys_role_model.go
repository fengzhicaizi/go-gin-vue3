package system

import "time"

type SysRole struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	MenuIds   string    `json:"menuIds"`
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
