package snykclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"io/ioutil"

	"github.com/securityclippy/snyker/pkg/snyk"
)

func (s *SnykClient) ListProjects() (*snyk.ListProjectsResp, error) {
	err := s.mustSetOrgID()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s/projects", baseOrgURL, s.OrgID)
	req, err := s.newRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	listProjResp := &snyk.ListProjectsResp{}
	err = json.Unmarshal(body, listProjResp)
	if err != nil {
		return nil, err
	}
	return listProjResp, nil
}


func (s *SnykClient) GetProject(projectContains string) (project *snyk.Project, err error) {
	projects, err := s.ListProjects()
	if err != nil {
		return nil, err
	}

	for _, p := range projects.Projects {
		if strings.Contains(p.Name, projectContains) {
			return p, nil
		}
	}
	return nil, fmt.Errorf("could not find project containing: %s", projectContains)
}

func (s *SnykClient) DeleteProjectByID(projectID string) error {
	err := s.mustSetOrgID()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s/project/%s", baseOrgURL, s.OrgID, projectID)
	req, err := s.newRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("[project: %s] %s", projectID, resp.Status)
	}
	return nil
}

func (s *SnykClient) ListProjectSettings(projectID string) (*snyk.ProjectSettings, error) {
	err := s.mustSetOrgID()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/%s/project/%s/settings", baseOrgURL, s.OrgID, projectID)
	fmt.Printf("listing: %s\n", url)
	resp, err := s.Get(url)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	psr := &snyk.ProjectSettings{}
	err = json.Unmarshal(body, psr)
	if err != nil {
		return nil, err
	}
	return psr, nil
}


func (s *SnykClient) UpdateProjectSettings(projectID string, settings *snyk.ProjectSettings) error {
	err := s.mustSetOrgID()
	if err != nil {
		return err
	}
	url := fmt.Sprintf("%s/%s/project/%s", baseOrgURL, s.OrgID, projectID)
	js, err := json.Marshal(settings)
	if err != nil {
		return err
	}
	req, err := s.newRequest("PUT", url, bytes.NewReader(js))
	if err != nil {
	}

	resp, err := s.Do(req)
	if err != nil {
		return err
	}


	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return nil
}