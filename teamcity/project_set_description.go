package teamcity

func (c *Client) SetProjectDescription(projectID, description string) error {
	err := c.SetProjectField(projectID, "description", description)
	if err != nil {
		return err
	}
	return nil
}
