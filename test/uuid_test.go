package test

import (
	"fmt"
	"github.com/superyyk/mayi_rpc/common"
	"github.com/superyyk/mayi_rpc/tool"
	"testing"
)

func TestUuid(t *testing.T) {
	//uid := tool.RandString(12)
	//uid := tool.Get_time_stamp()
	//uid := tool.Md5("yyk")
	//uid := tool.GetOrderUuid(time.Now())
	uid := tool.GooglePassword("123")
	//校验密码
	str := "$2a$10$FIhrxJ3WZmsb2vCwLzHU3uH2fIVMP3BtL07D6rDpIINNQgjCNKr5."
	re := tool.GoogleCheckPassword("123", str)
	fmt.Println(uid)
	fmt.Println("校验结果：", re)
}

func TestTimeOut(t *testing.T) {
	re := common.IsOutTime(1678902800, 1678903009, 5)
	fmt.Println(re)
}
