package model

type ConnInfo struct {
	Username string
	Password string
	Host     string
	Db       string
	Port     int
}

type BiInfo struct {
	Id     string `gorm:"column:fID",json:"id"`
	Pairid string `gorm:"column:pairId",json:"pairid"`
	Logo   string `gorm:"column:logo",json:"logo"`
	Price  string `gorm:"column:price",json:"price"`
}

type Area struct {
	Code string `gorm:"column:code",json:"code"`
	Name string `gorm:"column:name",json:"name"`
}

type Hp struct {
	Res  int    `json:"res"`
	Uid  string `json:"uid"`
	Data string `json:"data"`
	Code int    `json:"code"`
}

type Codes struct {
	Uid  string `gorm:"column:uid",json:"uid"`
	Code int    `gorm:"column:code",json:"code"`
	Tel  string `gorm:"column:tel",json:"tel"`
	Time int64  `gorm:"column:time",json:"time"`
}
