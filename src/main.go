package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

const (
	BASE_API = "https://api.mc"
)

var (
	config *Config            = &Config{}
	client *discordgo.Session = nil
)

func init() {
	var err error

	if err = config.ReadFile("config.yml"); err != nil {
		log.Fatal(err)
	}

	if client, err = discordgo.New("Bot " + config.Token); err != nil {
		log.Fatal(err)
	}
}

func main() {
	client.AddHandler(ReadyHandler)
	client.AddHandler(InteractionCreateHandler)

	if err := client.Open(); err != nil {
		log.Fatal(err)
	}

	defer client.Close()

	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt)
	<-s
}
