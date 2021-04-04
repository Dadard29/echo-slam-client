package models

type HealthApisResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	Content HealthApis `json:"Content"`
}

type HealthApis struct {
	List []struct {
		APIName string `json:"api_name"`
		Check   bool   `json:"check"`
	} `json:"list"`
}
