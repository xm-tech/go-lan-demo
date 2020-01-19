package main

import (
	"fmt"
	"time"

	redis "github.com/gomodule/redigo/redis"
)

func main() {
	conn, err := getRedisConn()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Do("Set", "maxm", 35)
	if err != nil {
		fmt.Println(fmt.Errorf("set cmd err, %v", err))
		panic("set cmd err")
	}

	r, err := redis.Int(conn.Do("get", "maxm"))
	if err != nil {
		panic("get failed")
	}
	fmt.Printf("r = %+v\n", r)
}

func getRedisConn() (redis.Conn, error) {
	conn, err := redis.Dial("tcp",
		"127.0.0.1:6379",
		redis.DialConnectTimeout(5*time.Second),
		redis.DialReadTimeout(2*time.Second),
		redis.DialWriteTimeout(2*time.Second),
		redis.DialPassword("Maxm#@!123"),
		redis.DialKeepAlive(2*time.Second),
	)
	return conn, err
}
