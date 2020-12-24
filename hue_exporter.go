package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var port string
var errorEmptyEnvVar = errors.New("getenv: environment variable empty")

func getEnv(key, defVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return defVal
}

func main() {
	log.Println("Starting...")
	port = getEnv("HUE_ENV_PORT", "8080")
	listenPort := flag.String("listen-address", ":"+port, "The address to listen on for HTTP requests.")
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Prometheus metrics available at /metrics")
	log.Println("Looking for bridge")
	b := NewBridge(os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY"))

	sensors := GetSensors(b)
	publishSensorsState(sensors)
	log.Fatal(http.ListenAndServe(*listenPort, nil))
}
