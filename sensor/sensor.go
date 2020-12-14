package sensor

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/M-Ayers/hue-exporter/bridge"

	"github.com/golang/gddo/httputil/header"
)

type Sensor struct {
	Name string
}

func PopulateSensors(b bridge.Bridge) Sensor {
	resp, err := http.Get(fmt.Sprintf("http://%s/api/%s/sensors", b.IpAddr, b.ApiKey))
	if resp.Header.Get("Content-Type") != "" {
		value, _ := header.ParseValueAndParams(resp.Header, "Content-Type")
		if value != "application/json" {
			fmt.Printf("Error with Content-Type.  Got %v\n", value)
		}
	}
	dec := json.NewDecoder(resp.Body)
	fmt.Println(dec)
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	for dec.More() {
		var s Sensor
		err := dec.Decode(&s)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", s.Name)
	}
	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
	return Sensor{"NameValue"}
}
