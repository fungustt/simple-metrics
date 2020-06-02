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

const ValueNotSet = 0

func NewMetrics() *Metrics {
	return &Metrics{
		sync.Map{},
	}
}

func (m *Metrics) Get(key string) uint64 {
	metric, exist := m.store.Load(key)

	if !exist {
		return ValueNotSet
	}

	return metric.(*Metric).value
}

func (m *Metrics) Inc(key string) {
	metric, exist := m.store.Load(key)

	if !exist {
		metric = &Metric{value: 0}
		m.store.Store(key, metric)
	}

	metric.(*Metric).inc()
}

func (m *Metric) inc() {
	atomic.AddUint64(&m.value, 1)
}