// Copyright 2017 NelsonJVF. All rights reserved.
package gojira

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

// Generic HTTP caller
// URL - Jira address
// Path - Rest api path - Useally "rest/api/2"
// Username - Jira username
// Password - Jira user password
// Parameters - Parameters to send to Jira
func HTTPRequest(URL string, path string, username string, password string, parameters url.Values) (http.Response, error) {

	var resp http.Response
	var jiraURL string
	var restAPIPath = "rest/api/2"

	if len(URL) == 0 {
		return resp, errors.New("Jira URL missing")
	}

	jiraURL = fmt.Sprintf("%s%s%s", URL, restAPIPath, path)

	timeoutVal := time.Duration(time.Duration(30) * time.Second)
	client := &http.Client{
		Timeout: timeoutVal,
	}
	r, _ := http.NewRequest("GET", jiraURL, bytes.NewBufferString(parameters.Encode()))

	r.SetBasicAuth(username, password)
	r.Header.Set("Content-Type", "application/json")

	res, errDo := client.Do(r)
	if errDo != nil {
		return resp, errDo
	}
	defer resp.Body.Close()

	resp = *res

	return resp, nil
}

// RequestIssue - Request specific Jira item, we should specify the project from that item
func RequestIssue(URL string, username string, password string, item string) (JiraIssueResponse, error) {
	var URLIssuePath string
	var data JiraIssueResponse

	URLIssuePath = fmt.Sprintf("/issue/%s", item)

	resp, e := HTTPRequest(URL, URLIssuePath, username, password, nil)
	if e != nil {
		return data, e
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errMsg := fmt.Sprintf("ioutil.ReadAll err   #%v ", err)
		return data, errors.New(errMsg)
	}

	e = json.Unmarshal(body, &data)
	if e != nil {
		errMsg := fmt.Sprintf("ioutil.ReadAll err   #%v ", err)
		return data, errors.New(errMsg)
	}

	return data, nil
}

// RequestSearch - Search in Jira, we should specify the project from that item
func RequestSearch(URL string, username string, password string, query string) (JiraSearchResponse, error) {
	var URLSearchPath string
	var data JiraSearchResponse

	URLSearchPath = fmt.Sprintf("/search/%s", query)

	resp, e := HTTPRequest(URL, URLSearchPath, username, password, nil)
	if e != nil {
		return data, e
	}

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		errMsg := fmt.Sprintf("GoJira - ioutil.ReadAll err #%v ", e)
		return data, errors.New(errMsg)
	}

	e = json.Unmarshal(body, &data)
	if e != nil {
		errMsg := fmt.Sprintf("GoJira - json.Unmarshal err #%v ", e)
		return data, errors.New(errMsg)
	}

	return data, nil
}
