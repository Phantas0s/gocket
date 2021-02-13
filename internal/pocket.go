package internal

import (
	"github.com/Phantas0s/gocket/internal/platform"
)

type Website struct {
	Title string
	URL   string
}

func List(consumerKey string, browser string, count int) (websites []Website) {
	auth, err := platform.Auth(consumerKey, browser)
	c := platform.NewClient(consumerKey, auth.AccessToken)

	opts := &platform.RetrieveOption{Sort: platform.SortNewest}
	if count != 0 {
		opts.Count = count
	}

	res, err := c.Retrieve(opts)
	if err != nil {
		panic(err)
	}

	for _, e := range res.List {
		websites = append(websites, Website{
			Title: e.Title(),
			URL:   e.URL(),
		})
	}

	return
}
