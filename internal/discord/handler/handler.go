package handler

import (
	csgoTypes "discord_dota2_cs2/internal/api/csgo_types"
	dotaTypes "discord_dota2_cs2/internal/api/dota_types"
	config "discord_dota2_cs2/internal/configs"
	"discord_dota2_cs2/internal/discord"
	"discord_dota2_cs2/internal/discord/types"
	"github.com/hugolgst/rich-go/client"
)

func SetDotaPresence(success chan bool, error chan types.DotaPresenceError, response *dotaTypes.GameDotaResponse) {
	var presence types.DotaPresence
	presence.SetDotaPresenceInfo(response)
	err := client.SetActivity("1342725398261403668", client.Activity{
		State:      presence.State,
		Details:    presence.Details,
		LargeImage: presence.MainImage,
		LargeText:  presence.LargeText,
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
		var dotaPresenceError types.DotaPresenceError
		dotaPresenceError.SetErrors(&presence, err)
		error <- dotaPresenceError
		return
	}
	success <- true
}

func SetCsGoPresence(success chan bool, error chan types.CsGoPresenceError, response *csgoTypes.GameCsgoResponse) {
	var presence types.CsGoPresence
	var settings config.SteamSettings
	settings.InitSettings()
	presence.SetCsgoPresenceInfo(response, &settings)
	err := client.SetActivity("1343901867016585216", client.Activity{
		State:      presence.State,
		Details:    presence.Details,
		LargeImage: "main",
		LargeText:  presence.LargeText,
		//LargeText:  "CS:GO 2",
		SmallImage: presence.SmallImage,
		SmallText:  presence.SmallText,
		Timestamps: &client.Timestamps{
			Start: &discord.GameTime,
		},
	})
	if err != nil {
		var csGoPresenceError types.CsGoPresenceError
		csGoPresenceError.SetErrors(&presence, err)
		error <- csGoPresenceError
		return
	}
	success <- true
}
