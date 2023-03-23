package model

type Weather struct {
	Location struct {
		Name      string `json:"name"`
		Region    string `json:"region"`
		Country   string `json:"country"`
		Localtime string `json:"localtime"`
	} `json:"location"`
	Current struct {
		TempC float32 `json:"temp_c"`
		TempF float32 `json:"temp_f"`
	} `json:"current"`
}
