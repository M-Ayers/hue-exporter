package main

import (
	"log"
	"os"
	"time"
)

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	log.Println("Starting")
	log.Println("Looking for bridge")
	b := New(os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY"))
	endpoint, err := b.GetApiEndpoint()
	if err != nil {
		log.Fatal("%s%s%s\n", getTime(), " - ", err)
	}
	log.Printf("Endpoint: %s", endpoint)
}
