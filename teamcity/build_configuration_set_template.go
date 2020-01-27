package teamcity

import (
	"bytes"
	"fmt"
)

func (c *Client) SetBuildConfigurationTemplate(buildConfID, templateID string) error {
	path := fmt.Sprintf("/app/rest/buildTypes/id:%s/template", buildConfID)

	if templateID != "" {
		body := bytes.NewBuffer([]byte("id:" + templateID))
		_, err := c.doNotJSONRequest("PUT", path, "application/json", "text/plain", body)
		if err != nil {
			return err
		}
	} else {
		return c.doRetryRequest("DELETE", path, nil, nil)
	}
	return nil
}
