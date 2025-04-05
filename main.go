package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	helloCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "hello_requests_total",
		Help: "Total number of hello requests.",
	})
)

func init() {
	prometheus.MustRegister(helloCounter)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	helloCounter.Inc()
	log.Println("Вызван эндпоинт /api/v1/hello")
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/api/v1/hello", helloHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Сервер стартует на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
