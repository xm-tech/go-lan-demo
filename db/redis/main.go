package main

import (
	"fmt"
	"time"

	redis "github.com/gomodule/redigo/redis"
)

func main() {
	set("maxm", 35)
	set("tq", 33)
	set("lqh", 30)
	age := get("maxm")
	fmt.Printf("age = %+v\n", age)
}

func set(key string, val interface{}) {
	conn, err := getRedisConnFromPool()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	_, err = conn.Do("Set", key, val)
	if err != nil {
		fmt.Println(fmt.Errorf("set cmd err, %v", err))
		panic("set cmd err")
	}
}

func get(key string) interface{} {
	conn, err := getRedisConnFromPool()
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	r, err := conn.Do("get", key)
	if err != nil {
		panic("get failed")
	}
	fmt.Printf("r = %+v\n", r)
	if val, ok := r.(string); ok {
		return val
	} else if val, ok := r.(int); ok {
		return val
	}
	return nil
}

func getRedisConnFromPool() (redis.Conn, error) {
	pool := &redis.Pool{
		Dial:            func() (redis.Conn, error) { return getRedisConn() },
		MaxIdle:         5,
		MaxActive:       18,
		IdleTimeout:     240 * time.Second,
		MaxConnLifetime: 300 * time.Second,
	}
	return pool.Get(), nil
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
