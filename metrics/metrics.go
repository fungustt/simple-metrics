package metrics

import (
	"sync"
	"sync/atomic"
)

type Metrics struct {
	store sync.Map
}

type Metric struct {
	value uint64
}

func NewMetrics() *Metrics {
	return &Metrics{
		sync.Map{},
	}
}

func (m *Metrics) Get(key string) uint64 {
	metric, _ := m.store.LoadOrStore(key, &Metric{value: 0})

	return metric.(*Metric).value
}

func (m *Metrics) Inc(key string) {
	metric, _ := m.store.LoadOrStore(key, &Metric{value: 0})
	metric.(*Metric).inc()
}

func (m *Metric) inc() {
	atomic.AddUint64(&m.value, 1)
}