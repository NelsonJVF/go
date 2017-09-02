// Copyright 2017 NelsonJVF. All rights reserved.

package gojira

import (
	"fmt"
  "net/http"
  "log"
	"io/ioutil"
	"encoding/json"
	"gopkg.in/yaml.v2"
)

/*
	Struct for Jira access information
 */
type configuration struct {
  Lable string `yaml:"lable"` // Some projects have more than one Jira, so just lable as you wish
  User string  `yaml:"user"` // Username for Jira
  Pass string  `yaml:"pass"` // Password from Jira Username
  Url string   `yaml:"url"` // URL to Jira hostname + port
}

/*
	Jira Issue Response Struct
 */
type JiraIssueResponse struct {
	Expand string `json:"expand"`
	ID     string `json:"id"`
	Self   string `json:"self"`
	Key    string `json:"key"`
	Fields struct {
		Issuetype struct {
			Self        string `json:"self"`
			ID          string `json:"id"`
			Description string `json:"description"`
			IconURL     string `json:"iconUrl"`
			Name        string `json:"name"`
			Subtask     bool   `json:"subtask"`
			AvatarID    int    `json:"avatarId"`
		} `json:"issuetype"`
		Components           []interface{} `json:"components"`
		Timespent            interface{}   `json:"timespent"`
		Timeoriginalestimate int           `json:"timeoriginalestimate"`
		Description          string        `json:"description"`
		Project              struct {
			Self       string `json:"self"`
			ID         string `json:"id"`
			Key        string `json:"key"`
			Name       string `json:"name"`
			AvatarUrls struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
		} `json:"project"`
		FixVersions        []interface{} `json:"fixVersions"`
		Aggregatetimespent interface{}   `json:"aggregatetimespent"`
		Resolution         interface{}   `json:"resolution"`
		Timetracking       struct {
			OriginalEstimate         string `json:"originalEstimate"`
			RemainingEstimate        string `json:"remainingEstimate"`
			OriginalEstimateSeconds  int    `json:"originalEstimateSeconds"`
			RemainingEstimateSeconds int    `json:"remainingEstimateSeconds"`
		} `json:"timetracking"`
		Customfield10600 interface{} `json:"customfield_10600"`
		Customfield10700 struct {
			Self  string `json:"self"`
			Value string `json:"value"`
			ID    string `json:"id"`
		} `json:"customfield_10700"`
		Attachment            []interface{} `json:"attachment"`
		Aggregatetimeestimate int           `json:"aggregatetimeestimate"`
		Resolutiondate        interface{}   `json:"resolutiondate"`
		Workratio             int           `json:"workratio"`
		Summary               string        `json:"summary"`
		LastViewed            string        `json:"lastViewed"`
		Watches               struct {
			Self       string `json:"self"`
			WatchCount int    `json:"watchCount"`
			IsWatching bool   `json:"isWatching"`
		} `json:"watches"`
		Creator struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"creator"`
		Subtasks []interface{} `json:"subtasks"`
		Created  string        `json:"created"`
		Reporter struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"reporter"`
		Customfield10000  string `json:"customfield_10000"`
		Aggregateprogress struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
			Percent  int `json:"percent"`
		} `json:"aggregateprogress"`
		Priority struct {
			Self    string `json:"self"`
			IconURL string `json:"iconUrl"`
			Name    string `json:"name"`
			ID      string `json:"id"`
		} `json:"priority"`
		Customfield10001              interface{}   `json:"customfield_10001"`
		Customfield10002              interface{}   `json:"customfield_10002"`
		Customfield10200              interface{}   `json:"customfield_10200"`
		Customfield10101              interface{}   `json:"customfield_10101"`
		Customfield10201              interface{}   `json:"customfield_10201"`
		Customfield10300              interface{}   `json:"customfield_10300"`
		Labels                        []string      `json:"labels"`
		Customfield10103              interface{}   `json:"customfield_10103"`
		Environment                   interface{}   `json:"environment"`
		Timeestimate                  int           `json:"timeestimate"`
		Aggregatetimeoriginalestimate int           `json:"aggregatetimeoriginalestimate"`
		Versions                      []interface{} `json:"versions"`
		Duedate                       interface{}   `json:"duedate"`
		Progress                      struct {
			Progress int `json:"progress"`
			Total    int `json:"total"`
			Percent  int `json:"percent"`
		} `json:"progress"`
		Comment struct {
			StartAt    int           `json:"startAt"`
			MaxResults int           `json:"maxResults"`
			Total      int           `json:"total"`
			Comments   []interface{} `json:"comments"`
		} `json:"comment"`
		Issuelinks []interface{} `json:"issuelinks"`
		Votes      struct {
			Self     string `json:"self"`
			Votes    int    `json:"votes"`
			HasVoted bool   `json:"hasVoted"`
		} `json:"votes"`
		Worklog struct {
			StartAt    int           `json:"startAt"`
			MaxResults int           `json:"maxResults"`
			Total      int           `json:"total"`
			Worklogs   []interface{} `json:"worklogs"`
		} `json:"worklog"`
		Assignee struct {
			Self         string `json:"self"`
			Name         string `json:"name"`
			Key          string `json:"key"`
			EmailAddress string `json:"emailAddress"`
			AvatarUrls   struct {
				Four8X48  string `json:"48x48"`
				Two4X24   string `json:"24x24"`
				One6X16   string `json:"16x16"`
				Three2X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			DisplayName string `json:"displayName"`
			Active      bool   `json:"active"`
			TimeZone    string `json:"timeZone"`
		} `json:"assignee"`
		Updated string `json:"updated"`
		Status  struct {
			Self           string `json:"self"`
			Description    string `json:"description"`
			IconURL        string `json:"iconUrl"`
			Name           string `json:"name"`
			ID             string `json:"id"`
			StatusCategory struct {
				Self      string `json:"self"`
				ID        int    `json:"id"`
				Key       string `json:"key"`
				ColorName string `json:"colorName"`
				Name      string `json:"name"`
			} `json:"statusCategory"`
		} `json:"status"`
	} `json:"fields"`
}

