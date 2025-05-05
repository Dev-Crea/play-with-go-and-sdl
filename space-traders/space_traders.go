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

	"sdl/playing/assets"
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
	u, err := url.Parse(assets.SPACE_TRADER_API)
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
