package systems

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
	X                   int32
	Y                   int32
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
	// getSystems()
}

/*
func getSystems() {
	endpoint := url.Values{}
	endpoint.Add("api", "systems")

	data := space_traders.GetSpaceTradersData(endpoint)

	var responseSystem Data

	err := json.Unmarshal(data, &responseSystem)
	if err != nil {
		panic(err)
	}
}
*/
