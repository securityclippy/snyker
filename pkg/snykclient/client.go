package snykclient

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	baseOrgURL = "https://snyk.io/api/v1/org"
)

type SnykClient struct {
	*http.Client
	AuthToken string
	Org       string
	OrgID     string
	ECRIntegrationID string
}

func NewSnykClient(authToken string) (*SnykClient, error) {
	c := &http.Client{
		Timeout: time.Second * 30,
	}

	sc := &SnykClient{
		Client:    c,
		AuthToken: authToken,
	}

	err := sc.mustSetOrgID()
	err = sc.mustSetOrgname()
	if err != nil {
		return nil, err
	}


	return sc, nil
}

func (s *SnykClient) newRequest(method, url string, payload io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	header := http.Header{}
	header.Set("Authorization", fmt.Sprintf("token: %s", s.AuthToken))
	header.Set("Content-Type", "application/json")
	header.Set("User-Agent", "Snyker")
	req.Header = header
	//req.Header.Set("Authorization", fmt.Sprintf("JWT %s", c.AuthToken))
	return req, nil
}

func (s *SnykClient) doRequest(method, url string, payload io.Reader) ([]byte, error) {
	req, err := s.newRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (s *SnykClient) getOrgID() (string, error) {
	out, err := s.ListOrgs()
	if err != nil {
		return "", nil
	}
	if s.Org != "" {
		for _, o := range out.Orgs {
			if o.Name == s.Org {
				return o.ID, nil
			}
		}
	}
	if len(out.Orgs) > 0 {

		return out.Orgs[0].ID, nil
	}
	return "", errors.New("no orgID found")
}

func (s *SnykClient) getOrgName() (string, error) {
	out, err := s.ListOrgs()
	if err != nil {
		return "", nil
	}
	if s.Org != "" {
		for _, o := range out.Orgs {
			if o.Name == s.Org {
				return o.ID, nil
			}
		}
	}

	if len(out.Orgs) > 0 {

		return out.Orgs[0].ID, nil
	}
	return "", errors.New("no orgID found")

}

func (s *SnykClient) setOrgID(id string) {
	s.OrgID = id
}

func (s *SnykClient) setOrgName(name string) {
	s.Org = name
}

func (s *SnykClient) mustSetOrgID() error {
	id, err := s.getOrgID()
	if err != nil {
		return err
	}
	s.setOrgID(id)
	return nil
}

func (s *SnykClient) mustSetOrgname() error {
	name, err := s.getOrgName()
	if err != nil {
		return err
	}

	s.setOrgName(name)
	return nil
}
