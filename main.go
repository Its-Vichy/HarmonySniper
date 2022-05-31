package main

import (
	"github.com/its-vichy/harmony/lib/app"
	"github.com/its-vichy/harmony/lib/components"
	"github.com/its-vichy/harmony/lib/config"
	"github.com/its-vichy/harmony/lib/utils"
	"github.com/zenthangplus/goccm"
)

func main() {
	utils.Log("Harmony is starting...")

	go app.RunWebsocketServer()
	c := goccm.New(config.LoadingThreads)

	for _, Token := range components.ZombiesTokens {
		c.Wait()

		go func(Token string) {
			components.ConnectToWebsocket(Token)
			c.Done()
		}(Token)
	}

	c.WaitAllDone()
	utils.PrintLogo()
	utils.BlockConsoleStd()
}