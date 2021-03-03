package platform

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

const (
	EventDelete  = "delete"
	EventArchive = "archive"
)

type event struct {
	action string
	ID     int
	title  string
	listID int
}

type Tview struct {
	URLs   []string
	Titles []string
	IDs    []int
}

func (t *Tview) List(
	archiver func(IDs []int),
	deleter func(IDs []int),
	noconfirm bool,
) {
	app := tview.NewApplication()
	list := tview.NewList()

	list.SetSelectedTextColor(tcell.ColorBlack)
	list.SetSelectedBackgroundColor(tcell.ColorBlue)
	list.SetSelectedFocusOnly(true).
		SetBackgroundColor(tcell.ColorBlack).
		SetBorder(true).
		SetBorderColor(tcell.ColorWhite)

	pages := tview.NewPages()

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		}
		return event
	})

	for i, v := range t.Titles {
		url := t.URLs[i]
		list.AddItem(v, url, 0, func() {
			openBrowser(url)
		})

	}
	pages.AddPage("list", list, true, true)

	delete := tview.NewModal().AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SendToFront("list")
			pages.SendToBack("delete")
			pages.HidePage("delete")

			if buttonLabel == "Yes" {
				t.act(deleter, list)
			}
		}).SetFocus(0)

	archive := tview.NewModal().AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SendToFront("list")
			pages.SendToBack("archive")
			pages.HidePage("archive")

			if buttonLabel == "Yes" {
				t.act(archiver, list)
			}
		}).SetFocus(0)

	pages.AddPage("delete", delete, false, false)
	pages.AddPage("archive", archive, false, false)

	delete.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'j' {
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		} else if event.Rune() == 'k' {
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		}
		return event
	})

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'j' {
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		} else if event.Rune() == 'k' {
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		} else if event.Rune() == 'G' {
			return tcell.NewEventKey(tcell.KeyPgUp, 0, tcell.ModNone)
		} else if event.Rune() == 'g' {
			return tcell.NewEventKey(tcell.KeyPgDn, 0, tcell.ModNone)
		} else if event.Rune() == 'x' {
			if noconfirm {
				t.act(deleter, list)
			} else {
				title, _ := list.GetItemText(list.GetCurrentItem())
				delete.SetText(fmt.Sprintf("Are you sure you want to delete \"%s\"?", title))
				pages.SendToBack("list")
				pages.SendToFront("delete")
				pages.ShowPage("delete")
			}
		} else if event.Rune() == 'a' {
			if noconfirm {
				t.act(archiver, list)
			} else {
				title, _ := list.GetItemText(list.GetCurrentItem())
				archive.SetText(fmt.Sprintf("Are you sure you want to archive \"%s\"?", title))
				pages.SendToBack("list")
				pages.SendToFront("archive")
				pages.ShowPage("archive")
			}
		}

		return event
	})

	txt := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	txt.SetBorder(true).
		SetBorderColor(tcell.ColorWhite).
		SetBackgroundColor(tcell.ColorBlack)

	layout := tview.NewFlex().
		SetDirection(tview.FlexRow).
		AddItem(pages, 0, 10, true).
		AddItem(txt, 0, 1, false)

	fmt.Fprint(txt, "'a' to archive | 'x' to delete")

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func (t *Tview) act(action func(IDs []int), list *tview.List) {
	go action([]int{t.IDs[list.GetCurrentItem()]})
	t.URLs = append(t.URLs[:list.GetCurrentItem()], t.URLs[list.GetCurrentItem()+1:]...)
	t.Titles = append(t.Titles[:list.GetCurrentItem()], t.Titles[list.GetCurrentItem()+1:]...)
	t.IDs = append(t.IDs[:list.GetCurrentItem()], t.IDs[list.GetCurrentItem()+1:]...)
	list.RemoveItem(list.GetCurrentItem())
}

// open opens the specified URL in the default browser of the user.
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
