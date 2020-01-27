package teamcity

import (
	"errors"
	"fmt"

	"github.com/icelander/teamcity-sdk-go/types"
)

func (c *Client) CreateProject(project *types.Project) error {
	path := fmt.Sprintf("/app/rest/%s/projects", c.version)
	var projectReturn *types.Project

	err := c.doRetryRequest("POST", path, project, &projectReturn)
	if err != nil {
		return err
	}

	if projectReturn == nil {
		return errors.New("project not created")
	}
	*project = *projectReturn

	return nil
}
