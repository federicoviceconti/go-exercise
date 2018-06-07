package main

import (
	"fmt"
	"time"
	"math/rand"
)

type Worker struct {
	id int
}

func main() {
	c := make(chan int)

	for i:=0; i< 4; i++ {
		worker := Worker{id: i}
		go worker.process(c)
	}

	for {
		select {
		case c <- rand.Int():
			fmt.Println(len(c))
			break
		case <-time.After(time.Microsecond * 1):
			fmt.Println("Timed")
		default:
			fmt.Println("dropped!")
		}
		time.Sleep(time.Millisecond* 50)
	}
}

func (w Worker) process(c chan int) {
	for {
		data := <- c
		fmt.Printf("worker %d got %d\n", w.id, data)
	}
}