package snykclient

import (
	"encoding/json"
	"fmt"
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
