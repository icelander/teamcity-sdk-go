package teamcity

import (
	"bytes"
	"fmt"
	"strings"
)

func (c *Client) SetProjectField(projectID, field string, content string) error {
	path := fmt.Sprintf("/app/rest/projects/id:%s/%s", projectID, strings.ToLower(field))

	body := bytes.NewBuffer([]byte(content))
	_, err := c.doNotJSONRequest("PUT", path, "text/plain", "text/plain", body)
	if err != nil {
		return err
	}
	return nil
}
