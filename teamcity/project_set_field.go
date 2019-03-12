package teamcity

import (
	"bytes"
	"fmt"
	"strings"
)

func (c *Client) SetProjectField(projectID, field string, content string) error {
	path := fmt.Sprintf("/httpAuth/app/rest/%s/projects/id:%s/%s", c.version, projectID, strings.ToLower(field))

	body := bytes.NewBuffer([]byte(content))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
