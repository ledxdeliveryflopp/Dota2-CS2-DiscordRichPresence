package dota_types

import (
	"encoding/json"
	"io"
)

type GameDotaResponse struct {
	State      `json:"map"`
	DotaPlayer `json:"player"`
	DotaHero   `json:"hero"`
}

func (r *GameDotaResponse) DecodeGameInfo(data io.Reader) error {
	err := json.NewDecoder(data).Decode(&r)
	if err != nil {
		return err
	}
	return nil
}
