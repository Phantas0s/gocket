package internal

import (
	"fmt"
	"os"

	"github.com/Phantas0s/gocket/internal/platform"
)

type Website struct {
	Title string
	URL   string
}

func List(
	consumerKey string,
	browser string,
	count int,
	sort string,
) (websites []Website) {
	auth, err := platform.Auth(consumerKey, browser)
	c := platform.NewClient(consumerKey, auth.AccessToken)

	opts := &platform.RetrieveOption{Sort: mapSort(sort)}
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

func mapSort(sort string) string {
	switch sort {
	case "newest":
		return platform.SortNewest
	case "oldest":
		return platform.SortOldest
	case "title":
		return platform.SortTitle
	case "url":
		return platform.SortSite
	default:
		os.Stderr.WriteString(fmt.Sprintf("ERROR: '%s' is not a valid sort. Default to 'newest'.\n\n", sort))
		return platform.SortNewest
	}
}
