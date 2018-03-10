package main

import (
	"fmt"

	"github.com/nelsonjvf/gojira"
)

func main() {

	var jiraURL = "http://jira-server.com/"
	var jiraUsername = "my-jira-user"
	var jiraPassword = "my-jira-password"

	fmt.Println("Staring Testing..")
	fmt.Println("Setting Configuration")

	fmt.Println(gojira.Config)

	gojira.RequestIssue(yourJiraURL, jiraUsername, jiraPassword, "ISSUE-12345")
	gojira.RequestSearch(yourJiraURL, jiraUsername, jiraPassword, "Bug with string")

}
