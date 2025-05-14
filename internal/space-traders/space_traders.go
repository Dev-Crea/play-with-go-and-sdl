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

	"sdl/playing/internal/render/assets"

	"github.com/rs/zerolog"
)

var LoggerAPI *zerolog.Logger

func InitAPI(logger zerolog.Logger) {
	LoggerAPI = &logger
}

func GetSpaceTradersData(endpoint url.Values) []byte {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodGet, apiEndpoint(endpoint), nil)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Get data error")
	}

	headerAgent(request)
	return send(request)
}

func PostSpaceTradersData(endpoint url.Values, body io.Reader) []byte {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiEndpoint(endpoint), body)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Post data error")
	}

	headerAgent(request)
	return send(request)
}

func PostSpaceTradersRegister(endpoint url.Values, body io.Reader) []byte {
	request, err := http.NewRequestWithContext(context.Background(), http.MethodPost, apiEndpoint(endpoint), body)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Post authentication error")
	}

	headerAccount(request)
	return send(request)
}

func readTokenFile(file string) []byte {
	safeFile := filepath.Clean(file)
	LoggerAPI.Debug().Str("File", file).Msg("Read token ")
	content, err := os.ReadFile(safeFile)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Read Token file error")

		os.Exit(1)
	}

	return content
}

func headerAccount(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ReadFileTokenApp()))
	request.Header.Add("Content-Type", "application/json")
}

func headerAgent(request *http.Request) {
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", ReadFileTokenAgent()))
	request.Header.Add("Content-Type", "application/json")
}

func send(request *http.Request) []byte {
	time.Sleep(100 * time.Millisecond)
	client := http.Client{Timeout: 30 * time.Second}

	response, err := client.Do(request)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("failed request")
	}

	defer response.Body.Close() //nolint

	if response.StatusCode == http.StatusUnauthorized {
		LoggerAPI.Error().Stack().Int("Code", response.StatusCode).Msg("Error authentication")
		ErrorUnauthorized()
	}

	data, err := io.ReadAll(response.Body)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Parse response failed")
	}

	LoggerAPI.Debug().RawJSON("Data", data).Msgf("[%s] %s", request.Method, request.URL.String())

	return data
}

func apiEndpoint(endpoint url.Values) string {
	u, err := url.Parse(assets.SPACE_TRADER_API)
	if err != nil {
		LoggerAPI.Error().Stack().Err(err).Msg("Parse parameters failed")
	}
	return u.JoinPath(endpoint["api"]...).String()
}
