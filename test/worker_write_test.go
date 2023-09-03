package test

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"
)

type CsvApi struct {
	f *os.File
}

const maxworkers = 100

type Workers struct {
	workerpool chan chan []string
	jobqueue   chan []string
}

func (h *CsvApi) CreateFile(filename string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("create file feailed", err)
		return
	}
	h.f = f
}

func (h *CsvApi) WriteHeader(heads []string) {
	writer := csv.NewWriter(h.f)
	writer.Write(heads)
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("写入：", err)
	}
	fmt.Println("write:", heads[0])
}
func (h *CsvApi) WriteHeaders(heads [][]string) {
	writer := csv.NewWriter(h.f)
	writer.WriteAll(heads)
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("写入：", err)
	}
	fmt.Println("write:", heads[0])
}

//var jobs chan []string

func Do(mycsv CsvApi, datas chan []string) {
	//var mycsv CsvApi
	for val := range datas {
		mycsv.WriteHeader(val)
	}

}

// var job []string


func Consumer(){

}
func TestCsv(t *testing.T) {
	var mycsv CsvApi
	mycsv.CreateFile("test.csv")
	//var data = []string{"id", "name", "age"}
	//var wg *sync.RWMutex
	//mycsv.WriteHeader(data)
	start := time.Now()
	jobs := make(chan []string, 1000)

	for i := 0; i < 1000000000; i++ {
		jj := []string{strconv.Itoa(i), "name" + strconv.Itoa(i), strconv.Itoa(i * i)}
	    jobs <- jj
		fmt.Println(i, "in ok")
	}
    //消费开启
	Consumer()


	//var d [][]string

	fmt.Println(time.Since(start))
}

func (w *Workers) DoWorker(mycsv CsvApi, jobs chan []string) {
	for val := range jobs {
		mycsv.WriteHeader(val)
	}
}

func TestWorker(t *testing.T) {
	fmt.Println("woker")

}
