package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {

	var ops uint64

	var wg sync.WaitGroup

	fmt.Printf("ops addr: %p\n", &ops)

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i < 1000; i++ {
				// 使用ops++会发生数据竞争的错误
				//ops++
				atomic.AddUint64(&ops, 1)
			}
		}()
	}

	wg.Wait()
	fmt.Println("ops:", ops)
}
