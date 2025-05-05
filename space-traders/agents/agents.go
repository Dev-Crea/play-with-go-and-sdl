package agents

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"sdl/playing/space-traders/contracts"

	space_traders "sdl/playing/space-traders"

	systems "sdl/playing/space-traders/systems"
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

	// getAgentStartContract()
	// postAgentNegociateContract()
}

func getAgentData() Data {
	endpoint := url.Values{}
	endpoint.Add("api", "my")
	endpoint.Add("api", "agent")
	data := space_traders.GetSpaceTradersData(endpoint)

	var responseAgent Data

	err := json.Unmarshal(data, &responseAgent)
	if err != nil {
		panic(err)
	}

	return responseAgent
}

func getAgentStartData() systems.Data {
	endpoint := url.Values{}
	endpoint.Add("api", "systems")
	endpoint.Add("api", MyAgent.GetLocationSystem())
	endpoint.Add("api", "waypoints")
	endpoint.Add("api", MyAgent.GetLocationWaypoint())
	data := space_traders.GetSpaceTradersData(endpoint)

	var responsePosition systems.Data

	err := json.Unmarshal(data, &responsePosition)
	if err != nil {
		panic(err)
	}

	return responsePosition
}

func getAgentStartContract() contracts.Data {
	endpoint := url.Values{}
	endpoint.Add("api", "my")
	endpoint.Add("api", "contracts")
	data := space_traders.GetSpaceTradersData(endpoint)

	var responseContract contracts.Data

	err := json.Unmarshal(data, &responseContract)
	if err != nil {
		fmt.Printf("Error : %s", err)
		panic(err)
	}

	return responseContract
}

/*
func postAgentNegociateContract() {
	endpoint := url.Values{}
	endpoint.Add("api", "my")
	endpoint.Add("api", "ships")
	endpoint.Add("api", "")
	endpoint.Add("api", "negotiate")
	endpoint.Add("api", "contract")

	data := space_traders.PostSpaceTradersData(endpoint, nil)
	fmt.Printf("> %s", data)
}
*/

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
