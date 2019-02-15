package sdk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Client struct {
	Token  string
	TeamID string
}

// ListTasks calls the ClickUp API to list all tasks, with filtering criteria passed by ListTaskRequest object
func (c *Client) ListTasks(req *ListTasksRequest) (*ListTasksResponse, error) {

	url := "https://api.clickup.com/api/v1/team/" + c.TeamID + "/task?" + req.Encode().Encode()

	request, _ := http.NewRequest("GET", url, nil)
	request.Header.Add("Authorization", c.Token)
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	b, e := ioutil.ReadAll(response.Body)
	if e != nil {
		return nil, e
	}
	list := ListTasksResponse{}
	err = json.Unmarshal(b, &list)
	if err != nil {
		apiError := ApiError{}
		if err = json.Unmarshal(b, &apiError); err == nil && apiError.Error != "" {
			return nil, fmt.Errorf(apiError.Error)
		} else {
			log.Println("Cannot decode body" + string(b))
			return nil, err
		}
	}

	// Additional Filtering on Response
	if req.FilterByTag != "" {
		var newList []*Task
		for _, t := range list.Tasks {
			var hasTag bool
			for _, tag := range t.Tags {
				if tag.Name == req.FilterByTag {
					hasTag = true
					break
				}
			}
			if hasTag {
				newList = append(newList, t)
			}
		}
		list.Tasks = newList
	}

	return &list, nil

}

// CreateTask creates a new tasks inside a given List
func (c *Client) CreateTask(listId string, request *PutTaskRequest) (string, error) {

	client := http.DefaultClient
	body, _ := json.Marshal(request)

	req, _ := http.NewRequest("POST", "https://api.clickup.com/api/v1/list/"+listId+"/task", bytes.NewBuffer(body))
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	type respStruct struct {
		ID string `json:"id"`
	}
	var response respStruct
	if e := json.Unmarshal(resp_body, &response); e == nil {
		return response.ID, nil
	}
	var apiError ApiError
	if e := json.Unmarshal(resp_body, &apiError); e == nil {
		return "", fmt.Errorf(apiError.Error)
	}
	return "", fmt.Errorf("cannot parse response body: %s", string(resp_body))

}

// UpdateTask updates a task by its ID
func (c *Client) UpdateTask(taskId string, request *PutTaskRequest) error {

	client := http.DefaultClient
	body, _ := json.Marshal(request)

	req, _ := http.NewRequest("PUT", "https://api.clickup.com/api/v1/task/"+taskId, bytes.NewBuffer(body))
	req.Header.Add("Authorization", c.Token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)
	var response map[string]interface{}
	if e := json.Unmarshal(resp_body, &response); e != nil {
		return e
	}
	if errorMessage, ok := response["err"]; ok {
		return fmt.Errorf(errorMessage.(string))
	}
	return nil

}
