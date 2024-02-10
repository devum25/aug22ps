package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	for i := 0; i < 2; i++ {
		emit(i)
	}
	wg.Wait()
}

func emit(id int) {
	reqId := uuid.New()
	fmt.Println("before go routine:", reqId, id)

	go func(id int) {
		fmt.Println("inside go routine:", reqId, id)

		time.Sleep(time.Second * 5)

		fmt.Print("after http call:", reqId, id)
	}(id)

}
