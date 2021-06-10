package internal

// This layer isolate from the 3rd party Pocket API.
// Might be useless for now, but if I've learned something in development:
// ALWAYS isolate 3rd party APIs you have no control on.
// It will be easier to switch API version if needed for example.

import (
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

func (p *pocket) List(count int, order string, search string, filter string, tag string) (websites []Website) {
	res, err := p.client.Retrieve(count, order, search, filter, tag)
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

func (p *pocket) ListArchive(count int, order string, search string, filter string, tag string) (websites []Website) {
	res, err := p.client.RetrieveArchive(count, order, search, filter, tag)
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
	_, err := p.client.Archive(IDs)
	if err != nil {
		panic(err)
	}
}

func (p *pocket) Unarchive(IDs []int) {
	_, err := p.client.Unarchive(IDs)
	if err != nil {
		panic(err)
	}
}

func (p *pocket) Delete(IDs []int) {
	_, err := p.client.Delete(IDs)
	if err != nil {
		panic(err)
	}
}

func (p *pocket) Add(URL string) {
	err := p.client.Add(URL)
	if err != nil {
		panic(err)
	}
}
