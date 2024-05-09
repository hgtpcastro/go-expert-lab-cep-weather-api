package cep

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/hgtpcastro/go-expert-lab-cep-weather-api/configs"
)

type cepHttp struct {
	config *configs.Config
}

func NewBuscaCepHttp(config *configs.Config) *cepHttp {
	return &cepHttp{config: config}
}

func (b *cepHttp) BuscaCep(cep string) (*CepDto, error) {
	resp, err := http.Get(fmt.Sprintf(b.config.CepApiUrl, cep))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data CepDto
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
