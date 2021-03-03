package internal

// Layer of indirection to be able to swap the viewer easily.

type viewer interface {
	List(archiver func([]int), deleter func([]int), adder func([]int), noconfirm bool)
}

type TUI struct {
	Viewer viewer
	Pocket *pocket
}

func (t *TUI) List(noconfirm bool) {
	t.Viewer.List(t.Pocket.Archive, t.Pocket.Delete, nil, noconfirm)
}

func (t *TUI) ListArchive(noconfirm bool) {
	t.Viewer.List(nil, t.Pocket.Delete, t.Pocket.Unarchive, noconfirm)
}
