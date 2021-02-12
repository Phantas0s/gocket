package internal

import (
	"fmt"

	"github.com/Phantas0s/gocket/internal/platform"
)

func DisplayList(
	consumerKey string,
	browser string,
	count int,
) error {
	auth, err := platform.Auth(consumerKey, browser)
	c := platform.NewClient(consumerKey, auth.AccessToken)

	res, err := c.Retrieve(&platform.RetrieveOption{})
	if err != nil {
		panic(err)
	}

	if count == 0 {
		for _, e := range res.List {
			fmt.Println(e.URL())
		}
	} else {
		l := res.FlattenList()
		for i := 0; i < count; i++ {
			fmt.Println(l[i].URL())
		}
	}

	return nil
}
