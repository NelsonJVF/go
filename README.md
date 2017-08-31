# GoJira [![GoDoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](http://godoc.org/github.com/nelsonjvf/gojira) [![Build Status](http://img.shields.io/travis/fatih/structs.svg?style=flat-square)]() [![Coverage Status](http://img.shields.io/coveralls/fatih/structs.svg?style=flat-square)]()

GoJira is a generic library to interact with Atlassian Jira REST API.

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

After that you can simply use the following methods to interact with JIRA:

```go
// Get Jira issue information
jira.RequestIssue("JIRA-4968")

// Search a string in Jira
jira.RequestSearch("Error on workspace")
```

### GoJira methods

We can get an issue information

```jira.RequestIssue(issueId)```

We can also search in Jira:

```jira.RequestSearch(query)```

## Credits

 * [Nelson Ferreira](https://github.com/nelsonjvf)

## License

The MIT License (MIT) - see LICENSE.md for more details
