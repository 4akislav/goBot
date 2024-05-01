package internal

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

const prefix string = "!go"

var locationName string

func ManageBot(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	args := strings.Split(m.Content, " ")

	if args[0] != prefix {
		return
	}

	if len(args) == 1 {
		s.ChannelMessageSend(m.ChannelID, "?")
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

		weatherData, err := GetWeather(locationName)
		if err != nil {
			s.ChannelMessageSend(m.ChannelID, "Сталась помилка, можливо ви не привильно ввели назву міста")
			return
		}

		message := fmt.Sprintf("Місто %s, країна %s, %f°C", weatherData.Location.Name, weatherData.Location.Country, weatherData.Current.Temp)
		s.ChannelMessageSend(m.ChannelID, message)

	}

	if args[1] == "weather" && len(args) == 2 {
		s.ChannelMessageSend(m.ChannelID, "Бачу ти хочеш дізнатись погоду в якомусь місті, напиши англійською назву міста. Приклад - !go weather Kyiv")
	}
	if args[1] != CommandsList.Dice && args[1] != CommandsList.ShowWeather && args[1] != CommandsList.TestHello {
		message := "Вибач!\nАле я не розумію що ти від мене хочеш, я знаю тільки певний перелік команд!\n" +
			"Та з радістю їх виконаю!\nТа пам'ятай що я маю своє їм'я, щоб мене покликати пиши просто !go \"команда\""
		s.ChannelMessageSend(m.ChannelID, message)
	}
}
