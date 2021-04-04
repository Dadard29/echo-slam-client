package backend

import (
	"echo-slam-client/backend/log"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type ApiClient struct {
	client *http.Client

	host      string
	JWT       string
	connected bool
}

func NewApiClient(host string, jwt string, connected bool) *ApiClient {
	return &ApiClient{
		client:    &http.Client{},
		host:      host,
		JWT:       jwt,
		connected: connected,
	}
}

func (c *ApiClient) Host() string {
	return c.host
}

func (c *ApiClient) Connected() bool {
	return c.connected
}

func (c *ApiClient) Connect(JWT string) {
	c.JWT = JWT
	c.connected = true
}

func (c *ApiClient) Disconnect() {
	c.JWT = ""
	c.connected = false
}

func (c *ApiClient) getUrl(endpoint string) string {
	return fmt.Sprintf("%s%s", c.host, endpoint)
}

func (c *ApiClient) setParams(req *http.Request, params map[string]string) error {
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

func (c *ApiClient) DoRequestProtectedBasic(
	endpoint string, httpMethod string, params map[string]string,
	username string, password string,
	model interface{}) error {

	req, err := http.NewRequest(httpMethod, c.getUrl(endpoint), nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(username, password)

	return c.doRequestMiddleware(req, model)
}

func (c *ApiClient) DoRequestProtectedJWT(
	endpoint string, httpMethod string,
	params map[string]string, model interface{}) error {

	if !c.connected {
		return errors.New("not connected")
	}

	req, err := http.NewRequest(httpMethod, c.getUrl(endpoint), nil)
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", c.JWT)

	err = c.setParams(req, params)

	return c.doRequestMiddleware(req, model)
}

func (c *ApiClient) DoRequest(
	endpoint string, httpMethod string,
	params map[string]string, model interface{}) error {

	req, err := http.NewRequest(httpMethod, c.getUrl(endpoint), nil)
	if err != nil {
		return err
	}

	err = c.setParams(req, params)

	return c.doRequestMiddleware(req, model)
}

func (c *ApiClient) doRequestMiddleware(req *http.Request, model interface{}) error {

	err := c.doRequestBase(req, model)
	if err != nil {
		log.Error(err)
		log.Debug(fmt.Sprintf("%v", model))
	}

	return err
}

func (c *ApiClient) doRequestBase(req *http.Request, model interface{}) error {

	resp, err := c.client.Do(req)
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
