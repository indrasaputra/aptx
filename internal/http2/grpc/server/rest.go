package server

import (
	"fmt"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// RegisterEndpointFunc defines function contract to register endpoint.
type RegisterEndpointFunc func(server *runtime.ServeMux) error

// Rest is responsible to act as HTTP/1.1 REST server.
// It composes grpc-gateway runtime.ServeMux.
type Rest struct {
	*runtime.ServeMux
	port string
}

// NewRest creates an instance of Rest.
func NewRest(port string) *Rest {
	return &Rest{
		ServeMux: runtime.NewServeMux(),
		port:     port,
	}
}

// EnablePrometheus enables prometheus endpoint.
// It can be accessed via /metrics.
func (r *Rest) EnablePrometheus() error {
	return r.ServeMux.HandlePath(http.MethodGet, "/metrics", prometheusHandler())
}

// RegisterEndpoints registers HTTP/1.1 REST endpoints.
// If there are some errors, it returns the first error it encounter and stop the iteration.
func (r *Rest) RegisterEndpoints(fns ...RegisterEndpointFunc) error {
	for _, fn := range fns {
		if err := fn(r.ServeMux); err != nil {
			return err
		}
	}
	return nil
}

// Run runs HTTP/1.1 runtime.ServeMux.
// It is blocking.
func (r *Rest) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", r.port), r.ServeMux)
}

func prometheusHandler() runtime.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, _ map[string]string) {
		promhttp.Handler().ServeHTTP(w, r)
	}
}
