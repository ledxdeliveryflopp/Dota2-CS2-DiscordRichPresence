package handler

import (
	csgoTypes "discord_dota2_cs2/internal/api/csgo_types"
	dotaTypes "discord_dota2_cs2/internal/api/dota_types"
	config "discord_dota2_cs2/internal/configs"
	"discord_dota2_cs2/internal/discord"
	"discord_dota2_cs2/internal/discord/types"
	"github.com/hugolgst/rich-go/client"
)

func SetDotaPresence(success chan bool, error chan error, response *dotaTypes.GameDotaResponse) {
	var presence types.DotaPresence
	presence.SetDotaPresenceInfo(response)
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

func SetCsGoPresence(success chan bool, error chan error, response *csgoTypes.GameCsgoResponse) {
	var presence types.CsGoPresence
	var settings config.SteamSettings
	settings.InitSettings()
	presence.SetCsgoPresenceInfo(response, &settings)
	err := client.SetActivity("1343901867016585216", client.Activity{
		State:      presence.State,
		Details:    presence.Details,
		LargeImage: "map image",
		LargeText:  "map name",
		SmallImage: "team image",
		SmallText:  "team name",
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
