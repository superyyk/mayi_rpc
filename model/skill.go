package model

type Skill struct {
	Id      int    `gorm:"column:id" json:"id"`
	Name    string `gorm:"column:name" json:"name"`
	Uid     string `gorm:"column:uid" json:"uid"`
	UserId  string `gorm:"column:userid" json:"user_id"`
	Time    int64  `gorm:"column:time" json:"time"`
	Status  int    `gorm:"column:status" json:"status"`
	OrderId string `gorm:"column:orderid" json:"order_id"`
}
