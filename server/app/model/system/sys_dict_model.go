package system

import "time"

type SysDictType struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Type      string    `json:"type"`
	Name      string    `json:"name"`
	Remark    string    `json:"remark"`
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}

type SysDictData struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	TypeId    int       `json:"typeId"`
	Label     string    `json:"label"`
	Value     string    `json:"value"`
	Remark    string    `json:"remark"`
	Sort      int       `json:"sort"`
	CreatedAt time.Time `json:"createdAt"` // 创建时间
	UpdatedAt time.Time `json:"updatedAt"` // 更新时间
}
