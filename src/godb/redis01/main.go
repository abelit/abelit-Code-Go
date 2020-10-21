package main

import (
	"fmt"
	"reflect"

	"github.com/garyburd/redigo/redis"
)

func main() {
	db, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("connect to redis error, ", err)
		return
	}

	fmt.Println(db)

	_, err = db.Do("set", "name", "abelit陈颖")

	if err != nil {
		fmt.Println("save data to redis error, ", err)
		return
	}
	fmt.Println("save data successfully!")

	_, err = db.Do("hset", "user", "name", "chenying", "age", 31, "sex", "male", "phone", "15285649896")

	if err != nil {
		fmt.Println("save hash data to redis error, ", err)
		return
	}
	fmt.Println("save hash data successfully!")

	// 获取set
	sv, err := redis.String(db.Do("get", "name"))

	if err != nil {
		fmt.Println("get err, ", err)
		return
	}
	fmt.Println(reflect.TypeOf(sv))
	fmt.Println(sv)

	// 获取hash
	hv, err := redis.StringMap(db.Do("hgetall", "user"))

	if err != nil {
		fmt.Println("get hash value err, ", err)
		return
	}

	fmt.Println(reflect.TypeOf(hv))
	fmt.Println(hv["name"])

}
