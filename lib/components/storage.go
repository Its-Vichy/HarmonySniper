package components

import (
	"regexp"

	"github.com/its-vichy/harmony/lib/utils"
)

var (
	TotalCheckedMessage int = 0
	
	ZombiesAccounts     []DiscordAccount
	ZombiesGuilds       []DiscordGuild

	StrBlacklist        []string = utils.LoadTokensFromFile("invites.txt")
	ZombiesTokens       []string = utils.LoadTokensFromFile("tokens.txt")
)

var (
	InviteRegex = regexp.MustCompile("(discord.gg/)([0-9a-zA-Z]+)")
)
