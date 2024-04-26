package main

import (
	"encoding/json"
	"fmt"
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

const prefix string = "!bot"

type Weather struct {
	Location struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"location"`
	Current struct {
		Temp float32 `json:"temp_c"`
	} `json:"current"`
}

func main() {
	godotenv.Load()

	weather_url := "https://api.weatherapi.com/v1/current.json?key=" + os.Getenv("WEATHER_TOKEN") + "&q=Kyiv&aqi=no&lang=ja"

	response, err := http.Get(weather_url)
	if err != nil {
		log.Fatal("error making HTTP request", err)
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("error reading response body", err)
	}

	var weatherData Weather
	err = json.Unmarshal(body, &weatherData)
	if err != nil {
		log.Fatal("error parsing JSON", err)
	}

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

		if args[1] == "weather" {
			message := fmt.Sprintf("Місто %s", weatherData.Location.Name)
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
