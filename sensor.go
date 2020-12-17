package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type SensorState struct {
	Key   string
	Value interface{}
}

type Sensor struct {
	Name             string
	Type             string
	ModelID          string
	ManufacturerName string
	SwVersion        string
	States           []SensorState
}

func ParseSensors(b *Bridge) []Sensor {

	endpoint, err := b.GetAPIEndpoint("sensors")
	if err != nil {
		log.Fatal(err)
	}
	resp, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var sensors map[string]interface{}
	json.Unmarshal([]byte(body), &sensors)

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
