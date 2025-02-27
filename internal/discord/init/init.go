package discord

import (
	config "discord_dota2_cs2/internal/configs"
	"discord_dota2_cs2/internal/discord"
	"fmt"
	"github.com/hugolgst/rich-go/client"
	"github.com/shirou/gopsutil/v4/process"
	log "github.com/sirupsen/logrus"
	"time"
)

func InitDiscordClient() {
	for {
		pid, code, name := FindGamesPid()
		config.MainLog.Println("-------------------------------------------")
		fmt.Println("Game process found")
		config.MainLog.Infof("game pid: %d, game name: %s, discord code: %s", pid, name, code)
		err := connectDiscordRPC(code)
		if err != nil {
			config.MainLog.Println("-------------------------------------------")
			config.MainLog.Panicln("error while connecting to discord RPC:", err)
			return
		}
		discord.GamePid = pid
		return
	}
}

func logProcesses(processList []string) {
	config.MainLog.Println("-------------------------------------------")
	for _, element := range processList {
		if len(element) > 2 {
			config.MainLog.Printf("proccess name: %s", element)
		}
	}
	return
}

func FindGamesPid() (int32, string, string) {
	for {
		processes, _ := process.Processes()
		var processList []string
		for _, game := range processes {
			name, _ := game.Name()
			if name == "cs2.exe" {
				pid := game.Pid
				return pid, "1343901867016585216", name
			} else if name == "dota2.exe" {
				pid := game.Pid
				return pid, "1342725398261403668", name
			} else {
				processList = append(processList, name)
			}
		}
		go logProcesses(processList)
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
			config.MainLog.Infof("proccess with %d pid dont exists", discord.GamePid)
			fmt.Println("Game process dont found")
			fmt.Println("disconnecting from discord rpc")
			err := client.Logout()
			if err != nil {
				config.MainLog.Println("-------------------------------------------")
				log.Error("error while disconnecting from discord rpc:", err)
			}
			config.MainLog.Info("disconnecting from discord rpc success")
			config.MainLog.Info("start process find loop")
			time.Sleep(5 * time.Second)
			pid, code, _ := FindGamesPid()
			err = connectDiscordRPC(code)
			if err != nil {
				config.MainLog.Println("-------------------------------------------")
				log.Panicln("error while connecting to discord RPC:", err)
			}
			config.MainLog.Infof("new process pid: %d", pid)
			discord.GamePid = pid
			fmt.Println("Reconnect to discord rpc success")
		}
		config.MainLog.Infof("proccess %d is running", discord.GamePid)
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
