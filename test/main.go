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

	fmt.Println("Staring Testing..")

	fmt.Println("Setting Configuration")

	fmt.Println(gojira.Config)

	gojira.RequestIssue("Test Jira Env", "ISSUE-12345")

	gojira.RequestSearch("Project", "Bug with string")

}