// Copyright 2017 NelsonJVF. All rights reserved.
package gojira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
	Jira(s) configuration(s)
*/
var Config []Configuration

/*
	Generic HTTP caller
*/
func HTTPRequest(project string, urlPath string) []byte {
	var userJira string
	var passJira string
	var urlJira string
	var restAPIPath = "rest/api/2"

	for _, c := range Config {
		if c.Lable == project {
			userJira = c.User
			passJira = c.Pass
			urlJira = c.Url
		}
	}

	if len(urlJira) == 0 {
		log.Printf(" ---------- Jira configuration is missing  ---------- ")
		return nil
	}

	urlJira = fmt.Sprintf("%s%s%s", urlJira, restAPIPath, urlPath)

	req, err := http.NewRequest("GET", urlJira, nil)
	if err != nil {
		log.Printf("http.NewRequest err   #%v ", err)
	}

	req.SetBasicAuth(userJira, passJira)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("http.DefaultClient.Do err   #%v ", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil.ReadAll err   #%v ", err)
	}

	return body
}

/*
	Request specific Jira item, we should specify the project from that item
*/
func RequestIssue(project string, item string) JiraIssueResponse {
	var urlIssuePath string
	var data JiraIssueResponse

	urlIssuePath = fmt.Sprintf("/issue/%s", item)

	response := HTTPRequest(project, urlIssuePath)

	err := json.Unmarshal(response, &data)
	if err != nil {
		log.Printf("json.Unmarshal err   #%v ", err)
	}

	return data
}

/*
	Search in Jira, we should specify the project from that item
*/
func RequestSearch(project string, query string) JiraSearchResponse {
	var urlSearchPath string
	var data JiraSearchResponse

	urlSearchPath = fmt.Sprintf("/search/%s", query)

	response := HTTPRequest(project, urlSearchPath)

	err := json.Unmarshal(response, &data)
	if err != nil {
		log.Printf("json.Unmarshal err   #%v ", err)
	}

	return data
}
