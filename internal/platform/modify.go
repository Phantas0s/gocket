package platform

const (
	add        = "add"
	archive    = "archive"
	unarchive  = "readd"
	favorite   = "favorite"
	unfavorite = "unfavorite"
	delete     = "delete"

	addTag     = "tags_add"
	removeTag  = "tags_remove"
	replaceTag = "tags_replace"
	clearTag   = "tags_clear"
	renameTag  = "tag_rename"
	deleteTag  = "tag_delete"
)

// Action represents one action in a bulk modify requests.
type Modify struct {
	Action string `json:"action"`
	ID     int    `json:"item_id,string"`
}

type modifyRequest struct {
	Modifies []Modify `json:"actions"`
	*Client
}

type ModifyResult struct {
	// The results for each of the requested actions.
	ActionResults []bool
	Status        int
}

func (c *Client) Archive(IDs []int) (*ModifyResult, error) {
	acs := []Modify{}
	for _, v := range IDs {
		acs = append(acs, Modify{Action: archive, ID: v})
	}
	return c.modify(acs)
}

func (c *Client) Unarchive(IDs []int) (*ModifyResult, error) {
	acs := []Modify{}
	for _, v := range IDs {
		acs = append(acs, Modify{Action: unarchive, ID: v})
	}
	return c.modify(acs)
}

func (c *Client) Delete(IDs []int) (*ModifyResult, error) {
	acs := []Modify{}
	for _, v := range IDs {
		acs = append(acs, Modify{Action: delete, ID: v})
	}
	return c.modify(acs)
}

func (c *Client) modify(modifies []Modify) (*ModifyResult, error) {
	data := modifyRequest{
		Client:   c,
		Modifies: modifies,
	}

	res := ModifyResult{}
	err := Post("/v3/send", data, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
