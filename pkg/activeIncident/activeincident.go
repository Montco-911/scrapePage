package activeIncident

import "time"

type ActiveIncident struct {
	Calls []struct {
		Call struct {
			XCord           string `json:"XCord"`
			YCord           string `json:"YCord"`
			Dispatched      string `json:"dispatched"`
			Eid             string `json:"eid"`
			Incidentno      string `json:"incidentno"`
			Incidentsubtype string `json:"incidentsubtype"`
			Incidenttype    string `json:"incidenttype"`
			Location        string `json:"location"`
			Mun             string `json:"mun"`
			Station         string `json:"station"`
			Type            string `json:"type"`
		} `json:"Call"`
		Status []string `json:"Status"`
	} `json:"Calls"`
	TimeStamp time.Time `json:"TimeStamp"`
}
