// Code generated by goctl. DO NOT EDIT.
package types

type AddReq struct {
	Id     int    `form:"id,optional"`
	Name   string `form:"name"`
	Age    int    `form:"age,optional"`
	Uid    string `form:"uid,optional"`
	Puber  string `form:"puber,optional"`
	Time   string `form:"time,optional"`
	Tip    string `form:"tip,optional"`
	Status int    `form:"status,optional,default=0"`
	Level  int    `form:"level,optional"`
	Tel    string `form:"tel,optional"`
}

type Resp struct {
	Res  string      `json:"res"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type UpdateReq struct {
	Id     int    `form:"id,optional"`
	Name   string `form:"name,optional"`
	Age    int    `form:"age,optional"`
	Uid    string `form:"uid,optional"`
	Puber  string `form:"puber,optional"`
	Time   string `form:"time,optional"`
	Tip    string `form:"tip,optional"`
	Status int    `form:"status,optional,default=0"`
	Level  int    `form:"level,optional"`
	Tel    string `form:"tel,optional"`
}

type SearchReq struct {
	Uid string `form:"uid"`
}
