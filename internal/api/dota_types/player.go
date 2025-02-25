package dota_types

type DotaPlayer struct {
	Name       string `json:"name"`
	Gold       int    `json:"gold"`
	Kills      int    `json:"kills"`
	Deaths     int    `json:"deaths"`
	Assists    int    `json:"assists"`
	KillStreak int    `json:"kill_streak"`
}
