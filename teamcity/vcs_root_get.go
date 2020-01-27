package teamcity

import (
	"fmt"

	"github.com/icelander/teamcity-sdk-go/types"
)

func (c *Client) GetVcsRoot(VcsRootId string) (*types.VcsRoot, error) {
	path := fmt.Sprintf("/app/rest/vcs-roots/id:%s", VcsRootId)
	var vcs *types.VcsRoot

	err := c.doRetryRequest("GET", path, nil, &vcs)
	if err != nil {
		return nil, err
	}

	return vcs, nil
}