/*
	Jira 'Search Response Struct
 */
type JiraSearchResponse struct {
	Expand     string `json:"expand"`
	StartAt    int    `json:"startAt"`
	MaxResults int    `json:"maxResults"`
	Total      int    `json:"total"`
	Issues     []struct {
		Expand string `json:"expand"`
		ID     string `json:"id"`
		Self   string `json:"self"`
		Key    string `json:"key"`
		Fields struct {
			Issuetype struct {
				Self        string `json:"self"`
				ID          string `json:"id"`
				Description string `json:"description"`
				IconURL     string `json:"iconUrl"`
				Name        string `json:"name"`
				Subtask     bool   `json:"subtask"`
				AvatarID    int    `json:"avatarId"`
			} `json:"issuetype"`
			Components           []interface{} `json:"components"`
			Timespent            interface{}   `json:"timespent"`
			Timeoriginalestimate interface{}   `json:"timeoriginalestimate"`
			Description          interface{}   `json:"description"`
			Project              struct {
				Self       string `json:"self"`
				ID         string `json:"id"`
				Key        string `json:"key"`
				Name       string `json:"name"`
				AvatarUrls struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
			} `json:"project"`
			FixVersions        []interface{} `json:"fixVersions"`
			Aggregatetimespent interface{}   `json:"aggregatetimespent"`
			Resolution         interface{}   `json:"resolution"`
			Customfield10600   interface{}   `json:"customfield_10600"`
			Customfield10700   struct {
				Self  string `json:"self"`
				Value string `json:"value"`
				ID    string `json:"id"`
			} `json:"customfield_10700"`
			Aggregatetimeestimate interface{} `json:"aggregatetimeestimate"`
			Resolutiondate        interface{} `json:"resolutiondate"`
			Workratio             int         `json:"workratio"`
			Summary               string      `json:"summary"`
			LastViewed            string      `json:"lastViewed"`
			Watches               struct {
				Self       string `json:"self"`
				WatchCount int    `json:"watchCount"`
				IsWatching bool   `json:"isWatching"`
			} `json:"watches"`
			Creator struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"creator"`
			Subtasks []interface{} `json:"subtasks"`
			Created  string        `json:"created"`
			Reporter struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"reporter"`
			Customfield10000  string `json:"customfield_10000"`
			Aggregateprogress struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"aggregateprogress"`
			Priority struct {
				Self    string `json:"self"`
				IconURL string `json:"iconUrl"`
				Name    string `json:"name"`
				ID      string `json:"id"`
			} `json:"priority"`
			Customfield10001              interface{}   `json:"customfield_10001"`
			Customfield10002              interface{}   `json:"customfield_10002"`
			Customfield10200              interface{}   `json:"customfield_10200"`
			Customfield10101              interface{}   `json:"customfield_10101"`
			Customfield10201              interface{}   `json:"customfield_10201"`
			Customfield10300              interface{}   `json:"customfield_10300"`
			Labels                        []string      `json:"labels"`
			Customfield10103              interface{}   `json:"customfield_10103"`
			Environment                   interface{}   `json:"environment"`
			Timeestimate                  interface{}   `json:"timeestimate"`
			Aggregatetimeoriginalestimate interface{}   `json:"aggregatetimeoriginalestimate"`
			Versions                      []interface{} `json:"versions"`
			Duedate                       interface{}   `json:"duedate"`
			Progress                      struct {
				Progress int `json:"progress"`
				Total    int `json:"total"`
			} `json:"progress"`
			Issuelinks []interface{} `json:"issuelinks"`
			Votes      struct {
				Self     string `json:"self"`
				Votes    int    `json:"votes"`
				HasVoted bool   `json:"hasVoted"`
			} `json:"votes"`
			Assignee struct {
				Self         string `json:"self"`
				Name         string `json:"name"`
				Key          string `json:"key"`
				EmailAddress string `json:"emailAddress"`
				AvatarUrls   struct {
					Four8X48  string `json:"48x48"`
					Two4X24   string `json:"24x24"`
					One6X16   string `json:"16x16"`
					Three2X32 string `json:"32x32"`
				} `json:"avatarUrls"`
				DisplayName string `json:"displayName"`
				Active      bool   `json:"active"`
				TimeZone    string `json:"timeZone"`
			} `json:"assignee"`
			Updated string `json:"updated"`
			Status  struct {
				Self           string `json:"self"`
				Description    string `json:"description"`
				IconURL        string `json:"iconUrl"`
				Name           string `json:"name"`
				ID             string `json:"id"`
				StatusCategory struct {
					Self      string `json:"self"`
					ID        int    `json:"id"`
					Key       string `json:"key"`
					ColorName string `json:"colorName"`
					Name      string `json:"name"`
				} `json:"statusCategory"`
			} `json:"status"`
		} `json:"fields"`
	} `json:"issues"`
}

var config []configuration

func setConf() {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func init() {
	setConf()
}

func checkError(e error) {
	if e != nil {
		log.Println(e)
	}
}

/*
	Generic HTTP caller
 */
func HTTPRequest(project string, urlPath string) []byte {

  var userJira string
  var passJira string
  var urlJira string

  for _, c := range config {
    if c.Lable == project {
      userJira = c.User
      passJira = c.Pass
      urlJira = c.Url
    }
  }

  urlJira = fmt.Sprintf("%s%s", urlJira, urlPath)

	req, err := http.NewRequest("GET", urlJira, nil)
	checkError(err)
	req.SetBasicAuth(userJira, passJira)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	checkError(err)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	return body
}

/*
	Request specific Jira item, we should specify the project from that item
 */
func RequestIssue(project string, item string) JiraIssueResponse {
  var url string
  var data JiraIssueResponse

  JiraIssuePath := "/rest/api/2/issue"

	url = fmt.Sprintf("%s/%s", JiraIssuePath, item)

  response := HTTPRequest(project, url)

	err := json.Unmarshal(response, &data)
	checkError(err)

  return data
}

/*
	Search in Jira, we should specify the project from that item
 */
func RequestSearch(project string, query string) {
  var url string
  var data JiraSearchResponse

  restApi := "/rest/api/2"
  searchPath := "/search"

	url = fmt.Sprintf("%s%s/%s", restApi, searchPath, query)

  response := HTTPRequest(project, url)

	err := json.Unmarshal(response, &data)
	checkError(err)

	fmt.Println(data)
}