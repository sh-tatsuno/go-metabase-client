package metabase

import (
	"encoding/json"
)

type Session struct {
	User     string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) CreateSession(user string, password string) (*string, error) {
	s := Session{
		User:     user,
		Password: password,
	}
	reqData, err := json.Marshal(&s)
	if err != nil {
		return nil, err
	}

	resData, err := c.postRequest("/api/session", reqData)
	if err != nil {
		return nil, err
	}

	res := struct {
		ID string `json:"id"`
	}{}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	c.Token = &(res.ID)
	return c.Token, err
}

func (c *Client) DeleteSession() error {

	err := c.deleteRequest("/api/session/")
	if err != nil {
		return err
	}

	return nil
}

// GET /api/session/password_reset_token_valid
// GET /api/session/properties
// POST /api/session/forgot_password
// POST /api/session/google_auth
// POST /api/session/reset_password
