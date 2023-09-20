package test

import (
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/superyyk/mayi_rpc/db"
	"github.com/superyyk/mayi_rpc/tool"

	"github.com/jinzhu/gorm"
)

var query = "go"
var matchs = 0

var maxWorkerCount = 32
var workerCount = 0
var searchRequest = make(chan string)
var fundMatchs = make(chan bool)
var workerDone = make(chan bool)
var DB *gorm.DB

type users struct {
	Id       int64  `gorm:"column=id",json:"id"`
	Pid      string `gorm:"column=pid",json:"pid"`
	Username string `gorm:"column=username",json:"username"`
	Uuid     string `gorm:"column=uuid",json:"uuid"`
	Mobile   string `gorm:"column=mobile",json:"mobile"`
	EMail    string `gorm:"column=email",json:"email"`
	Pids     string `gorm:"column=pids",json:"pids"`
}

func TestFindFile(t *testing.T) {
	start := time.Now()
	workerCount = 1
	num := 977
	DB = db.Db
	//var count = 0
	var users []users
	DB.Table("users").Where("comm=?", 0).Find(&users)
	//fmt.Println(users)

	go search(users, tool.Int_string(num), true)
	waitForWorkers()

	fmt.Println("找到：", matchs)
	fmt.Println(time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest: //有新工作
			workerCount++
			println(path)
			//go search(, true)
		case <-workerDone: //工人找到结束
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-fundMatchs: //新的结果被找到
			matchs++
		}
	}
}

func search(us []users, num string, master bool) {
	//files, err := os.ReadDir(path)

	// if err == nil {
	// 	for _, file := range files {
	// 		name := file.Name()
	// 		if name == query {
	// 			//matchs++
	// 			fundMatchs <- true
	// 		}
	// 		if file.IsDir() {
	// 			if workerCount < maxWorkerCount {
	// 				searchRequest <- path + name + "/"
	// 			} else {
	// 				search(path+name+"/", false)
	// 			}

	// 		}
	// 	}
	// 	if master {
	// 		workerDone <- true
	// 	}
	// }

	if len(us) > 0 {
		var ll []int
		for _, user := range us {

			pids := strings.Split(user.Pids, ",")
			if pids[0] == num { //找到
				fundMatchs <- true
			}
			ll = append(ll, len(pids))
		}
		fmt.Println("推广代数", tool.GetMaxNum(ll))

		if master {
			workerDone <- true
		}
	}

}
