package teamcity

import (
	"fmt"
)

func (c *Client) DeleteProject(projectID string) error {
	path := fmt.Sprintf("/app/rest/projects/id:%s", projectID)
	return c.doRetryRequest("DELETE", path, nil, nil)
}
