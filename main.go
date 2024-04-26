package main

import (
	"encoding/json"
	"fmt"
	"goBot/structs"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const prefix string = "!go"

func main() {
	godotenv.Load()

	var locationName string

	token := os.Getenv("BOT_TOKEN")

	sess, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatal(err)
	}

	sess.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		args := strings.Split(m.Content, " ")

		if args[0] != prefix {
			return
		}

		if len(args) == 1 {
			message := fmt.Sprintln("?")
			s.ChannelMessageSend(m.ChannelID, message)
			return
		}
		if args[1] == "hello" {
			s.ChannelMessageSend(m.ChannelID, "world!")
		}

		if args[1] == "dice" {
			nums := []string{
				"1",
				"2",
				"3",
				"4",
				"5",
				"6",
			}
			selection := rand.Intn(len(nums))

			s.ChannelMessageSend(m.ChannelID, nums[selection])
		}

		if len(args) > 2 && args[1] == "weather" {
			locationName = args[2]

			weatherData, err := getWeather(locationName)
			if err != nil {
				s.ChannelMessageSend(m.ChannelID, "Сталась помилка, можливо ви не привильно ввели назву міста")
				return
			}

			message := fmt.Sprintf("Місто %s, країна %s, %f°C", weatherData.Location.Name, weatherData.Location.Country, weatherData.Current.Temp)
			s.ChannelMessageSend(m.ChannelID, message)

		}

		if args[1] == "weather" && len(args) == 2 {
			message := "Бачу ти хочеш дізнатись погоду в якомусь місті, напиши англійською назву міста. Приклад - !go weather Kyiv"
			s.ChannelMessageSend(m.ChannelID, message)
		}
		if args[1] != structs.CommandsList.Dice && args[1] != structs.CommandsList.ShowWeather && args[1] != structs.CommandsList.TestHello {
			message := "Вибач!\nАле я не розумію що ти від мене хочеш, я знаю тільки певний перелік команд!\nТа з радістю їх виконаю!\nТа пам'ятай що я маю своє їм'я, щоб мене покликати пиши просто !go \"команда\""
			s.ChannelMessageSend(m.ChannelID, message)
		}
	})

	sess.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = sess.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer sess.Close()

	fmt.Println("the bot is online!")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func getWeather(locationName string) (*structs.Weather, error) {
	weather_url := "https://api.weatherapi.com/v1/current.json?key=" + os.Getenv("WEATHER_TOKEN") + "&lang=uk&q=" + locationName + "&aqi=no"

	response, err := http.Get(weather_url)
	if err != nil {
		log.Fatal("error making HTTP request", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error reading response body", err)
	}

	var weatherData structs.Weather
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatal("error parsing JSON", err)
	}
	return &weatherData, err
}
