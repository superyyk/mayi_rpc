// Code generated by goctl. DO NOT EDIT.
package types

type StepAddReq struct {
	Name      string `form:"name"`
	Productid string `form:"product_id"`
	Maker     string `form:"maker,optional"`
	Num       int    `form:"num,optional"`
	MakeTime  int64  `form:"maketime,optional"`
	Bad       int    `form:"bad,optional"`
	Time      int64  `form:"time,optional"`
	Status    int    `form:"status,optional"`
}

type StepSearchReq struct {
	Productid string `form:"product_id,optional"`
}

type StepUpdateReq struct {
	Id        int    `form:"id,optional"`
	Name      string `form:"name,optional"`
	Productid string `form:"product_id,optional"`
	Maker     string `form:"maker,optional"`
	Num       int    `form:"num,optional"`
	MakeTime  int64  `form:"maketime,optional"`
	Bad       int    `form:"bad,optional"`
	Time      int64  `form:"time,optional"`
	Status    int    `form:"status,optional"`
}

type Resp struct {
	Res  string      `json:"res"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
