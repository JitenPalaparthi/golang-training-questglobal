package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	var pingCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "ping_request_count",
			Help: "No of request handled by Ping handler",
		},
	)

	prometheus.MustRegister(pingCounter)

	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		pingCounter.Inc()
		fmt.Fprintln(w, "pong")
	})

	fmt.Println("Application started on port 9091")
	http.ListenAndServe(":9091", nil)
}
