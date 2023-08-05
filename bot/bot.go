package bot

import (
	"fmt"
	"strings"

	"github.com/akhil/discord-ping/config"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	u, err := goBot.User("@me")

	if err != nil {
		fmt.Println(err.Error())
	}

	BotID = u.ID

	goBot.AddHandler(messageHandler)

	err = goBot.Open()

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Bot is running!")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotID {
		return
	}
	args := strings.Split(m.Content, " ")
	if args[0] == config.BotPrefix {
		return
	}
	
	if m.Content == "hello" || m.Content == "Bonjour" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi, "+m.Author.Username+" I'm on 🔥 Fueg ")
	}
	prompt := m.Content[len(prefix):]

	response, err := getOpenAIResponse(prompt)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Une erreur s'est produite lors de la communication avec l'API OpenAI.")
		return
	}

	// Envoyez la réponse du modèle GPT-3.5 turbo sur le serveur Discord
	s.ChannelMessageSend(m.ChannelID,  response)
	
}
func getOpenAIResponse(prompt string) (string, error) {

	// apiURL := 
	// apiKey := 

	return response, nil
}
