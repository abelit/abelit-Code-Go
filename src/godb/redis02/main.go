package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
}

// B2S is a function, transfer bytes to string
func B2S(bs []uint8) string {
	ba := []byte{}
	for _, b := range bs {
		ba = append(ba, byte(b))
	}
	return string(ba)
}

var ErrNil = errors.New("nil returned")

func Byte2String(reply interface{}, err error) (string, error) {
	if err != nil {
		return "", err
	}

	switch reply := reply.(type) {
	case []byte:
		return string(reply), nil
	case string:
		return reply, nil
	case nil:
		return "", ErrNil
	}

	return "", fmt.Errorf("Un expected err, %v", reply)
}

func main() {
	conn := pool.Get()

	defer conn.Close()

	_, err := conn.Do("Set", "name", "好先生")

	if err != nil {
		fmt.Println("set err, ", err)
		return
	}

	v, err := conn.Do("get", "name")

	if err != nil {
		fmt.Println("get err, ", err)
		return
	}

	fmt.Println(reflect.TypeOf(v))

	// fmt.Println(B2S(v.([]uint8)))
	fmt.Println(string(v.([]uint8)))

	fmt.Println(Byte2String(v, err))
}
