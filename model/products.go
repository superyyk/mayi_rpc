package model

type Products struct {
	Id        int    `gorm:"column:id",json:"id"`
	Name      string `gorm:"column:name",json:"name"`
	Imgs      string `gorm:"column:imgs",json:"imgs"`
	Guige     string `gorm:"column:guige",json:"guige"`
	Suply     string `gorm:"column:suply",json:"suply"`
	Pertime   string `gorm:"column:pertime",json:"pertime"`
	Perhaocai string `gorm:"column:perhaocai",json:"perhaocai"`
	Uid       string `gorm:"column:uid",json:"uid"`
	Puber     string `gorm:"column:puber",json:"puber"`
	Sort      string `gorm:"column:sort",json:"sort"`
	Time      string `gorm:"column:time",json:"time"`
	Price     string `gorm:"column:price",json:"price"`
	Tip       string `gorm:"column:tip",json:"tip"`
	Unit      string `gorm:"column:unit",json:"unit"`
	Status    string `gorm:"column:status",json:"status"`
	Yuanliao  string `gorm:"column:yuanliao",json:"yuanliao"`
	Liucheng  []Liucheng
}

type Guige struct {
	Id     int    `gorm:"column:id",json:"id"`
	Name   string `gorm:"column:name",json:"name"`
	Status string `gorm:"column:status",json:"status"`
	Puber  string `gorm:"column:puber",json:"puber"`
}
type Yuanliao struct {
	Id     int    `gorm:"column:id",json:"id"`
	Name   string `gorm:"column:name",json:"name"`
	Status string `gorm:"column:status",json:"status"`
	Puber  string `gorm:"column:puber",json:"puber"`
}
