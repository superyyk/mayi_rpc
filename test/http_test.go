package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"mayi/common"
	"mayi/model"
	"testing"
)

func TestHttp(t *testing.T) {

	name := "小工单"
	//num := tool.SmsNum_6()
	tel := "18584911360"
	//uid, _ := tool.GetUuid(12)
	//content := "【" + name + "】您的验证码是：" + num
	//send_url := "http://47.105.221.205:7862/sms"
	//data := "action=send&account=" + account + "&password=" + password + "&mobile=" + tel + "&content=" + content + "&extno=" + extno + "&rt=json"
	//re, err := common.HttpPost(send_url, data)
	send_url := "http://www.kissnet.cn:39005/dianliang/index/send_rpc_sms?tel=" + tel + "&sign=" + name
	re, err := common.HttpGet(send_url)
	if err != nil {
		fmt.Println(errors.New("http_send error"))
	}
	var h *model.Hp
	err = json.Unmarshal([]byte(re), &h)
	fmt.Println(h.Code)

}
