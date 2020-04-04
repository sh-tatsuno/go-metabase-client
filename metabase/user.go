package metabase

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type User struct {
	ID                   int64       `json:"id"`
	Email                string      `json:"email"`
	LdapAuth             bool        `json:"ldap_auth"`
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	LastLogin            string      `json:"last_login"`
	IsActive             bool        `json:"is_active"`
	IsQbnewb             bool        `json:"is_qbnewb"`
	GroupIds             []int64     `json:"group_ids"`
	IsSuperuser          bool        `json:"is_superuser"`
	LoginAttributes      interface{} `json:"login_attributes"` // TODO: change to appropriate struct
	DateJoined           string      `json:"date_joined"`
	PersonalCollectionID int64       `json:"personal_collection_id"`
	CommonName           string      `json:"common_name"`
	GoogleAuth           bool        `json:"google_auth"`
	UpdatedAt            string      `json:"updated_at"`
}

type UserPatch struct {
	ID              int64       `json:"id"`
	Email           *string     `json:"email"`
	FirstName       *string     `json:"first_name"`
	LastName        *string     `json:"last_name"`
	GroupIds        []int64     `json:"group_ids"`
	IsSuperuser     *bool       `json:"is_superuser"`
	LoginAttributes interface{} `json:"login_attributes"` // TODO: change to appropriate struct
}

func (c *Client) DeleteUser(id int64) error {

	err := c.deleteRequest(fmt.Sprintf("/api/user/%d", id))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) GetUsers(includeDeactivated bool) ([]User, error) {

	query := url.Values{}

	// valid only when includeDeactivated is true
	if includeDeactivated {
		query.Add("include_deactivated", strconv.FormatBool(includeDeactivated))
	}

	resData, err := c.getRequest("/api/user", query)
	if err != nil {
		return nil, err
	}

	res := []User{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) GetUser(id int64) (*User, error) {

	resData, err := c.getRequest(fmt.Sprintf("/api/user/%d", id), nil)
	if err != nil {
		return nil, err
	}

	res := User{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) GetCurrentUser() (*User, error) {

	resData, err := c.getRequest("/api/user/current", nil)
	if err != nil {
		return nil, err
	}

	res := User{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) CreateUser(
	firstName string,
	lastName string,
	email string,
	password string,
	groupIDs []int64,
	loginAttributes interface{},
) (*User, error) {

	d := struct {
		FirstName       string      `json:"first_name"`
		LastName        string      `json:"last_name"`
		Email           string      `json:"email"`
		Password        string      `json:"password"`
		GroupIDs        []int64     `json:"group_ids"`
		LoginAttributes interface{} `json:"login_attributes"`
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

	resData, err := c.postRequest("/api/user", reqData)
	if err != nil {
		return nil, err
	}

	res := &User{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) SendInvite(id int64) (*bool, error) {

	resData, err := c.postRequest(fmt.Sprintf("/api/user/%d/send_invite", id), nil)
	if err != nil {
		return nil, err
	}

	res := struct {
		Success bool `json:"success"`
	}{}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res.Success, err
}

func (c *Client) UserPatch(u UserPatch) (*User, error) {

	reqData, err := json.Marshal(&u)
	if err != nil {
		return nil, err
	}

	resData, err := c.postRequest("/api/user", reqData)
	if err != nil {
		return nil, err
	}
	res := &User{}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
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

	resData, err := c.putRequest(fmt.Sprintf("/api/user/%d/password", d.ID), reqData)
	if err != nil {
		return nil, err
	}

	res := &User{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) Qbnewb(id int64) (*bool, error) {

	resData, err := c.putRequest(fmt.Sprintf("/api/user/%d/qbnewb", id), nil)
	if err != nil {
		return nil, err
	}

	res := struct {
		Success bool `json:"success"`
	}{}

	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res.Success, err
}

func (c *Client) Reactive(id int64) (*User, error) {

	resData, err := c.putRequest(fmt.Sprintf("/api/user/%d/reactive", id), nil)
	if err != nil {
		return nil, err
	}

	res := &User{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}
