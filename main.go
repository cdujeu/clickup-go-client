package main

import (
	"fmt"
	"log"

	"github.com/cdujeu/clickup-go-client/sdk"
)

func main() {

	// This is a sample code. You have to at least find the API Token and Team ID

	client := &sdk.Client{
		Token:  "PERSONAL_OR_OAUTH2_TOKEN",
		TeamID: "TEAM_ID",
	}

	r := &sdk.ListTasksRequest{
		SpaceIds:      []string{"SPACE_ID"},
		IncludeClosed: true,
	}

	list, err := client.ListTasks(r)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range list.Tasks {
		fmt.Println("Task " + t.Name)
		fmt.Println("----------------")
		fmt.Println(t.TextContent)
		fmt.Println("----------------")
	}
}
