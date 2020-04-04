package metabase

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	MemberCount int64  `json:"member_count"`
}

type Membership struct {
	MembershipID int64 `json:"membership_id"`
	GroupID      int64 `json:"group_id"`
	UserID       int64 `json:"user_id"`
}

func (c *Client) DeletePermissionsGroup(groupID int64) error {

	err := c.deleteRequest(fmt.Sprintf("/api/permissions/group/%d", groupID))
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeletePermissionsMembership(membershipID int64) error {

	err := c.deleteRequest(fmt.Sprintf("/api/permissions/membership/%d", membershipID))
	if err != nil {
		return err
	}

	return nil
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

func (c *Client) GetPermissionsGroup(groupID int64) (*Group, error) {

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

func (c *Client) CreatePermissionsMembership(groupID int64, userID int64) (*Membership, error) {

	d := struct {
		GroupID int64 `json:"group_id"`
		UserID  int64 `json:"user_id"`
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

func (c *Client) UpdatePermissionsGroup(groupID int64, groupName string) (*Group, error) {

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

// ================
// Graph Permission
// ================

type PermissionsGraph struct {
	Revision int64                               `json:"revision"`
	Groups   map[string](map[string]interface{}) `json:"groups"`
}

type BulkPermission struct {
	Native  string `json:"native,omitempty"`
	Schemas string `json:"schemas,omitempty"`
}

type StepPermission struct {
	Native  string                         `json:"native,omitempty"`
	Schemas map[string](map[string]string) `json:"schemas,omitempty"`
}

func (p *PermissionsGraph) UnmarshalJSON(data []byte) error {
	type Alias PermissionsGraph
	a := &struct {
		Groups map[string](map[string]json.RawMessage) `json:"groups"`
		*Alias
	}{
		Alias: (*Alias)(p),
	}

	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	m := map[string](map[string]interface{}){}
	for groupKey, groupValue := range a.Groups {
		graphMap := map[string]interface{}{}
		for graphKey, graphValue := range groupValue {
			var s StepPermission
			if err := json.Unmarshal(graphValue, &s); err == nil {
				graphMap[graphKey] = &s
			} else {
				var b BulkPermission
				if err := json.Unmarshal(graphValue, &b); err == nil {
					graphMap[graphKey] = &b
				} else {
					return err
				}
			}

		}
		m[groupKey] = graphMap
	}
	p.Groups = m

	return nil
}

func (c *Client) GetPermissionsGraphs() (*PermissionsGraph, error) {
	resData, err := c.getRequest("/api/permissions/graph", nil)
	if err != nil {
		return nil, err
	}

	res := PermissionsGraph{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdatePermissionsGraph(p PermissionsGraph) (*PermissionsGraph, error) {

	reqData, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	resData, err := c.putRequest("/api/permissions/graph", reqData)
	if err != nil {
		return nil, err
	}

	res := PermissionsGraph{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
