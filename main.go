package main

import (
	"fmt"
	"github.com/fungustt/simple-metrics/metrics"
	"sync"
)

func main() {
	m := metrics.NewMetrics()

	wg := sync.WaitGroup{}
	for i:= 1; i <= 10; i++ {
		wg.Add(1)
		go func(){
			m.Inc("test")
			fmt.Println(m.Get("test"))
			wg.Done()
		}()
	}

	wg.Wait()
}
