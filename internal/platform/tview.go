package platform

import (
	"fmt"
	"os/exec"
	"runtime"
	"strings"

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

type Entry struct {
	URL   string
	Title string
	ID    int
}

type Tview struct {
	Entries []Entry
}

func (t *Tview) List(
	archiver func(IDs []int),
	deleter func(IDs []int),
	adder func(IDs []int),
	noconfirm bool,
) {
	app := tview.NewApplication()
	list := tview.NewList()

	list.SetSelectedTextColor(tcell.ColorBlack)
	list.SetSelectedBackgroundColor(tcell.ColorBlue)
	// Don't go to the end of the list if go up at the beginning of it (and vice-versa).
	list.SetWrapAround(false)
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

	for _, v := range t.Entries {
		url := v.URL
		list.AddItem(v.Title, url, 0, func() {
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

	extraAction := tview.NewModal()
	extraName := ""
	if adder == nil {
		extraName = "archive"
		extraAction.AddButtons([]string{"Yes", "No"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.SendToFront("list")
				pages.SendToBack("archive")
				pages.HidePage("archive")

				if buttonLabel == "Yes" {
					t.act(archiver, list)
				}
			}).SetFocus(0)
	} else if archiver == nil {
		extraName = "add"
		extraAction.AddButtons([]string{"Yes", "No"}).
			SetDoneFunc(func(buttonIndex int, buttonLabel string) {
				pages.SendToFront("list")
				pages.SendToBack("add")
				pages.HidePage("add")

				if buttonLabel == "Yes" {
					t.act(adder, list)
				}
			}).SetFocus(0)
	}

	pages.AddPage("delete", delete, false, false)
	pages.AddPage("extra", extraAction, false, false)

	delete.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'j' {
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		} else if event.Rune() == 'k' {
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		}
		return event
	})

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch key := event.Rune(); key {
		case 'j':
			return tcell.NewEventKey(tcell.KeyDown, 0, tcell.ModNone)
		case 'k':
			return tcell.NewEventKey(tcell.KeyUp, 0, tcell.ModNone)
		case 'G':
			return tcell.NewEventKey(tcell.KeyEnd, 0, tcell.ModNone)
		case 'g':
			return tcell.NewEventKey(tcell.KeyHome, 0, tcell.ModNone)
		case 'x':
			if noconfirm {
				t.act(deleter, list)
			} else {
				title, _ := list.GetItemText(list.GetCurrentItem())
				delete.SetText(fmt.Sprintf("Are you sure you want to delete \"%s\"?", title))
				pages.SendToBack("list")
				pages.SendToFront("delete")
				pages.ShowPage("delete")
			}
		case 'a':
			if noconfirm {
				t.act(archiver, list)
			} else {
				title, _ := list.GetItemText(list.GetCurrentItem())
				extraAction.SetText(fmt.Sprintf("Are you sure you want to %s \"%s\"?", extraName, title))
				pages.SendToBack("list")
				pages.SendToFront("extra")
				pages.ShowPage("extra")
			}
		}

		switch key := event.Key(); key {
		case tcell.KeyCtrlD:
			return tcell.NewEventKey(tcell.KeyPgDn, 0, tcell.ModNone)
		case tcell.KeyCtrlU:
			return tcell.NewEventKey(tcell.KeyPgUp, 0, tcell.ModNone)
		}

		return event
	})

	statusBar := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	statusBar.SetBorder(true).
		SetBorderColor(tcell.ColorWhite).
		SetBackgroundColor(tcell.ColorBlack)

	layout := tview.NewGrid().
		SetRows(0, 3).
		SetColumns(0).
		SetBorders(false).
		AddItem(pages, 0, 0, 1, 1, 0, 0, true).
		AddItem(statusBar, 1, 0, 1, 1, 0, 0, false)

	fmt.Fprint(statusBar, fmt.Sprintf("'a' to %s | 'x' to delete", extraName))

	if err := app.SetRoot(layout, true).Run(); err != nil {
		panic(err)
	}
}

func (t *Tview) act(action func(IDs []int), list *tview.List) {
	go action([]int{t.Entries[list.GetCurrentItem()].ID})
	t.Entries = append(t.Entries[:list.GetCurrentItem()], t.Entries[list.GetCurrentItem()+1:]...)
	list.RemoveItem(list.GetCurrentItem())
}

// open opens the specified URL in the default browser of the user.
func openBrowser(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		url = strings.Replace(url, "&", "^&", -1)
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
