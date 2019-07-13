package core

import (
	"errors"
	"net/url"
	"strings"

	"github.com/datoga/sblm_cli/data"
)

//Client is a structure to manage context information about the client
type Client struct {
	requests *Requests
	verbose  bool
}

//NewClient creates a new client to make requests to spring boot logger endpoints
func NewClient(endpoint *url.URL, verbose bool) *Client {

	requests := NewRequests(endpoint, verbose)

	return &Client{
		requests: requests,
	}
}

//ListLoggers gets all the available loggers published by the actuator which begins by the logger string
func (client Client) ListLoggers(filter string) (data.Loggers, error) {
	actuatorData, err := client.requests.List()

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

	loggerLevel := data.LoggerLevel{
		ConfiguredLevel: newLevel,
	}

	err = client.requests.Edit(filter, loggerLevel)

	if err != nil {
		return -1, nil
	}

	return len(loggers), nil
}
