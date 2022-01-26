package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Worker struct {
	id int
}

func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}


func main() {
	// c := make(chan int, 100) // channel with buffer size 100
	c := make(chan int) // unbuffered channel
	for i := 0; i < 10; i++ {
		worker := Worker{id: i}

		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
		case <-time.After(time.Millisecond * 100):
		fmt.Println("timed out")
		}
		time.Sleep(time.Millisecond * 50)
	}
}