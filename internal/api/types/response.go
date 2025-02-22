package types

import (
	"encoding/json"
	"io"
)

type GameResponse struct {
	Player `json:"player"`
	Hero   `json:"hero"`
}

func (r *GameResponse) DecodeGameInfo(data io.Reader) error {
	err := json.NewDecoder(data).Decode(&r)
	if err != nil {
		return err
	}
	return nil
}
