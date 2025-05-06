package space_traders

import (
	"bytes"
	"encoding/json"
	"net/url"
	"os"
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
		panic(err)
	}

	data := PostSpaceTradersRegister(endpoint, bytes.NewReader(body))

	var responseRegister dataRegister
	err = json.Unmarshal(data, &responseRegister)
	if err != nil {
		panic(err)
	}

	replaceTokenAgent([]byte(responseRegister.Data.Token))
}

func replaceTokenAgent(tokenAgent []byte) {
	err := os.Remove(".token-agent")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(".token-agent", tokenAgent, 0o600)
	if err != nil {
		panic(err)
	}
}
