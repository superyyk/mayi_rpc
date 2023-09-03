// Code generated by goctl. DO NOT EDIT.
package types

type YuanliaoAddReq struct {
	Puder string `form:"uid"`
	Name  string `form:"name"`
}

type YuanliaoReq struct {
	Puder string `form:"uid"`
	Id    string `form:"id"`
}

type GuigeAddReq struct {
	Puder string `form:"uid"`
	Name  string `form:"name"`
}

type GuigeDelReq struct {
	Puder string `form:"uid"`
	Id    string `form:"id"`
}

type AddReq struct {
	Id        int    `form:"id,optional"`
	Name      string `form:"name"`
	Imgs      string `form:"imgs,optional"`
	Unit      string `form:"unit"` //计量单位
	Uid       string `form:"Uid,optional"`
	Guige     string `form:"guige,optional"`
	Suply     string `form:"suply,optional"`
	PerTime   string `form:"pertime,optional"`
	PerHaocai string `form:"perhaocai,optional"`
	Puber     string `form:"puber,optional"`
	Time      string `form:"time,optional"`
	Price     string `form:"price,optional"`
	Tip       string `form:"tip,optional"`
	Status    int    `form:"status,optional,default=0"`
	Yuanliao  string `form:"yuanliao,optional"`
}

type Resp struct {
	Res  string      `json:"res"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UpdateReq struct {
	Name      string `form:"name,optional"`
	Imgs      string `form:"imgs,optional"`
	Unit      string `form:"unit,optional"` //计量单位
	Uid       string `form:"Uid,optional"`
	Guige     string `form:"guige,optional"`
	Suply     string `form:"suply,optional"`
	PerTime   string `form:"pertime,optional"`
	PerHaocai string `form:"perhaocai,optional"`
	Puber     string `form:"puber,optional"`
	Time      string `form:"time,optional"`
	Price     string `form:"price,optional"`
	Tip       string `form:"tip,optional"`
	Status    int    `form:"status,optional"`
	Yuanliao  string `form:"yuanliao,optional"`
}

type SearchReq struct {
	Uid string `form:"uid"`
}
