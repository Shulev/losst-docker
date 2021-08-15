package gitlab

import (
	"net/http"
	"net/url"
	"time"
)

// Issue represents a GitLab issue.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html
type Issue struct {
	Assignee       *User      `json:"assignee"`
	Author         *User      `json:"author"`
	Description    string     `json:"description"`
	Milestone      *Milestone `json:"milestone"`
	Id             int64      `json:"id"`
	Iid            int64      `json:"iid"`
	Labels         *[]string  `json:"labels"`
	ProjectId      int64      `json:"project_id"`
	State          string     `json:"state"`
	Title          string     `json:"title"`
	UserNotesCount int        `json:"user_notes_count"`
	Subscribed     bool       `json:"subscribed"`
	CreatedAt      time.Time  `json:"created_at"`
	UpdatedAt      time.Time  `json:"updated_at"`
	DueDate        string     `json:"due_date"`
	Confidential   bool       `json:"confidential"`
	WebURL         string     `json:"web_url"`
}

// IssueRequest represents the available CreateIssue() and UpdateIssue() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#new-issues
type IssueRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	AssigneeId  *int64 `json:"assignee_id"`
	MilestoneId *int64 `json:"milestone_id"`
	Labels      string `json:"labels"`
	StateEvent  string `json:"state_event,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
}

// MoveIssueRequest moved issue to another project
type MoveIssueRequest struct {
	ToProjectID string `json:"to_project_id"`
}

// ListIssuesOptions represents the available ListIssues() options.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#list-issues
type IssueListOptions struct {
	// State filters issues based on their state.  Possible values are: open,
	// closed.  Default is "open".
	State string `url:"state,omitempty"`

	ListOptions
}

// ListIssues gets all issues created by authenticated user. This function
// takes pagination parameters page and per_page to restrict the list of issues.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#list-issues
func (g *Client) ListIssues(projectID string, o *IssueListOptions) ([]*Issue, *CollectionOptions, error) {
	path := getUrl([]string{"projects", url.QueryEscape(projectID), "issues"})
	u, err := addOptions(path, o)
	if err != nil {
		return nil, nil, err
	}

	req, _ := http.NewRequest("GET", u, nil)

	var ret []*Issue
	resp, err := g.Do(req, &ret)
	if err != nil {
		return nil, nil, err
	}

	return ret, NewCollectionOption(resp), nil
}

// CreateIssue creates a new project issue.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#new-issues
func (g *Client) CreateIssue(project_id string, issue *IssueRequest) (*Issue, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues"}
	req, _ := g.NewRequest("POST", path, issue)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}

// UpdateIssue updates an existing project issue. This function is also used
// to mark an issue as closed.
//
// GitLab API docs: http://doc.gitlab.com/ce/api/issues.html#edit-issues
func (g *Client) UpdateIssue(project_id, issue_id string, issue *IssueRequest) (*Issue, *http.Response, error) {
	path := []string{"projects", url.QueryEscape(project_id), "issues", issue_id}
	req, _ := g.NewRequest("PUT", path, issue)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}

// MoveAnIssue Moves an issue to a different project.
//
// Gitlab API docs: https://docs.gitlab.com/ee/api/issues.html#move-an-issue
func (g *Client) MoveAnIssue(projectID, issueID, toProjectId string) (*Issue, *http.Response, error) {

	body := MoveIssueRequest{
		ToProjectID: toProjectId,
	}

	path := []string{"projects", url.QueryEscape(projectID), "issues", issueID, "move"}
	req, _ := g.NewRequest("POST", path, body)

	var ret *Issue
	if res, err := g.Do(req, &ret); err != nil {
		return nil, res, err
	}

	return ret, nil, nil
}
