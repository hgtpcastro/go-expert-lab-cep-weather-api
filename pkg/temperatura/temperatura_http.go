package temperatura

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/hgtpcastro/go-expert-lab-cep-weather-api/configs"
)

type temperatura struct {
	config *configs.Config
}

func NewTemperatura(config *configs.Config) *temperatura {
	return &temperatura{config: config}
}

func (t *temperatura) BuscaTemperatura(localidade string) (float64, error) {
	params := url.Values{}
	params.Add("q", strings.ToLower(localidade))

	url := fmt.Sprintf(t.config.WeatherApiUrl, params.Encode())

	log.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data TemperaturaDto
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.Current.TempC, nil
}
