package test

import (
	"context"
	"fmt"
	"github.com/superyyk/mayi_rpc/db"
	"runtime"
	"testing"
	"time"
)

var (
	MaxWorker = runtime.NumCPU()       //os.Getenv("MAX_WORKERS")
	MaxQueue  = runtime.NumCPU() * 100 //os.Getenv("MAX_QUEUE")
)

type Payload struct {
	Num int
}
type Job struct {
	Payload Payload
}

var JobQueue chan Job

type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerpool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerpool,
		JobChannel: make(chan Job),
		quit:       make(chan bool),
	}
}

func (p Payload) Print(num int) {
	fmt.Println("doooo something:", num)
}

func (w Worker) Start() {
	go func() {
		for {
			//注册当前worker到worker队列
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				//收到一个工作请求
				job.Payload.Print(job.Payload.Num)
			case <-w.quit:
				// 收到停止工作信号
				return
			}

		}

	}()
}

// stop同志worker停止监听工作
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

type Dispatcher struct {
	WorkerPool chan chan Job
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < 1000; i++ {
		worker := NewWorker(d.WorkerPool)
		worker.Start()
	}
	go d.dispatch()
}
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-JobQueue:
			// a job request has been received
			go func(job Job) {
				// try to obtain a worker job channel that is available.
				// this will block until a worker is idle
				jobChannel := <-d.WorkerPool

				// dispatch the job to the worker job channel
				jobChannel <- job
			}(job)
		}
	}
}

func TestJobs(t *testing.T) {
	dispather := NewDispatcher(MaxWorker)
	rdb := db.Rdb
	rdb.Set(context.Background(), "ddd", "dddddddd", -1)
	dispather.Run()
	for i := 0; i < 1000; i++ {

		p := Payload{
			Num: i,
		}
		job := Job{
			Payload: p,
		}
		JobQueue <- job
		time.Sleep(time.Second)
	}
}
