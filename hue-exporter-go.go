package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

type State struct {
	ButtonEvent int    `json: "buttonevent"`
	Daylight    bool   `json: "daylight"`
	Presence    bool   `json: "presence"`
	Temperature int    `json: "temperature"`
	LightLevel  int    `json: "lightlevel"`
	Dark        bool   `json: "dark"`
	Flag        bool   `json: "flag"`
	LastUpdated string `json: "lastupdated"`
}

type Config struct {
	On             bool     `json: "on"`
	Battery        int      `json: "battery"`
	Configured     bool     `json: "configured"`
	Reachable      bool     `json: "reachable"`
	Alert          string   `json: "alert"`
	LedIndication  bool     `json: "ledindication"`
	UserTest       bool     `json: "usertest"`
	Pending        []string `json: "pending"` // todo - verify this is array of what
	SunriseOffset  int      `json: "sunriseoffset"`
	SunsetOffset   int      `json: "sunsetoffset"`
	Sensitivity    int      `json: "sensitivity"`
	SensitivityMax int      `json: "sensitivitymax"`
	TholdDark      int      `json: "tholddark"`
	TholdOffset    int      `json: "tholdoffset"`
}

type Events struct {
	ButtonEvent int    `json: "buttonevent"`
	EventType   string `json: "eventtype"`
}

type Inputs struct {
	RepeatIntervals []int    `json: "repeatintervals"`
	Events          []Events `json: "events"`
}

type Capabilities struct {
	Certified bool     `json: "certified"`
	Primary   bool     `json: "primary"`
	Inputs    []Inputs `json: "inputs"`
}

type SWUpdate struct {
	State       string `json: "state"`
	LastInstall string `json: "lastinstall"`
}

type SensorWrapper struct {
	Sensor Sensor
}

type Sensor struct {
	State            State  `json: "state"`
	Config           Config `json: "config"`
	Name             string `json: "name"`
	Type             string `json: "type"`
	ModelId          string `json: "modelid"`
	ManufacturerName string `json: "manufacturername"`
	ProductName      string `json: "productname"`
	DiveristyId      string `json: "diversityid"`
	UniqueId         string `json: "uniqueid"`
	SWVersion        string `json: "swversion"`
	Recycle          bool   `json: "recycle"`
}

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
