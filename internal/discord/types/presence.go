package types

import (
	csgoTypes "discord_dota2_cs2/internal/api/csgo_types"
	dotaTypes "discord_dota2_cs2/internal/api/dota_types"
	config "discord_dota2_cs2/internal/configs"
	"discord_dota2_cs2/internal/discord"
	"fmt"
)

type DotaPresence struct {
	State            string // KDA мб голда
	Details          string // Персонаж
	HeroReadableName string // Имя персонажа
	MainImage        string
	SmallImage       string
	LargeText        string
}

func (d *DotaPresence) SetDotaPresenceInfo(response *dotaTypes.GameDotaResponse) {
	d.MainImage = "main"
	switch {
	case *response == (dotaTypes.GameDotaResponse{}):
		d.State = "В меню"
	case response.State.GameState == "DOTA_GAMERULES_STATE_GAME_IN_PROGRESS":
		d.HeroReadableName = discord.DotaHeroes[response.DotaHero.Name]["name"]
		d.MainImage = fmt.Sprintf("https://courier.spectral.gg/images/dota/portraits/%s", discord.DotaHeroes[response.DotaHero.Name]["img"])
		d.State = fmt.Sprintf("KDA: %d/%d/%d,  lvl: %d",
			response.DotaPlayer.Kills, response.DotaPlayer.Deaths, response.DotaPlayer.Assists, response.DotaHero.Level)
		d.LargeText = fmt.Sprintf("gold: %d", response.DotaPlayer.Gold)
		d.Details = fmt.Sprintf("Персонаж: %s - %d%%HP", d.HeroReadableName, response.DotaHero.HealthPercent)
		d.SmallImage = "main"
	default:
		d.State = discord.DotaGameState[response.State.GameState]
	}
}

type CsGoPresence struct {
	State      string // Team - CT, KDA - 5/0/1, money - 4000
	Details    string // Map - Mirage, round - 2, KT/T score ratio - 2/1
	LargeText  string // hp/armor - 100/90
	SmallText  string
	SmallImage string
}

var teamDict = map[string]string{
	"T":  "t_team",
	"CT": "ct_team",
}

func (c *CsGoPresence) SetCsgoPresenceInfo(response *csgoTypes.GameCsgoResponse, settings *config.SteamSettings) {
	gameMode := response.CsGoPlayer.Activity
	switch {
	case gameMode == "menu":
		c.State = "В меню"
	case gameMode == "playing":
		if response.CsGoPlayer.SteamID != settings.SteamID {
			c.State = fmt.Sprintf("Наблюдает за: %s", response.CsGoPlayer.Name)
			return
		}
		c.State = fmt.Sprintf("Команда: %s, KDA: %d/%d/%d, mvps: %d",
			response.CsGoPlayer.Team,
			response.CsGoPlayer.Stats.Kills, response.CsGoPlayer.Stats.Deaths, response.CsGoPlayer.Stats.Assists,
			response.CsGoPlayer.Stats.Mvps)
		c.LargeText = fmt.Sprintf("HP/Броня: %d/%d", response.CsGoPlayer.State.Health,
			response.CsGoPlayer.State.Armor)
		c.SmallImage = teamDict[response.CsGoPlayer.Team]
		c.SmallText = fmt.Sprintf("Команда: %s", response.CsGoPlayer.Team)
		c.Details = fmt.Sprintf("Карта: %s, раунд: %d, раунды выиграные CT/T: %d/%d", response.GameMap.Name,
			response.GameMap.Round, response.GameMap.TeamCt.Score, response.GameMap.TeamT.Score)
	}
}
