package metabase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
)

type Session struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) NewSession(user string, password string) (*string, error) {
	s := Session{
		User:     user,
		Password: password,
	}
	reqData, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("POST", "/api/session", nil, bytes.NewBuffer(reqData))
	if err != nil {
		return nil, err
	}

	resp, err := c.do(req)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	result := struct {
		ID string `json:"id"`
	}{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	c.Token = &(result.ID)
	return c.Token, err
}
