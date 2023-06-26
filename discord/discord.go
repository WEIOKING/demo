package main

import (
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var botToken = "MTExNjE4ODkwMDkyMzI4OTY4Ng.GmfNmb.g5SjQOaIOKU-9KLnYdu2kJaRJDlEXif87oGyPA"

func main() {
	discord, err := discordgo.New("Bot " + botToken)
	if err != nil {
		log.Fatal(err)
	}
	proxyUrl, _ := url.Parse("http://127.0.0.1:7890")
	discord.Client = &http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxyUrl)}, Timeout: (20 * time.Second)}
	discord.Dialer = &websocket.Dialer{
		Proxy:            http.ProxyURL(proxyUrl),
		HandshakeTimeout: 45 * time.Second,
	}
	discord.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsGuildMembers | discordgo.IntentsGuildInvites | discordgo.IntentsGuilds
	discord.AddHandler(messageCreate)
	discord.AddHandler(ready)
	discord.AddHandler(guildMemberAdd)
	discord.AddHandler(guildMemberRemove)
	discord.AddHandler(inviteCreate)
	discord.AddHandler(inviteDelete)
	discord.AddHandler(guildCreate)
	discord.AddHandler(guildDelete)
	err = discord.Open()
	if err != nil {
		log.Println(err)
	}
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}

var inviteMap = map[string]map[string]int{}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	log.Println("Bot is ready!!!")
	guilds := event.Guilds
	for i := range guilds {
		guild, err := s.Guild(guilds[i].ID)
		if err != nil {
			log.Println("ready:", err)
			return
		}
		invites, err := s.GuildInvites(guild.ID)
		if err != nil {
			log.Println("ready:", err)
			return
		}
		m := inviteMap[guild.ID]
		if m == nil {
			m = map[string]int{}
			inviteMap[guild.ID] = m
		}
		for i2 := range invites {
			invite := invites[i2]
			code := invite.Code
			uses := invite.Uses
			m[code] = uses
		}
	}
	for guildId, m := range inviteMap {
		for code, uses := range m {
			log.Println("guildId: ", guildId, " code:", code, " uses:", uses)
		}
	}
}
func messageCreate(s *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Message.Author.Bot || event.Message.Author.System {
		return
	}
	guildID := event.GuildID
	content := event.Message.Content
	username := event.Message.Author.Username
	guild, err2 := s.Guild(guildID)
	if err2 != nil {
		log.Println("messageCreate:", err2)
		return
	}
	log.Println("guildName: " + guild.Name)
	log.Println("guildID: " + guildID)
	log.Println("Message Content: " + content)
	log.Println("username: " + username)
}

func guildMemberAdd(s *discordgo.Session, event *discordgo.GuildMemberAdd) {
	user := event.User
	guild, err := s.Guild(event.GuildID)
	if err != nil {
		log.Println("guildMemberAdd:", err)
		return
	}
	log.Println("guild name:", guild.Name, " user ", user.Username, " join")
}

func guildMemberRemove(s *discordgo.Session, event *discordgo.GuildMemberRemove) {
	user := event.User
	if user.Bot {
		return
	}
	guild, err := s.Guild(event.GuildID)
	if err != nil {
		log.Println("guildMemberRemove:", err)
		return
	}
	log.Println("guild name:", guild.Name, " user ", user.Username, " remove")
}

func inviteCreate(s *discordgo.Session, event *discordgo.InviteCreate) {
	code := event.Code
	guild, err := s.Guild(event.GuildID)
	if err != nil {
		log.Println("inviteDelete:", err)
		return
	}
	log.Println("guild name:", guild.Name, " invite code ", code, " create")
}

func inviteDelete(s *discordgo.Session, event *discordgo.InviteDelete) {
	code := event.Code
	guild, err := s.Guild(event.GuildID)
	if err != nil {
		log.Println("inviteDelete:", err)
		return
	}
	log.Println("guild name:", guild.Name, " invite code ", code, " delete")
}

func guildCreate(s *discordgo.Session, event *discordgo.GuildCreate) {
	guild := event.Guild
	log.Println("join guild ", guild.Name)
}

func guildDelete(s *discordgo.Session, event *discordgo.GuildDelete) {
	guild := event.BeforeDelete
	log.Println("leave guild ", guild.Name)
}
