package hass

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"polybar-lamp/config"
	"net/http"
	url2 "net/url"
	"path"
)

type APIClient struct {
	basePath *url2.URL
	apiKey   string
}

func (client APIClient) callService(serviceName string, data interface{}) error {
	servicePath := *client.basePath // make a copy
	servicePath.Path = path.Join(servicePath.Path, "services/light/"+serviceName)

	bodybuffer := makeBodyBuffer(data)

	request, err := http.NewRequest("POST", servicePath.String(), bodybuffer)

	if err != nil {
		panic(err)
	}

	addBearerTokenHeader(request, client.apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return err
	}
	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return fmt.Errorf("Service Call was not successful:\nResponse Body:\n%s", string(body))
	}
	return nil
}

func (client APIClient) GetState(entityID string) (LampState, error) {
	servicePath := *client.basePath // make a copy
	servicePath.Path = path.Join(servicePath.Path, "states", entityID)

	request, err := http.NewRequest("GET", servicePath.String(), nil)

	if err != nil {
		panic(err)
	}

	addBearerTokenHeader(request, client.apiKey)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return LampState{}, err
	}
	if response.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(response.Body)
		return LampState{}, fmt.Errorf("State Query was not successful:\nResponse Body:\n%s", string(body))
	}

	lampState := new(LampState)
	err = json.NewDecoder(response.Body).Decode(&lampState)

	return *lampState, nil
}

func (client APIClient) callTurnOn(data TurnOnData) error {
	return client.callService("turn_on", data)
}

func (client APIClient) callTurnOff(data TurnOffData) error {
	return client.callService("turn_off", data)
}

func (client APIClient) TurnOn(lampEntity string) error {
	return client.callTurnOn(TurnOnData{EntityID: lampEntity})
}

func (client APIClient) SetBrightness(lampEntity string, brightness int) error {
	return client.callTurnOn(TurnOnData{EntityID: lampEntity, BrightnessPercent: &brightness})
}

func (client APIClient) SetColorTemp(lampEntity string, colorTemp int) error {
	return client.callTurnOn(TurnOnData{EntityID: lampEntity, ColorTemp: &colorTemp})
}

func (client APIClient) TurnOff(lampEntity string) error {
	return client.callTurnOff(TurnOffData{EntityID: lampEntity})
}

func NewAPIClient(config config.PolybarLampConfig) *APIClient {
	basePath, err := url2.Parse(config.ApiBaseUrl)
	if err != nil {
		panic(err)
	}
	return &APIClient{basePath: basePath, apiKey: config.ApiKey}
}
