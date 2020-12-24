package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// SensorState struct stores a key/value string/interface{} pair
type SensorState struct {
	Key   string
	Value interface{}
}

// Sensor struct stores the attributes of a sensor along with any SensorState structs nested within
type Sensor struct {
	Name             string
	Type             string
	ModelID          string
	ManufacturerName string
	SwVersion        string
	States           []SensorState
}

func getSensorsURL(b *Bridge) string {
	endpoint, err := b.GetAPIEndpoint("sensors")
	if err != nil {
		log.Fatal(err)
	}
	return endpoint
}

func callBridge(endpoint string) []byte {
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return body
}

func getSensors(b *Bridge) map[string]interface{} {
	endpoint := getSensorsURL(b)
	body := callBridge(endpoint)
	var sensors map[string]interface{}
	json.Unmarshal(body, &sensors)
	return sensors
}

// GetSensors is the entryway into gathering sensor data from a hue bridge and returning sensor data
func GetSensors(b *Bridge) []Sensor {
	sensors := getSensors(b)
	var response []Sensor
	for key := range sensors {
		sensor := sensors[key].(map[string]interface{})
		sensorState := sensor["state"].(map[string]interface{})
		s := Sensor{
			Name: sensor["name"].(string),
			Type: sensor["type"].(string),
		}

		for key, val := range sensorState {
			s.States = append(
				s.States,
				SensorState{
					Key:   key,
					Value: val,
				})
		}

		response = append(response, s)
	}
	return response
}
