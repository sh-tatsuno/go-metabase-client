package metabase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/url"
	"strconv"
)

type User struct {
	ID                   int         `json:"id"`
	Email                string      `json:"email"`
	LdapAuth             bool        `json:"ldap_auth"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	LastLogin            string      `json:"last_login"`
	IsActive             bool        `json:"is_active"`
	IsQbnewb             bool        `json:"is_qbnewb"`
	GroupIds             []int       `json:"group_ids"`
	IsSuperuser          bool        `json:"is_superuser"`
	LoginAttributes      interface{} `json:"login_attributes"`
	DateJoined           string      `json:"date_joined"`
	PersonalCollectionID int         `json:"personal_collection_id"`
	CommonName           string      `json:"common_name"`
	GoogleAuth           bool        `json:"google_auth"`
	UpdatedAt            string      `json:"updated_at"`
}

type UpdateUser struct {
	ID              int         `json:"id"`
	Email           *string     `json:"email"`
	FirstName       *string     `json:"first_name"`
	LastName        *string     `json:"last_name"`
	GroupIds        []int       `json:"group_ids"`
	IsSuperuser     *bool       `json:"is_superuser"`
	LoginAttributes interface{} `json:"login_attributes"`
}

func (c *Client) DeleteUser(id int) (*bool, error) {

	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/user/%d", id), nil, nil)
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
		Success bool `json:"success"`
	}{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return &result.Success, err
}

func (c *Client) GetUsers(includeDeactivated bool) ([]User, error) {

	query := url.Values{}

	// valid only when includeDeactivated is true
	if includeDeactivated:
		query.Add("include_deactivated", strconv.FormatBool(includeDeactivated))

	req, err := c.newRequest("GET", "/api/user", query, nil)
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
	result := []User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Client) GetUser(id int) (*User, error) {

	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%d", id), nil, nil)
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
	result := User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c *Client) GetCurrentUser() (*User, error) {

	req, err := c.newRequest("GET", "/api/user/current", nil, nil)
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
	result := User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return &result, err
}

func (c *Client) CreateUser(
	firstName string,
	lastName string,
	email string,
	password string,
	groupIDs []int,
	loginAttributes interface{},
) (*User, error) {

	d := struct {
		firstName       string      `json:"first_name"`
		lastName        string      `json:"last_name"`
		email           string      `json:"email"`
		password        string      `json:"password"`
		groupIDs        []int       `json:"group_ids"`
		loginAttributes interface{} `json:"login_attributes"`
	}{
		firstName,
		lastName,
		email,
		password,
		groupIDs,
		loginAttributes,
	}
	reqData, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("POST", "/api/user", nil, bytes.NewBuffer(reqData))
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
	result := &User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Client) SendInvite(id int) (*bool, error) {

	req, err := c.newRequest("POST", fmt.Sprintf("/api/user/%d/send_invite", id), nil, nil)
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
		Success bool `json:"success"`
	}{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return &result.Success, err
}

func (c *Client) UpdateUser(u UpdateUser) (*User, error) {

	reqData, err := json.Marshal(&u)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("PUT", "/api/user", nil, bytes.NewBuffer(reqData))
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
	result := &User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Client) UpdatePassword(
	id string,
	password string,
	oldPassword string,
) (*User, error) {

	d := struct {
		ID          string `json:"id"`
		Password    string `json:"password"`
		OldPassword string `json:"old_password"`
	}{
		id,
		password,
		oldPassword,
	}
	reqData, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}

	req, err := c.newRequest("PUT", fmt.Sprintf("/api/user/%d/password", d.ID), nil, bytes.NewBuffer(reqData))
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
	result := &User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}

func (c *Client) Qbnewb(id int) (*bool, error) {

	req, err := c.newRequest("PUT", fmt.Sprintf("/api/user/%d/qbnewb", id), nil, nil)
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
		Success bool `json:"success"`
	}{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return &result.Success, err
}

func (c *Client) Reactive(id int) (*User, error) {

	req, err := c.newRequest("PUT", fmt.Sprintf("/api/user/%d/reactive", id), nil, nil)
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

	result := &User{}

	err = json.Unmarshal(respData, &result)
	if err != nil {
		return nil, err
	}

	return result, err
}
