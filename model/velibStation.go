package model

// Station is the struct that gather useful information for VelibFinder employees
type Station struct {
	Name string `json:"station"`
	Id string `json:"station_id"`
	VelibAvailable int `json:"num_bikes_available"`
	MechanicalVelib int `json:"num_mechanical_bikes"`
	ElectricVelib int `json:"num_electric_bikes"`
	StationImage string `json:"station_img"`
}

// VelibAPIResponse is the API response for the stations status
type VelibAPIResponse struct {
	LastUpdatedOther int `json:"lastUpdatedOther"`
	TTL              int `json:"ttl"`
	Data             struct {
		Stations []struct {
			StationCode            string `json:"stationCode"`
			StationID              int    `json:"station_id"`
			NumBikesAvailable      int    `json:"num_bikes_available"`
			NumBikesAvailableTypes []struct {
				Mechanical int `json:"mechanical,omitempty"`
				Ebike      int `json:"ebike,omitempty"`
			} `json:"num_bikes_available_types"`
			NumDocksAvailable int `json:"num_docks_available"`
			IsInstalled       int `json:"is_installed"`
			IsReturning       int `json:"is_returning"`
			IsRenting         int `json:"is_renting"`
			LastReported      int `json:"last_reported"`
		} `json:"stations"`
	} `json:"data"`
}