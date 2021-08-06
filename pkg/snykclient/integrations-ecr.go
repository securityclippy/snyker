package snykclient

import (
	"bytes"
	"github.com/securityclippy/snyker/pkg/snyk"

	"encoding/json"
	"errors"
	"fmt"
	//"github.com/securityclippy/snyker/pkg/snyk"
	"io/ioutil"
)

func (s *SnykClient) ImportECRImage(imageName string) error {

	// set ecr integration id
	if s.ECRIntegrationID == "" {
		resp, err := s.ListIntegrations()
		if err != nil {
			return err
		}
		ecrID, ok := resp["ecr"]
		if !ok {
			return errors.New("no ecr integration found")
		}

		s.ECRIntegrationID = ecrID
	}



	target := snyk.Target{
		Name:imageName,
	}

	ecrPost := snyk.ECRPost{Target:target}

	js, err := json.Marshal(ecrPost)
	if err != nil {
		return err
	}


	//body := []byte("{\n\"target\": {\n\"name\":\"528451384384.dkr.ecr.us-west-2.amazonaws.com/ajs-renderer:9d48ebf_1038\"\n}\n}")

	//https://snyk.io/api/v1/org/orgId/integrations/integrationId/import

	url := fmt.Sprintf("https://snyk.io/api/v1/org/%s/integrations/%s/import", s.OrgID, s.ECRIntegrationID)
	//url := fmt.Sprintf("http://127.0.0.1:8000/api/v1/org/%s/integrations/%s/import", s.OrgID, s.ECRIntegrationID)
	//url := fmt.Sprintf("https://snyk.io/api/v1/org/%s/integrations/%s/import", "segment-docker", s.ECRIntegrationID)

	fmt.Printf("url: %s\n\n", url)

	//fmt.Printf("PostPody: \n%s\n\n", string(js))

	req, err := s.newRequest("POST", url, bytes.NewReader(js))
	if err != nil {
		return err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println(string(b)+"\n")

	if err != nil {
		return err
	}

	fmt.Println(resp.Status)

	return nil
}