package main

import (
	"echo-slam-client/backend/log"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Connector struct {
	client *http.Client

	Host string `json:"host"`
	JWT string `json:"jwt"`
	Connected bool `json:"connected"`
}

func (e *Connector) Self() *Connector {
	return e
}

func (e *Connector) getUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", e.Host, endpoint)
}

func (e *Connector) setParams(req *http.Request, params map[string]string) error {
	if params != nil {
		q, err := url.ParseQuery(req.URL.RawQuery)
		if err != nil {
			return err
		}
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	return nil
}

func (e *Connector) doRequestProtectedBasic(
	endpoint string, httpMethod string, params map[string]string,
	username string, password string,
	model interface{}) error {

	req, err := http.NewRequest(httpMethod, e.getUrl(endpoint), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)

	return e.doRequestMiddleware(req, model)
}

func (e *Connector) doRequestProtectedJWT(
	endpoint string, httpMethod string,
	params map[string]string, model interface{}) error {

	if !e.Connected {
		return errors.New("not connected")
	}

	req, err := http.NewRequest(httpMethod, e.getUrl(endpoint), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", e.JWT)

	err = e.setParams(req, params)

	return e.doRequestMiddleware(req, model)
}

func (e *Connector) doRequest(
	endpoint string, httpMethod string,
	params map[string]string, model interface{}) error {

	req, err := http.NewRequest(httpMethod, e.getUrl(endpoint), nil)
	if err != nil {
		return err
	}

	err = e.setParams(req, params)

	return e.doRequestMiddleware(req, model)
}

func (e *Connector) doRequestMiddleware(req *http.Request, model interface{}) error {

	err := e.doRequestBase(req, model)
	if err != nil {
		log.Error(err)
		log.Debug(fmt.Sprintf("%v", model))
	}

	return err
}

func (e *Connector) doRequestBase(req *http.Request, model interface{}) error {

	resp, err := e.client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return errors.New(
			fmt.Sprintf(
				"failed to make request to %s: %d",
				resp.Request.URL.RawQuery, resp.StatusCode))
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, model)
	if err != nil {
		return err
	}

	return nil
}
