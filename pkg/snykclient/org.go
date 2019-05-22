package snykclient

import (
	"encoding/json"
	"io/ioutil"
)

type OrgResp struct {
	Orgs []struct {
		Name  string `json:"name"`
		ID    string `json:"id"`
		Group struct {
			Name string `json:"name"`
			ID   string `json:"id"`
		} `json:"group"`
	} `json:"orgs"`
}

func (s *SnykClient) ListOrgs() (*OrgResp, error) {
	url := "https://snyk.io/api/v1/orgs"
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

	orgResp := &OrgResp{}
	err = json.Unmarshal(body, orgResp)
	if err != nil {
		return nil, err
	}

	return orgResp, nil
}
