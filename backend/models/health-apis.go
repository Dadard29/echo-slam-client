package models

type HealthApisResponse struct {
	Status  bool   `json:"Status"`
	Message string `json:"Message"`
	Content HealthApis `json:"Content"`
}

type HealthApis struct {
	List []ApiCheck `json:"list"`
}

type ApiCheck struct {
	APIName string `json:"api_name"`
	Check   bool   `json:"check"`
}
