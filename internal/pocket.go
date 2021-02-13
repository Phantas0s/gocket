package internal

import (
	"github.com/Phantas0s/gocket/internal/platform"
)

func List(consumerKey string, browser string, count int) map[string]string {
	result := make(map[string]string, count)
	auth, err := platform.Auth(consumerKey, browser)
	c := platform.NewClient(consumerKey, auth.AccessToken)

	res, err := c.Retrieve(&platform.RetrieveOption{})
	if err != nil {
		panic(err)
	}

	if count == 0 {
		for _, e := range res.List {
			result[e.URL()] = e.Title()
		}
	} else {
		l := res.FlattenList()
		for i := 0; i < count; i++ {
			result[l[i].URL()] = l[i].Title()
		}
	}

	return result
}
