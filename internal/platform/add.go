package platform

type AddOption struct {
	URL   string `json:"url,omitempty"`
	Title string `json:"title,omitempty"`
	Tags  string `json:"tags,omitempty"`
}

type addRequest struct {
	*AddOption
	*Client
}

type AddResult struct{}

// Add only returns an error status, since adding an article doesn't have
// any other meaningful return value.
func (c *Client) Add(URL string) error {
	data := addRequest{
		AddOption: &AddOption{URL: URL},
		Client:    c,
	}

	res := &AddResult{}
	err := Post("/v3/add", data, res)
	if err != nil {
		return err
	}

	return nil
}
