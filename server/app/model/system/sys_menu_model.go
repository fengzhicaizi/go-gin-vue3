package system

import "time"

type SysMenu struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Pid       int       `json:"pid"`
	Name      string    `json:"name"`
	Code      string    `json:"code"`
	Path      string    `json:"path"`
	Sort      int       `json:"sort"`
	Mark      string    `json:"mark"`
	Stuatus   bool      `json:"stuatus"`
	Icon      string    `json:"icon"`
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
