package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/datoga/sblm_cli/data"
	"gopkg.in/resty.v1"
)

//Client is a structure to manage context information about the client
type Client struct {
	endpoint *url.URL
	verbose  bool
}

//NewClient creates a new client to make requests to spring boot logger endpoints
func NewClient(endpoint *url.URL, verbose bool) *Client {

	if verbose {
		resty.SetDebug(true)
	}

	return &Client{
		endpoint: endpoint,
		verbose:  verbose,
	}
}

//ListLoggers gets all the available loggers published by the actuator which begins by the logger string
func (client Client) ListLoggers(filter string) (data.Loggers, error) {
	actuatorData, err := client.list()

	if err != nil {
		return nil, err
	}

	if filter == "" {
		return actuatorData.Loggers, nil
	}

	filteredLoggers := make(data.Loggers)

	for loggerName, level := range actuatorData.Loggers {
		if strings.HasPrefix(loggerName, filter) {
			filteredLoggers[loggerName] = level
		}
	}

	return filteredLoggers, nil
}

//EditLoggers sets all the available loggers published by the actuator which begins by the logger string with the provided level
func (client Client) EditLoggers(filter string, newLevel string) (int, error) {
	if !data.IsValidLevel(newLevel) {
		return -1, errors.New("The level " + newLevel + " is not valid")
	}

	loggers, err := client.ListLoggers(filter)

	if err != nil {
		return -1, err
	}

	if len(loggers) == 0 {
		return 0, nil
	}

	err = client.edit(filter, newLevel)

	if err != nil {
		return -1, nil
	}

	return len(loggers), nil
}

func (client Client) list() (*data.ActuatorData, error) {
	resp, err := resty.R().Get(client.endpoint.String())

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, errors.New("Error with status code " + strconv.Itoa(resp.StatusCode()))
	}

	actuatorData := data.ActuatorData{}

	err = json.Unmarshal(resp.Body(), &actuatorData)

	if err != nil {
		return nil, err
	}

	return &actuatorData, nil
}

func (client Client) edit(logger string, newLevel string) error {
	endpoint := client.endpoint.String() + "/" + logger

	loggerLevel := data.LoggerLevel{
		ConfiguredLevel: newLevel,
	}

	loggerLevelStr, err := json.Marshal(loggerLevel)

	if err != nil {
		return err
	}

	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(loggerLevelStr).
		Post(endpoint)

	if err != nil {
		return err
	}

	if resp.StatusCode() != http.StatusOK {
		return errors.New("Error with status code " + strconv.Itoa(resp.StatusCode()))
	}

	return nil
}
