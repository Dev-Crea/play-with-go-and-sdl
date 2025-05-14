package models

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	space_traders "sdl/playing/internal/space-traders"
)

type AgentResponse struct {
	Symbol          string
	Headquarters    string
	Credits         int32
	StartingFaction string
	ShipCount       int32
	Location        Location
}

type AgentResponseData struct {
	Data AgentResponse
}

type Location struct {
	Sector   string
	System   string
	Waypoint string
}

var (
	MyAgent    AgentResponseData
	MyPosition SystemResponseData
	MyContract string
)

func InitAgent() {
	MyAgent = getAgentData()
	MyPosition = getAgentStartData()

	// getAgentStartContract()
	// postAgentNegociateContract()
}

func getAgentData() AgentResponseData {
	LoggerAPI.Debug().Msg("Get Agent data")

	endpoint := url.Values{}
	endpoint.Add("api", "my")
	endpoint.Add("api", "agent")

	data := space_traders.GetSpaceTradersData(endpoint)

	var responseAgent AgentResponseData

	err := json.Unmarshal(data, &responseAgent)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Error get agent data")
	}

	return responseAgent
}

func getAgentStartData() SystemResponseData {
	endpoint := url.Values{}
	endpoint.Add("api", "systems")
	endpoint.Add("api", MyAgent.GetLocationSystem())
	endpoint.Add("api", "waypoints")
	endpoint.Add("api", MyAgent.GetLocationWaypoint())
	data := space_traders.GetSpaceTradersData(endpoint)

	var responsePosition SystemResponseData

	err := json.Unmarshal(data, &responsePosition)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Error get agent start data")
	}

	return responsePosition
}

/*
func getAgentStartContract() ContractResponseData {
	endpoint := url.Values{}
	endpoint.Add("api", "my")
	endpoint.Add("api", "contracts")
	data := space_traders.GetSpaceTradersData(endpoint)

	var responseContract ContractResponseData

	err := json.Unmarshal(data, &responseContract)
	if err != nil {
		fmt.Printf("Error : %s", err)
		panic(err)
	}

	return responseContract
}
*/

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

func (x AgentResponseData) GetSymbol() string {
	return fmt.Sprintf("Symbol : %s", x.Data.Symbol)
}

func (x AgentResponseData) GetHeadquarter() string {
	return fmt.Sprintf("Headquarter : %s", x.Data.Headquarters)
}

func (x AgentResponseData) GetCredits() string {
	return fmt.Sprintf("Credits : %v ", x.Data.Credits)
}

func (x AgentResponseData) GetFaction() string {
	return fmt.Sprintf("Faction : %s", x.Data.StartingFaction)
}

func (x AgentResponseData) GetFleet() string {
	return fmt.Sprintf("Fleet : %v", x.Data.ShipCount)
}

func (x AgentResponseData) GetLocationSector() string {
	return x.parsingLocation()[0]
}

func (x AgentResponseData) GetLocationSystem() string {
	return fmt.Sprintf("%s-%s", MyAgent.GetLocationSector(), x.parsingLocation()[1])
}

func (x AgentResponseData) GetLocationWaypoint() string {
	return fmt.Sprintf("%s-%s", MyAgent.GetLocationSystem(), x.parsingLocation()[2])
}

func (x AgentResponseData) parsingLocation() []string {
	return strings.Split(x.Data.Headquarters, "-")
}
