package metric

import (
	"context"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel/trace"
)

var (
	httpReqDuration *prometheus.HistogramVec
	httpReqCounter  *prometheus.CounterVec
	httpInflight    *prometheus.GaugeVec
)

func init() {
	subsystem := "http_client"
	commonLabelNames := []string{"to_service"}

	httpReqCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name:      "requests_total",
		Namespace: namespace,
		Subsystem: subsystem,
	}, append(commonLabelNames, "code"))
	prometheus.MustRegister(httpReqCounter)

	httpReqDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:      "request_duration_seconds",
		Namespace: namespace,
		Subsystem: subsystem,
		Buckets:   []float64{.01, .1, .5, 1, 5, 10, 30, 60},
	}, commonLabelNames)
	prometheus.MustRegister(httpReqDuration)

	httpInflight = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Name:      "in_flight_requests",
		Namespace: namespace,
		Subsystem: subsystem,
	}, commonLabelNames)
	prometheus.MustRegister(httpInflight)

}

// InstrumentRoundTripper instruments the provided RoundTripper for prometheus metrics.
func InstrumentRoundTripper(toServiceName string, rt http.RoundTripper) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}
	commonLabels := prometheus.Labels{
		"to_service": toServiceName,
	}
	commonLabelValues := []string{toServiceName}
	opts := []promhttp.Option{
		promhttp.WithExemplarFromContext(getContextExemplar),
	}

	rt = promhttp.InstrumentRoundTripperDuration(httpReqDuration.MustCurryWith(commonLabels), rt, opts...)
	rt = promhttp.InstrumentRoundTripperCounter(httpReqCounter.MustCurryWith(commonLabels), rt, opts...)
	rt = promhttp.InstrumentRoundTripperInFlight(httpInflight.WithLabelValues(commonLabelValues...), rt)
	return rt
}

func getContextExemplar(ctx context.Context) prometheus.Labels {
	sctx := trace.SpanFromContext(ctx).SpanContext()
	if sctx.IsSampled() {
		return prometheus.Labels{"trace_id": sctx.TraceID().String()}
	}
	return prometheus.Labels{}
}
