package snyk

import "time"

type ProjectIssuesResp struct {
	Filters Filters `json:"filters"`
}
type Filters struct {
	Severities []string `json:"severities"`
	Types      []string `json:"types"`
	Ignored    bool     `json:"ignored"`
	Patched    bool     `json:"patched"`
}

type Org struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

type ListProjectsResp struct {
	ORG      Org        `json:"org"`
	Projects []*Project `json:"projects"`
}

type Project struct {
	Name                  string                `json:"name"`
	ID                    string                `json:"id"`
	Created               time.Time             `json:"created"`
	Origin                string                `json:"origin"`
	Type                  string                `json:"type"`
	ReadOnly              bool                `json:"readOnly"`
	TestFrequency         string                `json:"testFrequency"`
	TotalDependencies     int                   `json:"totalDependencies"`
	IssueCountsBySeverity IssueCountsBySeverity `json:"issueCountsBySeverity"`
	ImageID               string                `json:"imageId,omitempty"`
	ImageTag              string                `json:"imageTag,omitempty"`
}

type IssueCountsBySeverity struct {
	Low    int `json:"low"`
	High   int `json:"high"`
	Medium int `json:"medium"`
}


type ProjectSettings struct {
	PullRequestTestEnabled bool `json:"pullRequestTestEnabled"`
	PullRequestFailOnAnyVulns bool `json:"pullRequestFailOnAnyVulns"`
	PullRequestFailOnlyHighSeverity bool `json:"pullRequestFailOnlyForHighSeverity"`
}