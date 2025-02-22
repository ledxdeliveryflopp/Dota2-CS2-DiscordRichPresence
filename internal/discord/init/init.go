package discord

import (
	"github.com/hugolgst/rich-go/client"
)

func InitDiscordClient() error {
	err := client.Login("1342725398261403668")
	if err != nil {
		return err
	}
	return nil
}
