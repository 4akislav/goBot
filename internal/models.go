package internal

type Weather struct {
	Location `json:"location"`
	Current  `json:"current"`
}

type Location struct {
	Name    string `json:"name"`
	Country string `json:"country"`
}

type Current struct {
	Temp float32 `json:"temp_c"`
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
