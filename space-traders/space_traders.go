package space_traders

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"sdl/playing/constants"
)

func readTokenFile(file string) []byte {
	safeFile := filepath.Clean(file)
	content, err := os.ReadFile(safeFile)
	if err != nil {
		panic(err)
	}

	return content
}

func readFileTokenAgent() string {
	return string(readTokenFile(".token-agent"))
}

func readFileTokenApp() string {
	return string(readTokenFile(".token-account"))
}

// TODO Register agent
// Create form when server is reject token agent
/*
func readFileTokenGame() string {
	return string(readTokenFile(".token-agent"))
}

func CreateOrReadTokenAgent() {
	requestURL := "https://api.spacetraders.io/v2/register"
	// requestData := io.Reader("{'symbol': 'JEAN_JEAN', 'faction': 'AEGIS'}")
	requestData := url.Values{}
	requestData.Set("faction", "AEGIS")
	requestData.Set("symbol", "JEANJEAN")

	client := &http.Client{}
	req, err := http.NewRequest(http.MethodPost, requestURL, strings.NewReader(requestData.Encode()))
	if err != nil {
		panic(err)
	}

	req.Header.Add("Authorisation", "Bearer %s"+readFileTokenGame)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Status register %s", resp.Status)
}
*/
func headerAccount(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", readFileTokenApp()))
	request.Header.Add("Content-Type", "application/json")
}

func headerAgent(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", readFileTokenAgent()))
	request.Header.Add("Content-Type", "application/json")
}

func send(request *http.Request) []byte {
	time.Sleep(100 * time.Millisecond)
	client := http.Client{Timeout: 30 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close() //nolint

	if response.StatusCode == http.StatusUnauthorized {
		ErrorUnauthorized()
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("[%s] %s : %s\n", request.Method, request.URL.String(), data)

	return data
}

func apiEndpoint(endpoint url.Values) string {
	u, err := url.Parse(constants.SPACE_TRADER_API)
	if err != nil {
		panic(err)
	}
	return u.JoinPath(endpoint["api"]...).String()
}

func GetSpaceTradersData(endpoint url.Values) []byte {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiEndpoint(endpoint), nil)
	if err != nil {
		panic(err)
	}

	headerAgent(request)
	return send(request)
}

func PostSpaceTradersData(endpoint url.Values, body io.Reader) []byte {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiEndpoint(endpoint), body)
	if err != nil {
		panic(err)
	}

	headerAgent(request)
	return send(request)
}

func PostSpaceTradersRegister(endpoint url.Values, body io.Reader) []byte {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiEndpoint(endpoint), body)
	if err != nil {
		panic(err)
	}

	headerAccount(request)
	return send(request)
}
