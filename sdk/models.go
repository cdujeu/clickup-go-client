package sdk

import (
	"net/url"

	"github.com/ahmetalpbalkan/qs"
)

type TaskTag struct {
	Name  string `json:"name"`
	TagFg string `json:"tag_fg"`
	TagBg string `json:"tag_bg"`
}

type PutTaskRequest struct {
	Name      string `json:"name"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	Assignees []int  `json:"assignees,omitempty"`
	Priority  int    `json:"priority,omitempty"`
	DueDate   string `json:"due_date,omitempty"`
}

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	TextContent string `json:"text_content"`
	Status      struct {
		Status     string `json:"status"`
		Type       string `json:"type"`
		Orderindex int    `json:"orderindex"`
		Color      string `json:"color"`
	} `json:"status"`
	Orderindex  string `json:"orderindex"`
	DateCreated string `json:"date_created"`
	DateUpdated string `json:"date_updated"`
	DateClosed  string `json:"date_closed"`
	Creator     struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"creator"`
	Assignees []struct {
		ID             int    `json:"id"`
		Username       string `json:"username"`
		Color          string `json:"color"`
		ProfilePicture string `json:"profilePicture"`
	} `json:"assignees"`
	Tags         []*TaskTag  `json:"tags"`
	Parent       string      `json:"parent"`
	Priority     interface{} `json:"priority"`
	DueDate      interface{} `json:"due_date"`
	StartDate    interface{} `json:"start_date"`
	Points       float64     `json:"points"`
	TimeEstimate float64     `json:"time_estimate"`
	Space        struct {
		ID string `json:"id"`
	} `json:"space"`
	Project struct {
		ID string `json:"id"`
	} `json:"project"`
	List struct {
		ID string `json:"id"`
	} `json:"list"`
	URL string `json:"url"`
}

type ListTasksRequest struct {
	// Filters
	SpaceIds   []string `qs:"space_ids[],omitempty"`
	ProjectIds []string `qs:"project_ids[],omitempty"`
	ListIds    []string `qs:"list_ids[],omitempty"`
	Statuses   []string `qs:"statuses[],omitempty"`
	Assignees  []string `qs:"assignees[],omitempty"`

	// Pagination and ordering
	Page          int    `qs:"page,omitempty"`
	OrderBy       string `qs:"order_by,omitempty"`
	Reverse       bool   `qs:"reverse,omitempty"`
	SubTasks      bool   `qs:"subtasks,omitempty"`
	IncludeClosed bool   `qs:"include_closed,omitempty"`

	// Date filters
	DueDateGt     int `qs:"due_date_gt,omitempty"`
	DueDateLt     int `qs:"due_date_lt,omitempty"`
	DateCreatedGt int `qs:"date_created_gt,omitempty"`
	DateCreatedLt int `qs:"date_created_lt,omitempty"`
	DateUpdatedGt int `qs:"date_updated_gt,omitempty"`
	DateUpdatedLt int `qs:"date_updated_lt,omitempty"`

	FilterByTag string
}

func (l *ListTasksRequest) Encode() url.Values {
	return qs.Encode(*l)
}

type ListTasksResponse struct {
	Tasks []*Task `json:"tasks"`
}

type ApiError struct {
	Error string `json:"err"`
	Code  string `json:"ECODE"`
}
