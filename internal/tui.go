package internal

// Layer of indirection to be able to swap the viewer easily.

type viewer interface {
	Display(urls []string, titles []string)
}

type TUI struct {
	Instance viewer
}

func (t *TUI) Display(websites []Website) {
	urls := []string{}
	titles := []string{}
	for _, v := range websites {
		urls = append(urls, v.URL)
		titles = append(titles, v.Title)
	}
	t.Instance.Display(urls, titles)
}
