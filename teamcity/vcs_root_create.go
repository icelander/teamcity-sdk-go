package teamcity

import (
	"errors"

	"github.com/icelander/teamcity-sdk-go/types"
)

func (c *Client) CreateVcsRoot(vcs *types.VcsRoot) error {
	path := "/app/rest/vcs-roots"
	var vcsReturn *types.VcsRoot

	err := c.doRetryRequest("POST", path, vcs, &vcsReturn)
	if err != nil {
		return err
	}

	if vcsReturn == nil {
		return errors.New("VCS Root not created")
	}
	*vcs = *vcsReturn

	return nil
}
