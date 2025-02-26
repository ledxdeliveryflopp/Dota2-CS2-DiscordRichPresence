package csgo_types

type PlayerState struct {
	Health int `json:"health"`
	Armor  int `json:"armor"`
	Money  int `json:"money"`
}

type MatchStats struct {
	Kills   int `json:"kills"`
	Assists int `json:"assists"`
	Deaths  int `json:"deaths"`
	Mvps    int `json:"mvps"`
}

type CsGoPlayer struct {
	SteamID  string      `json:"steamid"`
	Name     string      `json:"name"`     // Nickname в Steam
	Team     string      `json:"team"`     // Команда за которую играет
	Activity string      `json:"activity"` // Что делает(в меню, в игре и тд)
	State    PlayerState `json:"state"`
	Stats    MatchStats  `json:"match_stats"`
}
