package gitlab

import (
	"net/http"
	"net/url"
)

// Label represents a GitLab label.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/labels.html
type Label struct {
	ID                    int64  `json:"id"`
	Color                 string `json:"color"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	OpenIssueCount        int    `json:"open_issue_count"`
	ClosedIssueCount      int    `json:"closed_issue_count"`
	OpenMergeRequestCount int    `json:"open_merge_requests_count"`
	Subscribed            bool   `json:"subscribed"`
	Priority              int    `json:"priority"`
}

// LabelRequest represents the available CreateLabel() and UpdateLabel() options.
type LabelRequest struct {
	Color   string `json:"color"`
	Name    string `json:"name"`
	NewName string `json:"new_name,omitempty"`
}

type LabelDeleteOptions struct {
	Name string `url:"name,omitempty"`
}

// ListLabels gets all labels for given project.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/labels.html#list-labels
func (g *Client) ListLabels(project_id string, o *ListOptions) ([]*Label, error) {
	path := getUrl([]string{"projects", url.QueryEscape(project_id), "labels"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Label
	if _, err := g.Do(req, &ret); err != nil {
		return nil, err
	}

	return ret, nil
}

// CreateIssue creates a new project issue.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#new-issues
func (g *Client) CreateLabel(project_id string, label *LabelRequest) (*Label, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "labels"}
	req, _ := g.NewRequest("POST", path, label)

	var ret *Label
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}

// EditLabel updates an existing project labels
//
// GitLab API docs: http://doc.gitlab.com/ce/api/labels.html#edit-an-existing-label
func (g *Client) EditLabel(project_id string, label *LabelRequest) (*Label, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "labels"}
	req, _ := g.NewRequest("PUT", path, label)

	var ret *Label
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}

//DeleteLabel deletes an existing project label
//
// GitLab API docs: http://doc.gitlab.com/ce/api/labels.html#delete-a-label
func (g *Client) DeleteLabel(project_id string, label *LabelRequest) (*Label, *http.Response, error) {
	o := &LabelDeleteOptions{Name: label.Name}
	path := []string{"projects", url.QueryEscape(project_id), "labels"}

	u, _ := addOptions(getUrl(path), o)

	req, _ := http.NewRequest("DELETE", u, nil)

	var ret *Label

	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}
