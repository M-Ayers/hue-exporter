package main

import (
	"fmt"
	"os"
	"time"
)

func getTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func main() {
	fmt.Printf("%s%s\n", getTime(), " - Starting")
	fmt.Printf("%s%s\n", getTime(), " - Looking for bridge")
	b := New(os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY"))
	endpoint, err := b.GetApiEndpoint()
	if err != nil {
		fmt.Printf("%s%s%s\n", getTime(), " - ", err)
	}
	fmt.Printf("%s%s%s\n", getTime(), " - Found ", endpoint)
}

// func queryHueHub() {
// 	// var sensors = PopulateSensors(Bridge{os.Getenv("HUE_IP_ADDR"), os.Getenv("HUE_API_KEY")})
// 	fmt.Println(sensors)
// }

// func main() {
// 	queryHueHub()
// }
