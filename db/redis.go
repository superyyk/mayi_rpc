package db

import (
	"context"
	"fmt"

	red "github.com/gomodule/redigo/redis"
	"github.com/redis/go-redis/v9"
)

// 基础使用说明  https://www.jianshu.com/p/77bc3013ed4d
// redis消息队列使用 https://zhuanlan.zhihu.com/p/416929879
var Rdb *redis.Client
var RedPool *red.Pool

func init() {
	Rdb = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "yyk*2012", // no password set
		DB:       0,          // use default DB
	})

	RedPool = &red.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //数据库最大连接数，0表示最大，无限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (red.Conn, error) { //初始化连接代码
			return red.Dial("tcp", "127.0.0.1:6379")
		},
	}

	//redis链接池用法
	// //取出连接
	// conn := RedPool.Get()
	// //关闭，释放连接
	// defer conn.Close()
	// //添加数据
	// _, err := conn.Do("Set", "name", "老王")
	// if err != nil {
	// 	fmt.Println("数据添加错误:", err)
	// 	return
	// }
	// //读取数据
	// n, err := red.String(conn.Do("Get", "name"))
	// if err != nil {
	// 	fmt.Println("数据读取错误:", err)
	// 	return
	// }
	// fmt.Println(n)
	// //连接池关闭
	// //pool.Close()

	//哨兵模式
	// rdb := redis.NewFailoverClient(&redis.FailoverOptions{
	//     MasterName:    "master",
	//     SentinelAddrs: []string{"x.x.x.x:26379", "xx.xx.xx.xx:26379", "xxx.xxx.xxx.xxx:26379"},
	// })

	//集群
	// rdb := redis.NewClusterClient(&redis.ClusterOptions{
	//     Addrs: []string{":7000", ":7001", ":7002", ":7003", ":7004", ":7005"},
	// })

	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Println("redis 链接失败")
	}
}

// func redisExample() {
//     err := rdb.Set("score", 100, 0).Err()
//     if err != nil {
//         fmt.Printf("set score failed, err:%v\n", err)
//         return
//     }

//     val, err := rdb.Get("score").Result()
//     if err != nil {
//         fmt.Printf("get score failed, err:%v\n", err)
//         return
//     }
//     fmt.Println("score", val)

//     val2, err := rdb.Get("name").Result()
//     if err == redis.Nil {
//         fmt.Println("name does not exist")
//     } else if err != nil {
//         fmt.Printf("get name failed, err:%v\n", err)
//         return
//     } else {
//         fmt.Println("name", val2)
//     }
// }

// func redisExample2() {
//     zsetKey := "language_rank"
//     languages := []redis.Z{
//         redis.Z{Score: 90.0, Member: "Golang"},
//         redis.Z{Score: 98.0, Member: "Java"},
//         redis.Z{Score: 95.0, Member: "Python"},
//         redis.Z{Score: 97.0, Member: "JavaScript"},
//         redis.Z{Score: 99.0, Member: "C/C++"},
//     }
//     // ZADD
//     num, err := rdb.ZAdd(zsetKey, languages...).Result()
//     if err != nil {
//         fmt.Printf("zadd failed, err:%v\n", err)
//         return
//     }
//     fmt.Printf("zadd %d succ.\n", num)

//     // 把Golang的分数加10
//     newScore, err := rdb.ZIncrBy(zsetKey, 10.0, "Golang").Result()
//     if err != nil {
//         fmt.Printf("zincrby failed, err:%v\n", err)
//         return
//     }
//     fmt.Printf("Golang's score is %f now.\n", newScore)

//     // 取分数最高的3个
//     ret, err := rdb.ZRevRangeWithScores(zsetKey, 0, 2).Result()
//     if err != nil {
//         fmt.Printf("zrevrange failed, err:%v\n", err)
//         return
//     }
//     for _, z := range ret {
//         fmt.Println(z.Member, z.Score)
//     }

//     // 取95~100分的
//     op := redis.ZRangeBy{
//         Min: "95",
//         Max: "100",
//     }
//     ret, err = rdb.ZRangeByScoreWithScores(zsetKey, op).Result()
//     if err != nil {
//         fmt.Printf("zrangebyscore failed, err:%v\n", err)
//         return
//     }
//     for _, z := range ret {
//         fmt.Println(z.Member, z.Score)
//     }
// }
