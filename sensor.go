package main

type SensorState struct {
	Key   string
	Value interface{}
}

type Sensor struct {
	Name             string
	Type             string
	ModelId          string
	ManufacturerName string
	SwVersion        string
	States           []SensorState
}

func ParseSensors(sensors map[string]interface{}) []Sensor {
	var response []Sensor
	for key, _ := range sensors {
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
