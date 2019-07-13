package core

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/datoga/sblm_cli/data"
	"gopkg.in/resty.v1"
)

//Requests encapsulates all the REST interaction with the endpoint
type Requests struct {
	endpoint *url.URL
}

//NewRequests creates a new struct to manage actuator requests
func NewRequests(endpoint *url.URL, verbose bool) *Requests {

	if verbose {
		resty.SetDebug(true)
	}

	return &Requests{
		endpoint: endpoint,
	}
}

//List make the list request for the actuator
func (requests Requests) List() (*data.ActuatorData, error) {
	resp, err := resty.R().Get(requests.endpoint.String())

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

//Edit make the list request for the actuator
func (requests Requests) Edit(logger string, loggerLevel data.LoggerLevel) error {
	endpoint := requests.endpoint.String() + "/" + logger

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
