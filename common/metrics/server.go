package metrics

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type StandaloneServer struct {
	server *http.Server
}

func NewStandaloneServer(port string) *StandaloneServer {
	mux := http.NewServeMux()

	mux.Handle("/metrics", promhttp.HandlerFor(prometheus.DefaultGatherer, promhttp.HandlerOpts{}))

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: mux,
	}

	return &StandaloneServer{server}
}

func (s *StandaloneServer) Start() {
	err := s.server.ListenAndServe()

	if err != nil {
		log.Fatal("error on start metrics standalone server", err)
	}
}

func (s *StandaloneServer) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
