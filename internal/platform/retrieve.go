package platform

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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
	Since       int64  `json:"since,omitempty"`
	Count       int    `json:"count,omitempty"`
	Offset      int    `json:"offset,omitempty"`
}

type retriveRequest struct {
	*RetrieveOption
	*Client
}

type RetrieveResult struct {
	List     []Item
	Status   int
	Complete int
	Since    int
}

func (r *RetrieveResult) UnmarshalJSON(data []byte) error {
	var d map[string]interface{}
	if err := json.Unmarshal(data, &d); err != nil {
		return err
	}

	switch v := d["list"].(type) {
	case []interface{}:
		if len(v) == 0 {
			r.List = nil
		}
	case map[string]interface{}:
		tmp := struct {
			List map[string]Item
		}{}
		if err := json.Unmarshal(data, &tmp); err != nil {
			return err
		}
		i := make([]Item, len(tmp.List))
		for _, v := range tmp.List {
			i[v.SortId] = v
		}
		r.List = i
	}

	r.Status = int(d["status"].(float64))
	r.Complete = int(d["complete"].(float64))
	r.Since = int(d["since"].(float64))

	return nil
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

func (c *Client) RetrieveArchive(
	count int,
	sort string,
	search string,
	filter string,
	tag string,
) (*RetrieveResult, error) {
	return c.fetch(count, sort, search, filter, tag, true)
}

func (c *Client) Retrieve(
	count int,
	sort string,
	search string,
	filter string,
	tag string,
) (*RetrieveResult, error) {
	return c.fetch(count, sort, search, filter, tag, false)
}

// Retrieve returns the in Pocket
func (c *Client) fetch(count int, sort string, search string, filter string, tag string, archive bool) (*RetrieveResult, error) {
	if sort == "" {
		sort = SortNewest
	}
	opts := &RetrieveOption{Sort: mapSort(sort)}
	if count != 0 {
		opts.Count = count
	}
	if search != "" {
		opts.Search = search
	}

	if tag != "" {
		opts.Tag = tag
	}

	if archive {
		opts.State = "archive"
	}

	opts.ContentType = filter

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

func mapSort(sort string) string {
	switch sort {
	case "newest":
		return SortNewest
	case "oldest":
		return SortOldest
	case "title":
		return SortTitle
	case "url":
		return SortSite
	default:
		os.Stderr.WriteString(fmt.Sprintf("ERROR: '%s' is not a valid order. Default to 'newest'.\n\n", sort))
		return SortNewest
	}
}
