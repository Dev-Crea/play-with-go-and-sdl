package models

import (
	"encoding/json"
	"net/url"

	space_traders "sdl/playing/internal/space-traders"
)

type Orbital struct {
	Symbol string
}

type Trait struct {
	Symbol      string
	Name        string
	Description string
}

type Faction struct {
	Symbol string
}

type Chart struct {
	SubmittedBy string
	SubmittedOn string
}

type SystemResponse struct {
	SystemSymbol        string
	Symbol              string
	Type                string
	X                   int32
	Y                   int32
	Orbitals            []Orbital
	Traits              []Trait
	Modifiers           []string
	Chart               Chart
	Faction             Faction
	IsUnderConstruction bool
}

type SystemResponseData struct {
	Data SystemResponse
}

func InitSystem() {
	getSystems()
}

func getSystems() {
	endpoint := url.Values{}
	endpoint.Add("api", "systems")

	data := space_traders.GetSpaceTradersData(endpoint)

	var responseSystem SystemResponseData

	err := json.Unmarshal(data, &responseSystem)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Error get systems data")
	}
}
