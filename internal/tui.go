package internal

// Layer of indirection to be able to swap the viewer easily.

type viewer interface {
	List(urls []string, titles []string, IDs []int, archiver func([]int), deleter func([]int))
}

type TUI struct {
	Instance viewer
	Pocket   *pocket
}

func (t *TUI) List(websites []Website) {
	urls := make([]string, len(websites))
	titles := make([]string, len(websites))
	IDs := make([]int, len(websites))

	for k, v := range websites {
		IDs[k] = v.ID
		urls[k] = v.URL
		titles[k] = v.Title
	}

	t.Instance.List(urls, titles, IDs, t.Pocket.Archive, t.Pocket.Delete)
}
