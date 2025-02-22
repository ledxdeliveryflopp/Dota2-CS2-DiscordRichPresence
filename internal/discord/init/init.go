package discord

import (
	"github.com/hugolgst/rich-go/client"
	"log"
)

func InitDiscordClient() error {
	err := client.Login("1342725398261403668")
	if err != nil {
		return err
	}
	log.Println("discord inited")
	return nil
}
