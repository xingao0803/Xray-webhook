package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "encoding/json"
)

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

    fmt.Printf("Get Violation:")
    fmt.Printf("%+v", violation)

}

func main() {
	http.HandleFunc("/xray/", handler)
	http.ListenAndServe(":9999", nil)
}
