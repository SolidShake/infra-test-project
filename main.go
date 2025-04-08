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
	httpRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Общее количество HTTP запросов, разбитых по коду ответа и методу",
		},
		[]string{"code", "method"},
	)
)

func init() {
	prometheus.MustRegister(helloCounter)
	prometheus.MustRegister(httpRequestsTotal)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	helloCounter.Inc()
	log.Println("Вызван эндпоинт /api/v1/hello")
	fmt.Fprintln(w, "Hello, World!")
}

func errorHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Вызван эндпоинт /api/v1/error")
	w.WriteHeader(http.StatusInternalServerError)
	fmt.Fprintln(w, "internal server error")
}

func main() {
	http.HandleFunc("/api/v1/hello", promhttp.InstrumentHandlerCounter(httpRequestsTotal, http.HandlerFunc(helloHandler)))
	http.HandleFunc("/api/v1/error", promhttp.InstrumentHandlerCounter(httpRequestsTotal, http.HandlerFunc(errorHandler)))
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Сервер стартует на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
