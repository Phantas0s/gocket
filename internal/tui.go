package internal

// Layer of indirection to be able to swap the viewer easily.

type viewer interface {
	List(archiver func([]int), deleter func([]int), noconfirm bool)
}

type TUI struct {
	Viewer viewer
	Pocket *pocket
}

func (t *TUI) List(noconfirm bool) {
	t.Viewer.List(t.Pocket.Archive, t.Pocket.Delete, noconfirm)
}
