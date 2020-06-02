package metrics

import (
	"sync"
	"testing"
)

func BenchmarkMetricsInc(b *testing.B) {
	b.ReportAllocs()
	m := NewMetrics()
	wg := sync.WaitGroup{}

	for i:= 1; i <= 10; i++ {
		wg.Add(1)
		go func(){
			m.Inc("test")
			wg.Done()
		}()
	}

	wg.Wait()
}
