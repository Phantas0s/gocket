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

func (c *Client) Add(URL string) error {
	data := addRequest{
		AddOption: &AddOption{URL: URL},
		Client:    c,
	}

	err := Post("/v3/add", data, &struct{}{})
	if err != nil {
		return err
	}

	return nil
}
