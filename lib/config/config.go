package config

const (
	Version = "0.0.1"
)

var (
	LoadingThreads int = 50
	MinimumServers int = 5
)

// [Under Dev] Websocket server (to make website dashboard, mobile app, etc.)
var (
	WebsocketServerPort   int  = 13254
	EnableWebsocketServer bool = false
)

// Dev & Debug
var (
	DebugMode bool = false
	SaveLogs  bool = true
)

// Token Cleaner
/*
	@EnableCleaner: The cleaner will leave/delete all servers that contains lower than X members.
*/
var (
	LeaveIfLowerMemberThan int  = 100
	EnableCleaner          bool = false
)
