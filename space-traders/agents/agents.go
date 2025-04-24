package agents

import (
	"encoding/json"
	"fmt"
	"strings"

	system "sdl/playing/scenes/systems"
	space_traders "sdl/playing/space-traders"
	"sdl/playing/space-traders/contracts"
	systems "sdl/playing/space-traders/systems"

	"github.com/veandco/go-sdl2/sdl"
)

type Agent struct {
	Symbol          string
	Headquarters    string
	Credits         int32
	StartingFaction string
	ShipCount       int32
	Location        Location
}

type Data struct {
	Data Agent
}

type Location struct {
	Sector   string
	System   string
	Waypoint string
}

var (
	MyAgent    Data
	MyPosition systems.Data
	MyContract string
)

func Init() {
	MyAgent = getAgentData()
	MyPosition = getAgentStartData()

	getAgentStartContract()
	postAgentNegociateContract()
}

func GetCurrentOrbitals(surface sdl.Surface) {
	// fmt.Printf("== GetCurrentOrbitals == \n")
	for key := range MyPosition.Data.Orbitals {
		// fmt.Printf("[%v]> %s\n", key, value)
		system.System(surface, int32((key+1)*50))
	}
}

func getAgentData() Data {
	data := space_traders.GetSpaceTradersData("/my/agent")

	var responseAgent Data

	json.Unmarshal(data, &responseAgent)

	return responseAgent
}

func getAgentStartData() systems.Data {
	data := space_traders.GetSpaceTradersData("/systems/" + MyAgent.GetLocationSystem() + "/waypoints/" + MyAgent.GetLocationWaypoint())

	var responsePosition systems.Data

	json.Unmarshal(data, &responsePosition)

	return responsePosition
}

func getAgentStartContract() contracts.Data {
	data := space_traders.GetSpaceTradersData("/my/contracts")

	var responseContract contracts.Data

	json.Unmarshal(data, &responseContract)

	return responseContract
}

func postAgentNegociateContract() {
	data := space_traders.PostSpaceTradersData("/my/ships/JEANJEAN/negociate/contract", nil)
	fmt.Printf("> %s", data)
}

func (x Data) GetSymbol() string {
	return "Symbol : " + x.Data.Symbol
}

func (x Data) GetHeadquarter() string {
	return "Headquarter : " + x.Data.Headquarters
}

func (x Data) GetCredits() string {
	return fmt.Sprintf("Credits : %v ", x.Data.Credits)
}

func (x Data) GetFaction() string {
	return "Faction : " + x.Data.StartingFaction
}

func (x Data) GetFleet() string {
	return fmt.Sprintf("Fleet : %v", x.Data.ShipCount)
}

func (x Data) GetLocationSector() string {
	return x.parsingLocation()[0]
}

func (x Data) GetLocationSystem() string {
	return MyAgent.GetLocationSector() + "-" + x.parsingLocation()[1]
}

func (x Data) GetLocationWaypoint() string {
	return MyAgent.GetLocationSystem() + "-" + x.parsingLocation()[2]
}

func (x Data) parsingLocation() []string {
	return strings.Split(x.Data.Headquarters, "-")
}
