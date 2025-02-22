package handler

import (
	apiTypes "discord_dota2/internal/api/types"
	"discord_dota2/internal/discord"
	"discord_dota2/internal/discord/types"
	"github.com/hugolgst/rich-go/client"
)

func SetDiscordPresence(success chan bool, error chan error, player *apiTypes.Player, hero *apiTypes.Hero) {
	var presence types.Presence
	presence.SetPresenceInfo(player, hero)
	err := client.SetActivity("1342725398261403668", client.Activity{
		State:      presence.State,
		Details:    presence.Details,
		LargeImage: presence.HeroCode,
		LargeText:  presence.HeroReadableName,
		SmallImage: presence.SmallImage,
		SmallText:  "Dota 2",
		//Party: &client.Party{
		//	ID:         "-1",
		//	Players:    15,
		//	MaxPlayers: 24,
		//},
		Timestamps: &client.Timestamps{
			Start: &discord.GameTime,
		},
	})
	if err != nil {
		error <- err
		return
	}
	success <- true
}
