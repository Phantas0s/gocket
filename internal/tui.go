package internal

// Layer of indirection to be able to swap the viewer easily.

type viewer interface {
	List(urls []string, titles []string)
}

type TUI struct {
	Instance viewer
}

func (t *TUI) List(websites []Website) {
	urls := []string{}
	titles := []string{}
	for _, v := range websites {
		urls = append(urls, v.URL)
		titles = append(titles, v.Title)
	}
	t.Instance.List(urls, titles)
}
