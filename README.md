# GoJira  [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/fatih/structs) [![Build Status](http://img.shields.io/travis/fatih/structs.svg?style=flat-square)](https://travis-ci.org/fatih/structs) [![Coverage Status](http://img.shields.io/coveralls/fatih/structs.svg?style=flat-square)](https://coveralls.io/r/fatih/structs)
GoLang projects

GoJira it's a generic library to interact with Atlassian Jira REST API.

## Install

```bash
go get github.com/nelsonjvf/gojira
```

## Usage and Examples

First of all you need to add your Jira information:

```go
config = []configuration{
  configuration {
    Lable: "Lable",
    User: "User",
    Pass: "Pass",
    Url: "URL",
  },
  configuration {
    Lable: "Lable",
    User: "User",
    Pass: "Pass",
    Url: "URL",
  },
}
```

...

```go
// Get Jira issue information
jira.RequestIssue("JIRA-4968")

// Search a string in Jira
jira.RequestSearch("Error on workspace")
```

### GoJira methods

We can get a issue information

```jira.jiraIssue(issueId)```

We can also search in Jira:

```RequestSearch(query)```

## Credits

 * [Nelson Ferreira](https://github.com/nelsonjvf)

## License

The MIT License (MIT) - see LICENSE.md for more details
