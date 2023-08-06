package bot

import (
	"fmt"
	"strings"

	"bot/config"

	"github.com/bwmarrin/discordgo"
)



var BotID string


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

	if m.Content == "hello" || m.Content == "Bonjour" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "Hi, "+ m.Author.Username +" I'm on 🔥 Fueg ")
	}
	args := strings.Split(m.Content, " ")
	if args[0] == config.BotPrefix {
		return
	}
	


	prompt := m.Content[len(config.BotPrefix):]

	apikey:=config.Apikey
	organization:=config.Apiorg
 	client := NewClient(apikey, organization)

 r := CreateCompletionsRequest{
  Model: "gpt-3.5-turbo",
  Messages: []Message{
   {
    Role:    m.Author.Username,
    Content: prompt,
   },
  },
  Temperature: 0.7,
 }

 completions, err := client.CreateCompletions(r)
 if err != nil {
  panic(err)
 }

 fmt.Println(completions)

	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Une erreur s'est produite lors de la communication avec l'API OpenAI.")
		return
	}

	// Envoyez la réponse du modèle GPT-3.5 turbo sur le serveur Discord
	
	if len(completions.Choices) > 0 {
		messageContent := completions.Choices[0].Message.Content
		s.ChannelMessageSend(m.ChannelID, messageContent)
	} else {
		s.ChannelMessageSend(m.ChannelID, "Aucune réponse n'a été reçue du modèle.")
	}
	
}


