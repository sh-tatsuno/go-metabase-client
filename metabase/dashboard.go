package metabase

import (
	"encoding/json"
	"fmt"
	"net/url"
)

// DELETE /api/dashboard/:dashboard-id/public_link
// DELETE /api/dashboard/:id
// DELETE /api/dashboard/:id/cards
// DELETE /api/dashboard/:id/favorite
// GET /api/dashboard/
// GET /api/dashboard/:id
// GET /api/dashboard/:id/related
// GET /api/dashboard/:id/revisions
// GET /api/dashboard/embeddable
// GET /api/dashboard/public
// POST /api/dashboard/
// POST /api/dashboard/:dashboard-id/public_link
// POST /api/dashboard/:from-dashboard-id/copy
// POST /api/dashboard/:id/cards
// POST /api/dashboard/:id/favorite
// POST /api/dashboard/:id/revert
// POST /api/dashboard/save
// POST /api/dashboard/save/collection/:parent-collection-id
// PUT /api/dashboard/:id
// PUT /api/dashboard/:id/cards

type Dashboard struct {
	ID                   int           `json:"id"`
	Name                 string        `json:"name"`
	Description          string        `json:"description,omitempty"`
	Archived             bool          `json:"archived,omitempty"`
	CollectionPosition   interface{}   `json:"collection_position,omitempty"`
	Creator              Creator       `json:"creator"`
	EnableEmbedding      bool          `json:"enable_embedding,omitempty"`
	CollectionID         interface{}   `json:"collection_id,omitempty"`
	ShowInGettingStarted bool          `json:"show_in_getting_started,omitempty"`
	Caveats              interface{}   `json:"caveats,omitempty"`
	CreatorID            int           `json:"creator_id"`
	UpdatedAt            string        `json:"updated_at"`
	MadePublicByID       interface{}   `json:"made_public_by_id,omitempty"`
	EmbeddingParams      interface{}   `json:"embedding_params,omitempty"`
	Position             interface{}   `json:"position,omitempty"`
	Parameters           []interface{} `json:"parameters"`
	Favorite             bool          `json:"favorite"`
	CreatedAt            string        `json:"created_at"`
	PublicUUID           interface{}   `json:"public_uuid,omitempty"`
	PointsOfInterest     interface{}   `json:"points_of_interest,omitempty"`
}

type Creator struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	FirstName   string `json:"first_name"`
	LastLogin   string `json:"last_login"`
	IsQbnewb    bool   `json:"is_qbnewb"`
	IsSuperuser bool   `json:"is_superuser"`
	LastName    string `json:"last_name"`
	DateJoined  string `json:"date_joined"`
	CommonName  string `json:"common_name"`
}

type DashboardRevision struct {
	ID          int    `json:"id"`
	Description string `json:"description, omitempty"`
	IsCreation  bool   `json:"is_creation"`
	IsReversion bool   `json:"is_reversion"`
	Timestamp   string `json:"timestamp"`
	User        struct {
		ID         int    `json:"id"`
		FirstName  string `json:"first_name"`
		LastName   string `json:"last_name"`
		CommonName string `json:"common_name"`
	} `json:"user"`
	Message string      `json:"message, omitempty"`
	Diff    interface{} `json:"diff"`
}

func (c *Client) GetDashboards(s string) ([]Dashboard, error) {

	query := url.Values{}

	// all, archived, mine
	types := []string{"all", "archived", "mine"}
	for _, t := range types {
		if s == t {
			query.Add("f", s)
			break
		}
	}

	resData, err := c.getRequest("/api/dashboard", query)
	if err != nil {
		return nil, err
	}

	res := []Dashboard{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (c *Client) GetDashboard(id int64) (*Dashboard, error) {

	resData, err := c.getRequest(fmt.Sprintf("/api/dashboard/%d", id), nil)
	if err != nil {
		return nil, err
	}

	res := Dashboard{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return &res, err
}

func (c *Client) GetDashboardRevisions(id int64) ([]DashboardRevision, error) {

	resData, err := c.getRequest(fmt.Sprintf("/api/dashboard/%d/revisions", id), nil)
	if err != nil {
		return nil, err
	}

	res := []DashboardRevision{}
	err = json.Unmarshal(resData, &res)
	if err != nil {
		return nil, err
	}

	return res, err
}
