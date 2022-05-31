package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/its-vichy/harmony/lib/components"
	"github.com/its-vichy/harmony/lib/config"
	"github.com/its-vichy/harmony/lib/utils"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn) {
	// Listen mess
	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				log.Println(err)
				return
			}

			utils.Debug(fmt.Sprintf("Received data on websocket server: %s", string(p)))

			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}

		}
	}()

	// update dsats
	go func() {
		for {
			time.Sleep(1 * time.Second)

			json_OpDstatUpdate, err := json.Marshal(&components.OpDstatUpdate{
				Op:                  "update_dstat",
				TotalCheckedMessage: components.TotalCheckedMessage,
			})

			if err != nil {
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(string(json_OpDstatUpdate))); err != nil {
				log.Println(err)
				return
			}

			// snipers

			json_OpSniperUpdate, err := json.Marshal(&components.OpSniperUpdate{
				Op:          "token_update",
				AccountList: components.ZombiesAccounts,
			})

			if err != nil {
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(string(json_OpSniperUpdate))); err != nil {
				log.Println(err)
				return
			}

			// Guilds

			json_OpGuildUpdate, err := json.Marshal(&components.OpGuildUpdate{
				Op:        "guild_update",
				GuildList: components.ZombiesGuilds,
			})

			if err != nil {
				return
			}

			if err := conn.WriteMessage(websocket.TextMessage, []byte(string(json_OpGuildUpdate))); err != nil {
				log.Println(err)
				return
			}
		}
	}()
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	utils.Debug("New connection to websocket server")
	reader(ws)
}

func SetupRoutes() {
	http.HandleFunc("/ws", WsEndpoint)
}

func RunWebsocketServer() {
	if config.EnableWebsocketServer {
		utils.Log("Websocket server is starting...")

		SetupRoutes()
		log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", config.WebsocketServerPort), nil))
	}
}
