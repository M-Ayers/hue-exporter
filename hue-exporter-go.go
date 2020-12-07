package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func log(logText string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "-", logText)
}

func queryHueHub(hostname string, apiKey string) {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/%s/sensors", hostname, apiKey))
	if err != nil {
		log(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
}

func main() {
	queryHueHub(os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY"))
}
