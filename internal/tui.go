package internal

import "github.com/Phantas0s/gocket/internal/platform"

// Layer of indirection to be able to swap the viewer easily.

type viewer interface {
	List(archiver func([]int), deleter func([]int), adder func([]int), noconfirm bool)
}

type TUI struct {
	Viewer viewer
	Pocket *pocket
}

func (t *TUI) List(list []Website, noconfirm bool) {
	t.Viewer = &platform.Tview{Entries: toEntries(list)}
	t.Viewer.List(t.Pocket.Archive, t.Pocket.Delete, nil, noconfirm)
}

func (t *TUI) ListArchive(list []Website, noconfirm bool) {
	t.Viewer = &platform.Tview{Entries: toEntries(list)}
	t.Viewer.List(nil, t.Pocket.Delete, t.Pocket.Unarchive, noconfirm)
}

func toEntries(list []Website) []platform.Entry {
	e := make([]platform.Entry, len(list))
	for i, v := range list {
		e[i] = platform.Entry{URL: v.URL, Title: v.Title, ID: v.ID}
	}

	return e
}
