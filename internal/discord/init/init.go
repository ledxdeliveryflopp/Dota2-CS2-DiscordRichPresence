package discord

import (
	"discord_dota2_cs2/internal/discord"
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"github.com/shirou/gopsutil/v4/process"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitDiscordClient() {
	for {
		pid, code := FindGamesPid()
		fmt.Println("Game process found")
		log.Infof("game pid: %d, discord code: %s", pid, code)
		err := connectDiscordRPC(code)
		if err != nil {
			log.Panicln("error while connecting to discord RPC:", err)
		}
		discord.GamePid = pid
		log.Info("start check game status loop")
		return
	}
}

func FindGamesPid() (int32, string) {
	for {
		processes, _ := process.Processes()
		for _, game := range processes {
			name, _ := game.Name()
			if name == "cs2.exe" {
				pid := game.Pid
				log.Info("game cs2")
				return pid, "1343901867016585216"
			} else if name == "dota2.exe" {
				pid := game.Pid
				log.Info("game Dota 2")
				return pid, "1342725398261403668"
			}
		}
		fmt.Println("cs2 or dota2 process dont found, repeat after 15 seconds")
		time.Sleep(15 * time.Second)
	}
}

func connectDiscordRPC(discordCode string) error {
	err := client.Login(discordCode)
	if err != nil {
		return err
	}
	fmt.Println("discord inited")
	return nil
}

func CheckGameIsRunning() {
	for {
		time.Sleep(15 * time.Second)
		gameStatus := CheckProcessExist()
		if gameStatus == false {
			log.Infof("proccess with %d pid dont exists", discord.GamePid)
			fmt.Println("Game process dont found")
			fmt.Println("disconnecting from discord rpc")
			err := client.Logout()
			if err != nil {
				log.Error("error while disconnecting from discord rpc:", err)
			}
			log.Info("disconnecting from discord rpc success")
			log.Info("start process find loop")
			time.Sleep(5 * time.Second)
			pid, code := FindGamesPid()
			err = connectDiscordRPC(code)
			if err != nil {
				log.Panicln("error while connecting to discord RPC:", err)
			}
			log.Infof("new process pid: %d", pid)
			discord.GamePid = pid
			fmt.Println("Reconnect to discord rpc success")
		}
		log.Infof("proccess %d is running", discord.GamePid)
	}
}

func CheckProcessExist() bool {
	exist, _ := process.PidExists(discord.GamePid)
	if exist == true {
		return true
	} else {
		return false
	}
}
