package teamcity

import (
	"bytes"
	"fmt"
	"strconv"
)

func (c *Client) SetBuildConfigurationPaused(buildConfID string, state bool) error {
	path := fmt.Sprintf("/app/rest/buildTypes/id:%s/paused", buildConfID)

	body := bytes.NewBuffer([]byte(strconv.FormatBool(state)))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
