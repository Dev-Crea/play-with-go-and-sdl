package space_traders

import (
	"bytes"
	"encoding/json"
	"net/url"
	"os"
)

// TODO : Change token manipulation
const (
	TOKEN_AGENT_PATH   = "../.token-agent"
	TOKEN_ACCOUNT_PATH = "../.token-account" //nolint
)

type requestRegister struct {
	Symbol  string `json:"symbol"`
	Faction string `json:"faction"`
}

type responseRegister struct {
	Token string `json:"token"`
}

type dataRegister struct {
	Data responseRegister
}

func ErrorUnauthorized() {
	endpoint := url.Values{}
	endpoint.Add("api", "register")

	requestRegister := requestRegister{
		Symbol:  "Roger2",
		Faction: "COBALT",
	}

	body, err := json.Marshal(requestRegister)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Fields(requestRegister).Msg("Create body error")
	}

	data := PostSpaceTradersRegister(endpoint, bytes.NewReader(body))

	var responseRegister dataRegister
	err = json.Unmarshal(data, &responseRegister)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Fields(responseRegister).Msg("Parse body error")
	}

	replaceTokenAgent([]byte(responseRegister.Data.Token))
}

func replaceTokenAgent(tokenAgent []byte) {
	err := os.Remove(TOKEN_AGENT_PATH)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Failed remove token agent")
	}

	err = os.WriteFile(TOKEN_AGENT_PATH, tokenAgent, 0o600)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Failed write token agent")
	}
}

func ReadFileTokenAgent() string {
	return string(readTokenFile(TOKEN_AGENT_PATH))
}

func ReadFileTokenApp() string {
	return string(readTokenFile(TOKEN_ACCOUNT_PATH))
}
