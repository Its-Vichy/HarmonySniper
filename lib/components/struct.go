package components

//import "github.com/bwmarrin/discordgo"

type DiscordAccount struct {
	ID               string
	Username         string
	Avatar           string
	Discriminator    string
	Token            string
	GuildSize        int
	RecievedMessages int
	//DiscordSession   *discordgo.Session
}

type DiscordGuild struct {
	ID       string
	Name     string
	Members  int
	Messages int
	Avatar   string
}

// Websocket OP Payloads

type OpDstatUpdate struct {
	Op                  string
	TotalCheckedMessage int
}

type OpSniperUpdate struct {
	Op          string
	AccountList []DiscordAccount
}

type OpGuildUpdate struct {
	Op        string
	GuildList []DiscordGuild
}
