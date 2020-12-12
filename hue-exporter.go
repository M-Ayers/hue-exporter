package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type HueHubResponse struct {
	SensorList []SensorWrapper `json:"sensorList"`
}

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
	var v SensorWrapper
	json.Unmarshal(body, &v)
	// data := v.(map[string]interface{})
	// for k, v := range data {
	// 	switch v := v.(type) {
	// 	case string:
	// 		fmt.Println(k, v, "(string)")
	// 	default:
	// 		fmt.Println(k, v, "(unknown)")
	// 	}
	// }
	fmt.Println(v)
}

func main() {
	queryHueHub(os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY"))
}
