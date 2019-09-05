package main

import (
    "fmt"
    "strconv"
    "net/http"
    "io/ioutil"
    "encoding/json"
    jira "github.com/andygrunwald/go-jira"
)

const jira_user = "xxxx"
const jira_pass = "xxxxxx"

type Violation struct {
  Created            string    `json:"created"`
  TopSeverity        string    `json:"top_severity"`
  WatchName          string    `json:"watch_name"`
  PolicyName         string    `json:"policy_name"`
  Issues             Issues    `json:"issues"`
}

type Issue struct {
  Severity           string             `json:"severity"`
  Type               string             `json:"type"` // Issue type license/security
  Summary            string             `json:"summary"`
  Description        string             `json:"description"`
  ImpactedArtifacts  ImpactedArtifacts  `json:"impacted_artifacts"`
  CVE                string             `json:"cve"`
}

type Issues []Issue

type ImpactedArtifact struct {
  Name             string         `json:"name"` // Artifact name
  DisplayName      string         `json:"display_name"`
  Path             string         `json:"path"`  // Artifact path in Artifactory
  PackageType      string         `json:"pkg_type"`
  SHA256           string         `json:"sha256"` // Artifact SHA 256 checksum
  SHA1             string         `json:"sha1"`
  Depth            int            `json:"depth"`  // Artifact depth in its hierarchy
  ParentSHA        string         `json:"parent_sha"`
  InfectedFiles    InfectedFiles  `json:"infected_files"`
}

type ImpactedArtifacts []ImpactedArtifact

type InfectedFile struct {
Name           string    `json:"name"`
Path           string    `json:"path"`  // artifact path in Artifactory
SHA256         string    `json:"sha256"`// artifact SHA 256 checksum
Depth          int       `json:"depth"` // Artifact depth in its hierarchy
ParentSHA      string    `json:"parent_sha"` // Parent artifact SHA1
DisplayName    string    `json:"display_name"`
PackageType    string    `json:"pkg_type"`
}

type InfectedFiles []InfectedFile

func getIssuesCount(violation Violation) int {
    
    return len(violation.Issues)

}


func handler(writer http.ResponseWriter, request *http.Request) {
    var violation Violation

    fmt.Fprintf(writer, "Get the Alert!")

    body, err := ioutil.ReadAll(request.Body)
    if err != nil {
        fmt.Printf("read body err, %v\n", err)
        return
    }

    if err := json.Unmarshal(body, &violation); err != nil {
        fmt.Printf("unmarshal json err, %v\n", err)
        return
    }

    fmt.Printf("Get Violation:\n")
    fmt.Printf("%+v\n", violation)

    tp := jira.BasicAuthTransport{
        Username: jira_user,
        Password: jira_pass,
    }

    jiraClient, err := jira.NewClient(tp.Client(), "http://jira.xxxxxx.com:8081")
    if err != nil {
        fmt.Printf("Can not create Jira Client, %v\n", err)
        return
    }
    fmt.Printf("Created Jira Client!\n")

    i := jira.Issue{
        Fields: &jira.IssueFields{
            Assignee: &jira.User{
                Name: jira_user,
            },
            Reporter: &jira.User{
                Name: jira_user,
            },
            Description: "The watch \"" + violation.WatchName + "\" with policy \"" + violation.PolicyName + "\" created a violation with " + strconv.Itoa(getIssuesCount(violation)) + " number of issues",
            Type: jira.IssueType{
                Name: "Bug",
            },
            Project: jira.Project{
                Key: "XRAYW",
            },
            Summary: "Violation was found with \"" + violation.TopSeverity + "\" severity",
        },
    }

    issue, _, err := jiraClient.Issue.Create(&i)
    if err != nil {
        fmt.Printf("Can not create Jira Issue, %v\n", err)
        return
    }
    fmt.Printf("Created Jira Issue: %s\n", issue.Key)

}

func main() {
	http.HandleFunc("/xray/", handler)
	http.ListenAndServe(":9999", nil)
}
