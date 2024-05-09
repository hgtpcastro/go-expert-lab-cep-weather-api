package httphandlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hgtpcastro/go-expert-lab-cep-weather-api/configs"
	cep1 "github.com/hgtpcastro/go-expert-lab-cep-weather-api/pkg/cep"
	"github.com/hgtpcastro/go-expert-lab-cep-weather-api/pkg/temperatura"
)

type cepWeatherHandler struct {
	config *configs.Config
}

func NewCepWeatherHandler(config *configs.Config) *cepWeatherHandler {
	return &cepWeatherHandler{config: config}
}

func (h *cepWeatherHandler) Handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cep := r.URL.Query().Get("cep")
	if len(cep) != 8 {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprint(w, "invalid zipcode")
		return
	}

	cepResponse, err := cep1.NewBuscaCepHttp(h.config).BuscaCep(cep)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "can not find zipcode")
		return
	}

	if cepResponse.Localidade == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "can not find zipcode")
		return
	}

	log.Printf("localidade: %s", (cepResponse.Localidade))

	tempC, err := temperatura.NewTemperatura(h.config).BuscaTemperatura(cepResponse.Localidade)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "error getting temperature")
		return
	}

	tempF := temperatura.NewConversor().CelsiusToFahrenheit(tempC)
	tempK := temperatura.NewConversor().CelsiusToKelvin(tempC)

	response := map[string]float64{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
