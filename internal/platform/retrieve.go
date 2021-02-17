package platform

import (
	"bytes"
	"strconv"
	"time"
)

const (
	ContentTypeArticle = "article"
	ContentTypeVideo   = "video"
	ContentTypeImage   = "image"

	SortNewest = "newest"
	SortOldest = "oldest"
	SortTitle  = "title"
	SortSite   = "site"

	StateUnread  = "unread"
	StateArchive = "archive"
	StateAll     = "all"

	DetailTypeSimple   = "simple"
	DetailTypeComplete = "complete"

	FavoriteFilterUnspecified = ""
	FavoriteFilterUnfavorited = "0"
	FavoriteFilterFavorited   = "1"

	ItemStatusUnread   = 0
	ItemStatusArchived = 1
	ItemStatusDeleted  = 2

	ItemMediaAttachmentNoMedia  = 0
	ItemMediaAttachmentHasMedia = 1
	ItemMediaAttachmentIsMedia  = 2
)

// RetrieveOption is the options for retrieve API.
type RetrieveOption struct {
	State       string `json:"state,omitempty"`
	Favorite    string `json:"favorite,omitempty"`
	Tag         string `json:"tag,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Sort        string `json:"sort,omitempty"`
	DetailType  string `json:"detailType,omitempty"`
	Search      string `json:"search,omitempty"`
	Domain      string `json:"domain,omitempty"`
	Since       int    `json:"since,omitempty"`
	Count       int    `json:"count,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}

type retriveRequest struct {
	*RetrieveOption
	*Client
}

type RetrieveResult struct {
	List     map[string]Item
	Status   int
	Complete int
	Since    int
}

func (r RetrieveResult) FlattenList() []Item {
	newList := []Item{}
	for _, v := range r.List {
		newList = append(newList, v)
	}
	return newList
}

type Item struct {
	ItemID        int    `json:"item_id,string"`
	ResolvedId    int    `json:"resolved_id,string"`
	GivenURL      string `json:"given_url"`
	ResolvedURL   string `json:"resolved_url"`
	GivenTitle    string `json:"given_title"`
	ResolvedTitle string `json:"resolved_title"`
	Favorite      int    `json:",string"`
	Status        int    `json:",string"`
	Excerpt       string
	IsArticle     int `json:"is_article,string"`
	HasImage      int `json:"has_image,string"`
	HasVideo      int `json:"has_video,string"`
	WordCount     int `json:"word_count,string"`

	// Fields for detailed response
	Tags    map[string]map[string]interface{}
	Authors map[string]map[string]interface{}
	Images  map[string]map[string]interface{}
	Videos  map[string]map[string]interface{}

	// Fields that are not documented but exist
	SortId        int  `json:"sort_id"`
	TimeAdded     Time `json:"time_added"`
	TimeUpdated   Time `json:"time_updated"`
	TimeRead      Time `json:"time_read"`
	TimeFavorited Time `json:"time_favorited"`
}

type Time time.Time

func (t *Time) UnmarshalJSON(b []byte) error {
	i, err := strconv.ParseInt(string(bytes.Trim(b, `"`)), 10, 64)
	if err != nil {
		return err
	}

	*t = Time(time.Unix(i, 0))

	return nil
}

// URL returns ResolvedURL or GivenURL
func (item Item) URL() string {
	url := item.ResolvedURL
	if url == "" {
		url = item.GivenURL
	}
	return url
}

// Title returns ResolvedTitle or GivenTitle
func (item Item) Title() string {
	title := item.ResolvedTitle
	if title == "" {
		title = item.GivenTitle
	}
	return title
}

// Retrieve returns the in Pocket
func (c *Client) Retrieve(count int, sort string) (*RetrieveResult, error) {
	opts := &RetrieveOption{Sort: sort}
	if count != 0 {
		opts.Count = count
	}

	data := retriveRequest{
		Client:         c,
		RetrieveOption: opts,
	}

	res := &RetrieveResult{}
	err := Post("/v3/get", data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
