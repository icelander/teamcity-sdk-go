package teamcity

import (
	"fmt"

	"github.com/icelander/teamcity-sdk-go/types"
)

func (c *Client) GetAgentPoolById(pool int) (*types.AgentPools, error) {
	path := fmt.Sprintf("/app/rest/agentPools/id:%d", pool)
	var agp *types.AgentPools

	err := c.doRetryRequest("GET", path, nil, &agp)
	if err != nil {
		return nil, err
	}

	return agp, nil
}

func (c *Client) GetAgentPoolByName(pool string) (*types.AgentPools, error) {
	path := fmt.Sprintf("/app/rest/agentPools/name:%s", pool)
	var agp *types.AgentPools

	err := c.doRetryRequest("GET", path, nil, &agp)
	if err != nil {
		return nil, err
	}

	return agp, nil
}
