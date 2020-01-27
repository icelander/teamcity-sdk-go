package teamcity

import (
	"errors"

	"github.com/icelander/teamcity-sdk-go/types"
)

func (c *Client) CreateBuildConfiguration(buildConfig *types.BuildConfiguration) error {
	path := "/app/rest/buildTypes"
	var buildConfigReturn *types.BuildConfiguration

	err := c.doRetryRequest("POST", path, buildConfig, &buildConfigReturn)
	if err != nil {
		return err
	}

	if buildConfigReturn == nil {
		return errors.New("build configuration not created")
	}
	*buildConfig = *buildConfigReturn

	return nil
}
