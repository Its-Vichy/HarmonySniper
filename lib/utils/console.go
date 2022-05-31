package utils

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/inancgumus/screen"
	"github.com/its-vichy/harmony/lib/config"
)

func BlockConsoleStd() {
	Sc := make(chan os.Signal, 1)
	signal.Notify(Sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-Sc
}

func PrintLogo() {
	screen.Clear()
	screen.MoveTopLeft()

	fmt.Printf(`
	 _   _                                        
	| | | | __ _ _ __ _ __ ___   ___  _ __  _   _ 
	| |_| |/ _  | '__| '_ ' _ \ / _ \| '_ \| | | |
	|  _  | (_| | |  | | | | | | (_) | | | | |_| |
	|_| |_|\__,_|_|  |_| |_| |_|\___/|_| |_|\__, |
     V%s                                      |___/

`, config.Version)

}

func AppendFile(FileName string, Content string) {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}

	_, err = file.WriteString(Content + "\n")
	if err != nil {
		return
	}
}

func Log(Content string) {
	fmt.Printf("[%s] %s\n", time.Now().Format("02-Jan-2006 15:04:05"), strings.ReplaceAll(Content, "\n", " "))

	if config.SaveLogs {
		AppendFile("Logs.txt", fmt.Sprintf("[%s] %s\n", time.Now().Format("02-Jan-2006 15:04:05"), Content))
	}
}

func Debug(Content string) {
	if config.DebugMode {
		Log(fmt.Sprintf("[DEBUG] %s", strings.ReplaceAll(Content, "\n", " ")))

		if config.SaveLogs {
			AppendFile("Logs.txt", fmt.Sprintf("[DEBUG] [%s] %s\n", time.Now().Format("02-Jan-2006 15:04:05"), strings.ReplaceAll(Content, "\n", " ")))
		}
	}
}
