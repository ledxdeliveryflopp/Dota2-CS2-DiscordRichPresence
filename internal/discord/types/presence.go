package types

import (
	csgoTypes "discord_dota2_cs2/internal/api/csgo_types"
	dotaTypes "discord_dota2_cs2/internal/api/dota_types"
	"discord_dota2_cs2/internal/configs"
	"discord_dota2_cs2/internal/discord"
	"fmt"
)

type DotaPresence struct {
	State            string // KDA мб голда
	Details          string // Персонаж
	HeroCode         string // Код персонажа
	HeroReadableName string // Имя персонажа
	SmallImage       string
}

func (d *DotaPresence) SetDotaPresenceInfo(response *dotaTypes.GameDotaResponse) {
	switch {
	case *response == (dotaTypes.GameDotaResponse{}):
		d.State = "В меню"
	case response.State.GameState == "DOTA_GAMERULES_STATE_GAME_IN_PROGRESS":
		d.HeroCode = response.DotaHero.Name
		d.HeroReadableName = discord.DotaHeroes[response.DotaHero.Name]
		d.State = fmt.Sprintf("KDA: %d/%d/%d,  lvl: %d, gold: %d",
			response.DotaPlayer.Kills, response.DotaPlayer.Deaths, response.DotaPlayer.Assists, response.DotaHero.Level, response.DotaPlayer.Gold)
		d.Details = fmt.Sprintf("Персонаж: %s - %d%%HP", d.HeroReadableName, response.DotaHero.HealthPercent)
		d.SmallImage = "main"
	default:
		d.State = discord.DotaGameState[response.State.GameState]
	}
}

type CsGoPresence struct {
	State   string // Team - CT, heatlh/armor - 100/90,  KDA - 5/0/1, money - 4000
	Details string // Map - Mirage, round - 2, KT/T score ratio - 2/1
}

func (c *CsGoPresence) SetCsgoPresenceInfo(response *csgoTypes.GameCsgoResponse) {
	gameMode := response.CsGoPlayer.Activity
	switch {
	case gameMode == "menu":
		c.State = "В меню"
	case gameMode == "playing":
		if response.CsGoPlayer.SteamID != configs.Settings.SteamID {
			c.State = fmt.Sprintf("Наблюдает за %s", response.CsGoPlayer.Name)
			return
		}
		c.State = fmt.Sprintf("Team - %s | HP/Armor - %d/%d | KDA- %d/%d/%d | mvps - %d",
			response.CsGoPlayer.Team, response.CsGoPlayer.State.Health, response.CsGoPlayer.State.Armor,
			response.CsGoPlayer.Stats.Kills, response.CsGoPlayer.Stats.Deaths, response.CsGoPlayer.Stats.Assists,
			response.CsGoPlayer.Stats.Mvps)
		c.Details = fmt.Sprintf("Map - %s | round - %d | CT/T score - %d/%d", response.GameMap.Name,
			response.GameMap.Round, response.GameMap.TeamCt.Score, response.GameMap.TeamT.Score)
	}
}
