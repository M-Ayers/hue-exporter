package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
		log.Fatal(err)
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	ParseSensors(result["sensors"].(map[string]interface{}))
}
