package snykclient

import (
	"encoding/json"
	"fmt"
)

const (
	integrationsEndpoint = "orgId/integrations"
)


func (s *SnykClient) ListIntegrations() (map[string]string, error) {
	url := fmt.Sprintf("https://snyk.io/api/v1/org/%s/integrations", s.OrgID)
	resp, err := s.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	integrations := make(map[string]string, 0)

	err = json.Unmarshal(resp, &integrations)
	if err != nil {
		return nil, err
	}

	return integrations, nil
}
