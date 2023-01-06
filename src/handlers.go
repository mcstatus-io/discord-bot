package main

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func ReadyHandler(s *discordgo.Session, r *discordgo.Ready) {
	log.Printf("Logged in as %s\n", s.State.User)

	s.ApplicationCommandCreate(s.State.User.ID, "", &discordgo.ApplicationCommand{
		Name:        "status",
		Description: "Retrieves the status of a Java Edition Minecraft server.",
		Options: []*discordgo.ApplicationCommandOption{
			{
				Type:        discordgo.ApplicationCommandOptionString,
				Name:        "host",
				Description: "The host of the server",
				Required:    true,
			},
			{
				Type:        discordgo.ApplicationCommandOptionInteger,
				Name:        "port",
				Description: "The port number of the server, defaults to 25565 (range: 0-65536)",
				MinValue:    PointerOf(float64(0)),
				MaxValue:    65536,
				Required:    false,
			},
		},
	})
}

func InteractionCreateHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	log.Println("Interaction")
	StatusHandler(s, i)
}
