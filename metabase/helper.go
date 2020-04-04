package metabase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

func (c *Client) getRequest(endpoint string, query url.Values) ([]byte, error) {

	req, err := c.newRequest("GET", endpoint, query, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("err in get command. endpoint: %s, code: %v", endpoint, resp.Status)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resData, nil
}

func (c *Client) postRequest(endpoint string, body []byte) ([]byte, error) {

	req, err := c.newRequest("POST", endpoint, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("err in post command. endpoint: %s, code: %v", endpoint, resp.Status)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resData, nil
}

func (c *Client) putRequest(endpoint string, body []byte) ([]byte, error) {

	req, err := c.newRequest("PUT", endpoint, nil, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("err in put command. endpoint: %s, code: %v", endpoint, resp.Status)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return resData, nil
}

func (c *Client) deleteRequest(endpoint string) error {

	req, err := c.newRequest("DELETE", endpoint, nil, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("err in delete command. endpoint: %s, code: %v", endpoint, resp.Status)
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	res := struct {
		Success bool `json:"success"`
	}{}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return err
	}

	return nil
}
