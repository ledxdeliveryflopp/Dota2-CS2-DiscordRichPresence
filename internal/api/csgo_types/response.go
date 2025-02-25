package csgo_types

import (
	"encoding/json"
	"io"
)

type GameCsgoResponse struct {
	CsGoPlayer `json:"player"`
	GameMap    `json:"map"`
}

func (r *GameCsgoResponse) DecodeGameInfo(data io.Reader) error {
	err := json.NewDecoder(data).Decode(&r)
	if err != nil {
		return err
	}
	return nil
}
