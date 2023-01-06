package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func StatusHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))

	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	hostString := optionMap["host"].StringValue()

	if optionMap["port"] != nil {
		hostString += ":" + strconv.Itoa(int(optionMap["port"].IntValue()))
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: ":hourglass_flowing_sand: Please wait while the server status of `%s` is retrieved...",
		},
	})
}
