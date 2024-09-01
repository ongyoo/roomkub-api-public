package metric

import "github.com/prometheus/client_golang/prometheus"

const (
	namespace = "room"
	subsystem = "kub"
)

var (
	cacheHit    *prometheus.CounterVec
	cacheMissed *prometheus.CounterVec
)

func init() {
	cacheHit = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      "cache_hit",
		Namespace: namespace,
		Subsystem: subsystem,
	}, []string{"name"})
	prometheus.MustRegister(cacheHit)

	cacheMissed = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      "cache_miss",
		Namespace: namespace,
		Subsystem: subsystem,
	}, []string{"name"})
	prometheus.MustRegister(cacheMissed)
}

// IncCacheHit increments the cache hit counter
func IncCacheHit(name string) {
	cacheHit.With(prometheus.Labels{
		"name": name,
	}).Inc()
}

// IncCacheMiss increments the cache miss counter
func IncCacheMiss(name string) {
	cacheMissed.With(prometheus.Labels{
		"name": name,
	}).Inc()
}
