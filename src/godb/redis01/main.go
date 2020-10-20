package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	db, err := redis.Dial("tcp", "127.0.0.1:6379")

	if err != nil {
		fmt.Println("connect to redis error, ", err)
		return
	}

	fmt.Println(db)

	_, err = db.Do("set", "name", "abelit")

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
}
