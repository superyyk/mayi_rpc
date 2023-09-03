package test

import (
	"context"
	"fmt"
	"mayi/db"
	"mayi/tool"
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	var num int
	val := make(chan map[string]interface{}, 10000)
	redis := db.Rdb
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("退出监控...")
				return
			case data := <-val:
				//fmt.Println(data)
				for i := range data {
					redis.LPush(ctx, i, tool.GetOrderUuid(time.Now()))
				}

			default:
				time.Sleep(time.Second)
				wg := sync.WaitGroup{}
				lc := sync.RWMutex{}
				for i := 0; i <= 10; i++ {
					wg.Add(1)
					d := make(map[string]interface{})
					go func(i int) {
						defer func() {
							fmt.Println("默认监控工作中....:", i)
							//time.Sleep(time.Second * 1)
							lc.Lock()
							num++
							lc.Unlock()
							wg.Done()
						}()

						d["num->"+strconv.Itoa(i)] = num * 2
						val <- d

					}(i)
				}
				wg.Wait()
			}

		}
	}(ctx)
	time.Sleep(time.Second * 10)
	fmt.Println("监控可以停止了....")
	cancel()
	time.Sleep(time.Second * 5)
	fmt.Println("job 完成:", num)

}
