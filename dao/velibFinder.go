package dao

import (
	"VelibFinder/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// GetVelibsStation reads JSON file with Velib stations close to Splio HQ
func GetVelibsStation() ([]model.Station, error) {
	stationFile, _ := ioutil.ReadFile("./stations.json")

	var stations []model.Station
	err := json.Unmarshal(stationFile, &stations); if err != nil {
		log.Printf("Error reading station json file, error : %v", err)
		return nil, err
	}

	return stations, nil
}

// GetAvailableVelibsForStation fetch the velib API and finds the number of available Velibs for given stations
func GetAvailableVelibsForStation(allStations []model.Station) ([]model.Station, error) {

	response, err := http.Get("https://velib-metropole-opendata.smoove.pro/opendata/Velib_Metropole/station_status.json")
	if err != nil {
		log.Printf("Error fetching the Velib API, error : %v", err)
		return nil, err
	}

	var vr model.VelibAPIResponse

	err = json.NewDecoder(response.Body).Decode(&vr)
	if err != nil {
		log.Printf("Error reading Velib API response, error : %v", err)
		return nil, err
	}

	// Get information concerning stations of interest
	for _, stations := range vr.Data.Stations {
		for i, s := range allStations {
			if stations.StationCode == s.Id {
				allStations[i].VelibAvailable = stations.NumBikesAvailable
				allStations[i].MechanicalVelib = stations.NumBikesAvailableTypes[0].Mechanical
				allStations[i].ElectricVelib = stations.NumBikesAvailableTypes[1].Ebike
			}
		}
	}

	return allStations, nil
}

