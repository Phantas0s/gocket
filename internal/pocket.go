package internal

// This layer isolate from the 3rd party Pocket API.
// Might be useless for now, but if I've learned something in development:
// ALWAYS isolate 3rd party APIs you have no control on.
// It will be easier to switch API version for example.

import (
	"fmt"
	"os"

	"github.com/Phantas0s/gocket/internal/platform"
)

type pocket struct {
	client *platform.Client
}

type Website struct {
	ID    int
	Title string
	URL   string
}

func CreatePocket(consumerKey string) *pocket {
	auth, err := platform.Auth(consumerKey)
	if err != nil {
		panic(err)
	}

	c := platform.NewClient(consumerKey, auth.AccessToken)

	return &pocket{
		client: c,
	}
}

func (p *pocket) List(count int, sort string) (websites []Website) {
	res, err := p.client.Retrieve(count, mapSort(sort))
	if err != nil {
		panic(err)
	}

	for _, e := range res.List {
		websites = append(websites, Website{
			ID:    e.ItemID,
			Title: e.Title(),
			URL:   e.URL(),
		})
	}

	return
}

func (p *pocket) Archive(IDs []int) {
	//TODO do something with result?
	_, err := p.client.Archive(IDs)
	if err != nil {
		panic(err)
	}
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
