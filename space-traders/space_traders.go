package space_traders

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

const api = "https://api.spacetraders.io/v2"

func readTokenFile(file string) []byte {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	// fmt.Println("File : %s", string(content))
	return content
}

func readFileTokenAgent() string {
	return string(readTokenFile(".token-agent"))
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

func header(request http.Request) {
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

	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	return data
}

func GetSpaceTradersData(endpoint string) []byte {
	request, err := http.NewRequest(http.MethodGet, api+endpoint, nil)
	if err != nil {
		panic(err)
	}

	header(*request)
	data := send(request)

	fmt.Printf("Request [GET]: %s%s\n", api, endpoint)
	fmt.Printf("Data : %s\n", data)

	return data
}

func PostSpaceTradersData(endpoint string, body io.Reader) []byte {
	request, err := http.NewRequest(http.MethodPost, api+endpoint, body)
	if err != nil {
		panic(err)
	}

	header(*request)
	data := send(request)

	fmt.Printf("Request [POST]: %s%s\n", api, endpoint)
	fmt.Printf("Data : %s\n", data)

	return data
}
