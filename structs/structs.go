package structs

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Temp float32 `json:"temp_c"`
	} `json:"current"`
}

type Commands struct {
	ShowWeather string
	TestHello   string
	Dice        string
}

var CommandsList = Commands{
	ShowWeather: "weather",
	TestHello:   "hello",
	Dice:        "dice",
}
