# GoJira [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/nelsonjvf/gojira/pkg)

GoJira is a basic and generic package to interact with Atlassian Jira REST API.
The idea of this package is to make your life easy, instead of learning the JIRA REST API, you just need to set your configuration and get the information.
All the methos available return an Go object with the smae struter of the response from JIRA REST API.
Fell free to add make comments and review the code.

## Install

```bash
go get github.com/nelsonjvf/gojira/pkg
```

## Usage and Examples

Simply use the methods available to interact with JIRA:

```go
// Get Jira issue information
gojira.RequestIssue("http://jira-server.com/", "user", "pass", "JIRA-4968")

// Search a string in Jira
gojira.RequestSearch("http://jira-server.com/", "user", "pass", "Error on workspace")
```

There is the full code to test it easly:

```go
package main

import (
	"io/ioutil"
	"log"

	"github.com/nelsonjvf/gojira"
)

func main() {

	var jiraURL = "http://jira-server.com/"
	var jiraUsername = "my-jira-user"
	var jiraPassword = "my-jira-password"

	log.Println("Starting test..")
	log.Println("Calling RequestIssue method:")

	issueToSearch := "ISSUE-12345"
	requestIssueResponse := gojira.RequestIssue(yourJiraURL, jiraUsername, jiraPassword, issueToSearch)
	log.Println(requestIssueResponse)

	log.Println("Calling RequestSearch method:")

	stringToSearch := "Bug with comma"
	requestSearchResponse := gojira.RequestSearch(yourJiraURL, jiraUsername, jiraPassword, stringToSearch)
	log.Println(requestSearchResponse)
}
```

### GoJira methods

We can get an issue information:

```RequestIssue(URL, user, pass, issueId)```

We can also search in Jira:

```RequestSearch(URL, user, pass, query)```

## Credits

 * [Nelson Ferreira](https://github.com/nelsonjvf)

## License

The MIT License (MIT) - see LICENSE.md for more details
