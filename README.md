# GoJira [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/nelsonjvf/gojira) [![Build Status](http://img.shields.io/travis/fatih/structs.svg?style=flat-square)]() [![Coverage Status](http://img.shields.io/coveralls/fatih/structs.svg?style=flat-square)]()

GoJira is a generic library to interact with Atlassian Jira REST API.

## Install

```bash
go get github.com/nelsonjvf/gojira
```

## Usage and Examples

First of all we need to configure and set your Jira information. For that we can use our config.yaml file as example and the following init function:

```go
func init() {
	// Use yaml configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &gojira.Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}
```

After that you can simply use the methods available to interact with JIRA:

```go
// Get Jira issue information
gojira.RequestIssue("Dev Team", "JIRA-4968")

// Search a string in Jira
gojira.RequestSearch("Project", "Error on workspace")
```

There is the full code to test it easly:

```go
package main

import (
	"fmt"
	"github.com/nelsonjvf/gojira"
	"io/ioutil"
	"log"
	"gopkg.in/yaml.v2"
)

func init() {
	// Use yaml configuration file
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &gojira.Config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func main() {
	fmt.Println("Starting test..")

	fmt.Println("Calling RequestIssue method:")

	issueToSearch := "ISSUE-12345"
	requestIssueResponse := gojira.RequestIssue("Test Jira Env", issueToSearch)
    fmt.Println(requestIssueResponse)

	fmt.Println("Calling RequestSearch method:")

	stringToSearch := "Bug with comma"
	requestSearchResponse := gojira.RequestSearch("Project", stringToSearch)
	fmt.Println(requestSearchResponse)
}
```

### GoJira methods

We can get an issue information

```RequestIssue(project, issueId)```

We can also search in Jira:

```RequestSearch(project, query)```

## Credits

 * [Nelson Ferreira](https://github.com/nelsonjvf)

## License

The MIT License (MIT) - see LICENSE.md for more details
