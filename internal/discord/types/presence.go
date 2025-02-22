package types

import (
	"discord_dota2/internal/api/types"
	"discord_dota2/internal/discord"
	"fmt"
	"log"
)

type Presence struct {
	State            string // KDA мб голда
	Details          string // Персонаж
	HeroCode         string // Код персонажа
	HeroReadableName string // Имя персонажа
	SmallImage       string
}

func (p *Presence) SetPresenceInfo(player *types.Player, hero *types.Hero) {
	if hero.Level == 0 {
		p.State = "В меню"
		return
	}
	p.HeroCode = hero.Name
	p.HeroReadableName = discord.Heroes[hero.Name]
	p.State = fmt.Sprintf("KDA: %d/%d/%d,  lvl: %d, gold: %d",
		player.Kills, player.Deaths, player.Assists, hero.Level, player.Gold)
	if len(p.HeroReadableName) == 0 {
		log.Printf("Can't find hero in map: %s", hero.Name)
		p.Details = fmt.Sprintf("Персонаж: Неизвестно - %d%%HP", hero.HealthPercent)
	} else {
		p.Details = fmt.Sprintf("Персонаж: %s - %d%%HP", p.HeroReadableName, hero.HealthPercent)
	}
	p.SmallImage = "main"
}
