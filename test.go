package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wait sync.WaitGroup
var count = 0

var lock sync.Mutex

func main() {
	wait.Add(10)
	for i := 0; i < 10; i++ {
		go func(data *int) {
			lock.Lock()
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			temp := *data
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
			ans := 1
			*data = temp + ans
			fmt.Println(*data)
			lock.Unlock()
			wait.Done()
		}(&count)
	}
	wait.Wait()
	fmt.Println("最终结果", count)
}
