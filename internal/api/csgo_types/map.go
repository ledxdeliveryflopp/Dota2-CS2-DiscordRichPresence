package csgo_types

type TeamCt struct {
	Score int `json:"score"`
}

type TeamT struct {
	Score int `json:"score"`
}

type GameMap struct {
	Mode   string           `json:"mode"`  // Режим игры
	Name   string           `json:"name"`  // Название карты
	Phase  string           `json:"phase"` // Фаза игры?
	Round  int              `json:"round"` // раунд
	TeamCt `json:"team_ct"` // Счет КТ команды
	TeamT  `json:"team_t"`  // Счет Т команды
}
