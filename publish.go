package main

import "log"

// Publish iterates through Sensor array and publishes a Prometheus metric based on the sensor type
func publishSensorsState(sensors []Sensor) {
	for _, sensor := range sensors {
		switch sensor.Type {
		case "CLIPGenericFlag":
			log.Printf("CLIPGenericFlag Sensor: %s", sensor.Name)
		case "CLIPGenericStatus":
			log.Printf("CLIPGenericStatus Sensor: %s", sensor.Name)
		case "ZLLLightLevel":
			log.Printf("ZLLLightLevel Sensor: %s", sensor.Name)
		case "ZLLPresence":
			log.Printf("ZLLPresence Sensor: %s", sensor.Name)
		case "ZLLSwitch":
			log.Printf("ZLLSwitch Sensor: %s", sensor.Name)
		case "Geofence":
			log.Printf("Geofence Sensor: %s", sensor.Name)
		case "ZLLTemperature":
			log.Printf("ZLLTemperature Sensor: %s", sensor.Name)
		case "Daylight":
			log.Printf("Daylight Sensor: %s", sensor.Name)
		default:
			log.Printf("Other sensor: %s", sensor.Name)
		}
	}
}

// More types I dont have
// CLIPSwitch
// ZLLZGPSwitch
// ZGPSwitch
// CLIPOpenClose
// CLIPPresence
// CLIPTemperature
// CLIPHumidity
// CLIPLightlevel
// ZLLRelativeRotary
