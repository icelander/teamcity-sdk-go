package teamcity

import (
	"fmt"

	"github.com/icelander/teamcity-sdk-go/types"
)

// GetProject gets a project based on the Project ID
func (c *Client) GetProject(projectID string) (*types.Project, error) {
	path := fmt.Sprintf("/app/rest/projects/id:%s", projectID)
	var project *types.Project

	err := c.doRetryRequest("GET", path, nil, &project)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// GetProjects returns all projects
func (c *Client) GetProjects() ([]types.Project, error) {
	path := "/app/rest/projects"
	var projects struct {
		Count int64
		HREF string
		Project []types.Project
	}

	err := c.doRequest("GET", path, nil, &projects)

	if err != nil {
		return nil, err
	}

	return projects.Project, nil
}

// GetShortProjects returns all projects in short form
func (c *Client) GetShortProjects() ([]types.ProjectShort, error) {
	path := "/app/rest/projects"
	var projects struct {
		Count int64
		HREF string
		Project []types.ProjectShort
	}

	err := c.doRequest("GET", path, nil, &projects)

	if err != nil {
		return nil, err
	}

	return projects.Project, nil
}