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
	listID int
}

type Tview struct{}

func (v *Tview) List(
	URLs []string,
	titles []string,
	IDs []int,
	archiver func(IDs []int),
	deleter func(IDs []int),
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

	for i, v := range titles {
		url := URLs[i]
		list.AddItem(v, url, 0, func() {
			openBrowser(url)
		})

	}
	pages.AddPage("list", list, true, true)

	e := event{}
	listChangedFunc := func(mod *tview.Modal, action string, e *event) func(int, string, string, rune) {
		return func(i int, main string, sec string, r rune) {
			mod.SetText(fmt.Sprintf("Are you sure you want to %s \"%s\"?", action, main))
			e.ID = IDs[i]
			e.listID = i
		}
	}

	delete := tview.NewModal().AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SendToFront("list")
			pages.SendToBack("delete")
			pages.HidePage("delete")
			if buttonLabel == "Yes" {
				deleter([]int{e.ID})
				list.RemoveItem(e.listID)
				URLs = append(URLs[:e.listID], URLs[e.listID+1:]...)
				titles = append(titles[:e.listID], titles[e.listID+1:]...)
				IDs = append(IDs[:e.listID], IDs[e.listID+1:]...)
			}
		}).SetFocus(0)

	archive := tview.NewModal().AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SendToFront("list")
			pages.SendToBack("archive")
			pages.HidePage("archive")
			if buttonLabel == "Yes" {
				archiver([]int{e.ID})
				list.RemoveItem(e.listID)
				URLs = append(URLs[:e.listID], URLs[e.listID+1:]...)
				titles = append(titles[:e.listID], titles[e.listID+1:]...)
				IDs = append(IDs[:e.listID], IDs[e.listID+1:]...)
			}
		}).SetFocus(0)
	list.SetChangedFunc(listChangedFunc(archive, "archive", &e))

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
			//TODO CTRL+D / CTRL+U?
		} else if event.Rune() == 'x' {
			pages.SendToBack("list")
			pages.SendToFront("delete")
			pages.ShowPage("delete")
		} else if event.Rune() == 'a' {
			pages.SendToBack("list")
			pages.SendToFront("archive")
			pages.ShowPage("archive")
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

func help(app *tview.Application) {
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
