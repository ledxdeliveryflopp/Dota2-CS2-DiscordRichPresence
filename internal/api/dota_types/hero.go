package dota_types

type DotaHero struct {
	Name            string `json:"name"`
	Level           int    `json:"level,omitempty"`
	Alive           bool   `json:"alive"`
	HealthPercent   int    `json:"health_percent"`
	AghanimsScepter bool   `json:"aghanims_scepter"`
}
