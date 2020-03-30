package metabase

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MemberCount int    `json:"member_count"`
}

type Membership struct {
	MembershipID int `json:"membership_id"`
	GroupID      int `json:"group_id"`
	UserID       int `json:"user_id"`
}

func (c *Client) DeletePermissionsGroup(groupID int) (*bool, error) {

	res, err := c.deleteRequest(fmt.Sprintf("/api/permissions/group/%d", groupID))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) DeletePermissionsMembership(membershipID int) (*bool, error) {

	res, err := c.deleteRequest(fmt.Sprintf("/api/permissions/membership/%d", membershipID))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetPermissionsGroups() ([]Group, error) {

	resData, err := c.getRequest("/api/permissions/group", nil)
	if err != nil {
		return nil, err
	}

	res := []Group{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) GetPermissionsGroup(groupID int) (*Group, error) {

	resData, err := c.getRequest(fmt.Sprintf("/api/permissions/group/%d", groupID), nil)
	if err != nil {
		return nil, err
	}

	res := Group{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetPermissionsMemberships() (map[string][]Membership, error) {

	resData, err := c.getRequest("/api/permissions/membership", nil)
	if err != nil {
		return nil, err
	}

	res := map[string][]Membership{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) CreatePermissionsMembership(groupID int, userID int) (*Membership, error) {

	d := struct {
		GroupID int `json:"group_id"`
		UserID  int `json:"user_id"`
	}{
		groupID,
		userID,
	}
	reqData, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}

	resData, err := c.postRequest("/api/permissions/membership", reqData)
	if err != nil {
		return nil, err
	}

	res := Membership{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreatePermissionsGroup(groupName string) (*Group, error) {

	d := struct {
		Name string `json:"name"`
	}{
		groupName,
	}
	reqData, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}

	resData, err := c.postRequest("/api/permissions/group", reqData)
	if err != nil {
		return nil, err
	}

	res := Group{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdatePermissionsGroup(groupID int, groupName string) (*Group, error) {

	d := struct {
		Name string `json:"name"`
	}{
		groupName,
	}
	reqData, err := json.Marshal(&d)
	if err != nil {
		return nil, err
	}

	resData, err := c.putRequest(fmt.Sprintf("/api/permissions/group/%d", groupID), reqData)
	if err != nil {
		return nil, err
	}

	res := Group{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

// GET /api/permissions/graph
// PUT /api/permissions/graph
