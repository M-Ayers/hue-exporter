package hue_exporter

import (
	"fmt"
	"os"
	"time"

	"github.com/M-Ayers/hue-exporter/bridge"
	"github.com/M-Ayers/hue-exporter/sensor"
)

func log(logText string) {
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "-", logText)
}

func queryHueHub(b bridge.Bridge) {
	var sensors = sensor.PopulateSensors(b)
	fmt.Println(sensors)
}

func main() {
	queryHueHub(bridge.GetBridge(os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY")))
}
