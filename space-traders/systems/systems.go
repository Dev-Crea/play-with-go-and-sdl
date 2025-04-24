package systems

import (
	"encoding/json"

	space_traders "sdl/playing/space-traders"
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

type System struct {
	SystemSymbol        string
	Symbol              string
	Type                string
	X                   uint32
	Y                   uint32
	Orbitals            []Orbital
	Traits              []Trait
	Modifiers           []string
	Chart               Chart
	Faction             Faction
	IsUnderConstruction bool
}

type Data struct {
	Data System
}

func Init() {
	getSystems()
}

func getSystems() {
	data := space_traders.GetSpaceTradersData("/systems")

	var responseSystem Data

	json.Unmarshal(data, &responseSystem)
}
