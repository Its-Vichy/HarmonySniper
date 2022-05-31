package components

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/its-vichy/harmony/lib/config"
	"github.com/its-vichy/harmony/lib/utils"
)

func HandleMessageEvent(session *discordgo.Session, message *discordgo.MessageCreate) {
	for _, match := range InviteRegex.FindAllString(message.Content, -1) {
		InviteCode := strings.Split(match, "/")[1]

		if !utils.IncludeStr(InviteCode, StrBlacklist) {
			StrBlacklist = append(StrBlacklist, InviteCode)

			utils.AppendFile("invites.txt", InviteCode)

			utils.Log(fmt.Sprintf("Sniped Invite: %s\n", InviteCode))
		}
	}

	// Not optimized, but it works (better with a map, btw websocket server doesnt use maps)
	go func() {
		TotalCheckedMessage++

		for i, acc := range ZombiesAccounts {
			if acc.ID == session.State.User.ID {
				ZombiesAccounts[i].RecievedMessages++
			}
		}

		for i, guild := range ZombiesGuilds {
			if guild.ID == message.GuildID {
				ZombiesGuilds[i].Messages++
			}
		}
	}()
}

func ConnectToWebsocket(Token string) {
	Client, err := discordgo.New(Token)
	log.SetOutput(ioutil.Discard)

	if err == nil {
		if Client.Open() == nil {
			ServerCount := len(Client.State.Guilds)

			if ServerCount > config.MinimumServers {
				Account := DiscordAccount{
					ID:               Client.State.User.ID,
					Username:         Client.State.User.Username,
					Discriminator:    Client.State.User.Discriminator,
					Avatar:           Client.State.User.Avatar,
					Token:            Token,
					GuildSize:        ServerCount,
					RecievedMessages: 0,
					//DiscordSession:   Client,
				}

				ZombiesAccounts = append(ZombiesAccounts, Account)

				for _, guild := range Client.State.Guilds {
					if config.EnableCleaner && guild.MemberCount < config.LeaveIfLowerMemberThan {
						if guild.OwnerID == Client.State.User.ID {
							_, err := Client.GuildDelete(guild.ID)

							if err != nil {
								fmt.Printf("Can't delete server: %s [%s]\n", guild.Name, err)
							} else {
								fmt.Printf("Deleted server: %s, not enought members !\n", guild.Name)
							}

						} else {
							err := Client.GuildLeave(guild.ID)

							if err != nil {
								fmt.Printf("Can't leave server: %s [%s]\n", guild.Name, err)
							} else {
								fmt.Printf("Leaved server: %s, not enought members !\n", guild.Name)
							}
						}

						file, err := os.OpenFile("to_clean.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
						if err != nil {
							log.Fatal(err)
							continue
						}

						_, err = file.WriteString(fmt.Sprintf("%s:%s\n", Token, guild.ID))
						if err != nil {
							log.Fatal(err)
							continue
						}

						file.Close()
					}

					G := DiscordGuild{
						Name:     guild.Name,
						Members:  guild.MemberCount,
						Messages: 0,
						ID:       guild.ID,
						Avatar:   guild.IconURL(),
					}

					ZombiesGuilds = append(ZombiesGuilds, G)
				}

				Client.AddHandler(HandleMessageEvent)
			} else {
				Client.Close()
			}
		}
	}
}