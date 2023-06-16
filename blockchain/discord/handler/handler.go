package handler

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
)

type MyStruct struct {
}

func messageCreateHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Message.Author.Bot || m.Message.Author.System {
		return
	}
	guildID := m.GuildID
	content := m.Message.Content
	username := m.Message.Author.Username
	guild, err2 := s.Guild(guildID)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(guild.Name)
	fmt.Println(guildID)
	fmt.Println(content)
	fmt.Println(username)
}
