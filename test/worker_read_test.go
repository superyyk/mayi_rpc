package test

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mayi/db"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

type city struct {
	Name string
	Lat  string
	Lng  string
}

var rdb *redis.Client

func createCity(c city) {
	// time.Sleep(10 * time.Millisecond)
	rdb = db.Rdb
	ctx := context.Background()
	j := Struct2Map(c)
	jj, _ := json.Marshal(j)
	jjj := string(jj)
	rdb.Set(ctx, c.Lat, jjj, -1)
	fmt.Println("存入redis:", "成功")
}

// 结构体转map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)
	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		data[t.Field(i).Name] = v.Field(i).Interface()
	}
	return data
}

func readData(cityChn chan []city) {
	var cities []city
	csvFile, err := os.Open("test.csv")
	if err != nil {
		fmt.Println("打开csv文件失败")
	}
	defer csvFile.Close()
	csvLines, _ := csv.NewReader(csvFile).ReadAll()
	for _, line := range csvLines {
		if line[0] != "" {

			cities = append(cities, city{
				Name: line[0],
				Lat:  line[1],
				Lng:  line[2],
			})
		}

	}
	cityChn <- cities
}

func worker(cityChn chan city) {
	for val := range cityChn {
		createCity(val)
	}
}

func TestWorkerRead(t *testing.T) {
	startTime := time.Now()
	cities := make(chan []city)
	go readData(cities)
	var workers int = runtime.GOMAXPROCS(runtime.NumCPU()) * 10000
	jobs := make(chan city, 1000)
	for w := 0; w <= workers; w++ {
		go worker(jobs)
	}
	counter := 0
	for _, val := range <-cities {
		counter++
		jobs <- val
	}
	fmt.Println("共处理:", counter)
	fmt.Println("耗时:", time.Since(startTime))

}

func TestRedis(t *testing.T) {
	rdb = db.Rdb
	ctx := context.Background()
	rdb.Set(ctx, "name22", "yyk", -1)
}
