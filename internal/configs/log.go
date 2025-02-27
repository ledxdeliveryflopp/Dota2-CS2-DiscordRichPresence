package configs

import (
	log "github.com/sirupsen/logrus"
	"os"
)

var CsGoLog = log.New()
var DotaLog = log.New()
var MainLog = log.New()

func initMainLog() {
	mainLog, err := os.OpenFile("main.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	MainLog.Formatter = &log.TextFormatter{}
	if err != nil {
		MainLog.Out = os.Stderr
		MainLog.Println("-------------------------------------------")
		MainLog.Errorf("Error while create main log file: %s", err)
		MainLog.Println("Using os.Stderr")
		return
	}
	MainLog.Out = mainLog
	MainLog.Println("main log inited")
	return
}

func initCsGoLog() {
	csGoLog, err := os.OpenFile("csgo.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	CsGoLog.Formatter = &log.TextFormatter{}
	if err != nil {
		CsGoLog.Out = os.Stderr
		MainLog.Println("-------------------------------------------")
		log.Errorf("Error while create csgo log file: %s", err)
		log.Println("Using os.Stderr")
		return
	}
	CsGoLog.Out = csGoLog
	MainLog.Println("CS:GO log inited")
	return
}

func initDotaLog() {
	dotaLogFile, err := os.OpenFile("dota.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	DotaLog.Formatter = &log.TextFormatter{}
	if err != nil {
		DotaLog.Out = os.Stderr
		MainLog.Println("-------------------------------------------")
		log.Errorf("Error while create dota log file: %s", err)
		log.Println("Using os.Stderr")
		return
	}
	DotaLog.Out = dotaLogFile
	MainLog.Println("Dota log inited")
	return
}

func InitLogrus() {
	initMainLog()
	initCsGoLog()
	initDotaLog()
	return
}
