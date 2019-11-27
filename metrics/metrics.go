package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"net/http"
	"strconv"
)

type StatusWrapperForResponseWriter struct {
	http.ResponseWriter
	Status int
}

func (w *StatusWrapperForResponseWriter) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}

const (
	PrometheusNamespace = "api_service"
)

var (
	AccessHits = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: PrometheusNamespace,
		Name:      "hits_by_http_status",
		Help:      "Total hits ordered by http response statuses",
	},
		[]string{"http_status", "path", "method"},
	)
)

func CountHitsMiddleware(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ww := &StatusWrapperForResponseWriter{
			ResponseWriter: w,
			Status:         http.StatusOK,
		}
		next.ServeHTTP(ww, r)

		AccessHits.With(prometheus.Labels{
			"http_status": strconv.Itoa(ww.Status),
			"path":        r.URL.Path,
			"method":      r.Method,
		}).Inc()
	})
}
