package test

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"
)

// 工厂模型
type Factory struct {
	Wg        *sync.WaitGroup
	MaxWorker int       //最大机器数
	MaxJobs   int       //最大工作数量
	JobQueue  chan int  //工作队列管道
	Quit      chan bool //是否关闭机器
}

// 创建工厂模型
func NewFactory(maxWorker int, wg *sync.WaitGroup) Factory {
	return Factory{
		Wg:        wg,
		MaxWorker: maxWorker,
		JobQueue:  make(chan int, maxWorker),
		Quit:      make(chan bool),
	}
}

// 设置最大订单数量
func (f *Factory) SetMaxJobs(teskNum int) {
	f.MaxJobs = teskNum
}

// 开启机器上班
func (f *Factory) Start() {
	//机器开机，MaxWorker
	for i := 0; i < f.MaxWorker; i++ {
		go func() {
			for {
				select {
				case i := <-f.JobQueue:
					//收到工作开工
					f.doWorker(i)
				case <-f.Quit:
					log.Println("机器关机")
					return
				}
			}
		}()
	}
}

// 分配每个任务到管道中
func (f *Factory) AddTask(taskNum int) {
	f.Wg.Add(1)
	//把任务分配的管道
	f.JobQueue <- taskNum
}

type CsvApi1 struct {
	f *os.File
}

var mycsv CsvApi1

// 工作
func (f *Factory) doWorker(taskNum int) {
	// 生产产品的工作
	//time.Sleep(time.Second)
	j := []string{strconv.Itoa(taskNum), "name" + strconv.Itoa(taskNum), strconv.Itoa(taskNum * taskNum)}

	mycsv.WriteHeader1(j)

	f.Wg.Done()
	log.Println("完工：", taskNum)
}

func (h *CsvApi1) WriteHeader1(heads []string) {
	writer := csv.NewWriter(h.f)
	writer.Write(heads)
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("写入：", err)
	}
	fmt.Println("write:", heads[0])
}
func (h *CsvApi1) CreateFile1(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("create file feailed", err)
		return
	}
	h.f = f
}

func TestFactory(t *testing.T) {
	startTime := time.Now()

	mycsv.CreateFile1("test.csv")
	// 配置工厂核数
	gomaxprocs := runtime.GOMAXPROCS(runtime.NumCPU())
	log.Println("核数：", gomaxprocs)
	//配置监控
	wg := new(sync.WaitGroup)
	//开工厂
	factory := NewFactory(1000*gomaxprocs, wg)
	//订单数量
	factory.SetMaxJobs(1000000000)
	//开始上班
	factory.Start()
	log.Println("开始生产")
	//将所有的订单添加到任务队列
	for i := 0; i < factory.MaxJobs; i++ {
		if i%100 == 0 {
			time.Sleep(time.Microsecond * 10)

		}
		factory.AddTask(i)
	}
	factory.Wg.Wait()

	log.Println("结束耗时：", time.Since(startTime))
}
