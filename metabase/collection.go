package metabase

import (
	"encoding/json"
	"fmt"
)

type Collection struct {
	ID                 json.Number         `json:"id"`
	Description        string              `json:"description,omitempty"`
	Archived           bool                `json:"archived,omitempty"`
	Slug               string              `json:"slug,omitempty"`
	Color              string              `json:"color,omitempty"` // TODO: validate color ^#[0-9A-Fa-f]{6}$
	CanWrite           bool                `json:"can_write"`
	Name               string              `json:"name"`
	PersonalOwnerID    int64               `json:"personal_owner_id"`
	EffectiveAncestors []EffectiveAncestor `json:"effective_ancestors"`
	EffectiveLocation  string              `json:"effective_location,omitempty"`
	ParentID           json.Number         `json:"parent_id,omitempty"`
	Location           string              `json:"location"`
}

type EffectiveAncestor struct {
	ID                             json.Number `json:"id"`
	MetabaseModelsCollectionIsRoot bool        `json:"metabase.models.collection/is-root?"`
	Name                           string      `json:"name"`
	CanWrite                       bool        `json:"can_write"`
}

type CollectionPatch struct {
	ID          json.Number `json:"id"`
	Name        string      `json:"name,omitempty"`
	Description string      `json:"description,omitempty"`
	Archived    bool        `json:"archived,omitempty"`
	Color       string      `json:"color,omitempty"`
	ParentID    int64       `json:"parent_id,omitempty"`
}

type CollectionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Color       string `json:"color,omitempty"`
	ParentID    int64  `json:"parent_id,omitempty"`
}

type CollectionItem struct {
	ID                 int64       `json:"id"`
	Name               string      `json:"name"`
	Description        string      `json:"description"`
	CollectionPosition interface{} `json:"collection_position,omitempty"` // TODO: change to appropriate struct
	Model              string      `json:"model"`
	CanWrite           bool        `json:"can_write,omitempty"`
}

type CollenctionGraphPermission struct {
	Revision int64                                     `json:"revision"`
	Groups   map[string]CollectionGraphPermissionGroup `json:"groups"`
}

type CollectionGraphPermissionGroup struct {
	Root string `json:"root"`
}

func (c *Client) GetCollenctions() ([]Collection, error) {

	resData, err := c.getRequest("/api/collection", nil)
	if err != nil {
		return nil, err
	}

	res := []Collection{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) GetCollenction(id int64) (*Collection, error) {

	resData, err := c.getRequest(fmt.Sprintf("/api/collection/%d", id), nil)
	if err != nil {
		return nil, err
	}

	res := Collection{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) GetRootCollenction() (*Collection, error) {

	resData, err := c.getRequest("/api/collection/root", nil)
	if err != nil {
		return nil, err
	}

	res := Collection{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) GetCollenctionItems(id int64) ([]CollectionItem, error) {

	resData, err := c.getRequest(fmt.Sprintf("/api/collection/%d/items", id), nil)
	if err != nil {
		return nil, err
	}

	res := []CollectionItem{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) GetRootCollenctionItems() ([]CollectionItem, error) {

	resData, err := c.getRequest("/api/collection/root/items", nil)
	if err != nil {
		return nil, err
	}

	res := []CollectionItem{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) GetCollenctionGraphPermission() (*CollenctionGraphPermission, error) {

	resData, err := c.getRequest("/api/collection/graph", nil)
	if err != nil {
		return nil, err
	}

	res := CollenctionGraphPermission{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) CreateCollenction(cp CollectionRequest) (*Collection, error) {

	reqData, err := json.Marshal(&cp)
	if err != nil {
		return nil, err
	}

	resData, err := c.postRequest("/api/collection", reqData)
	if err != nil {
		return nil, err
	}

	res := Collection{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) UpdateCollenction(cp CollectionPatch) (*Collection, error) {

	reqData, err := json.Marshal(&cp)
	if err != nil {
		return nil, err
	}

	resData, err := c.postRequest("/api/collection", reqData)
	if err != nil {
		return nil, err
	}

	res := Collection{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) UpdateCollenctionGraphPermission(p CollenctionGraphPermission) (*CollenctionGraphPermission, error) {

	reqData, err := json.Marshal(&p)
	if err != nil {
		return nil, err
	}

	resData, err := c.putRequest("/api/collection/graph", reqData)
	if err != nil {
		return nil, err
	}

	res := CollenctionGraphPermission{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}
