package platform

import (
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tview struct{}

func (v *Tview) List(urls []string, titles []string) {
	app := tview.NewApplication()
	list := tview.NewList()
	list.SetBackgroundColor(tcell.ColorDefault)
	list.SetSelectedBackgroundColor(tcell.ColorDefault)
	list.SetSelectedTextColor(tcell.ColorRed)
	list.SetSelectedFocusOnly(true)

	pages := tview.NewPages()
	modal := tview.NewModal().SetText("Hello").SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		pages.HidePage("modal")
	})

	pages.AddPage("modal", modal, false, false)

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'H' {
			pages.ShowPage("modal")
		}
		return event
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})

	for i, v := range titles {
		url := urls[i]
		list.AddItem(v, url, rune(i), func() {
			openBrowser(url)
		})

	}

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'j' {
			list.SetCurrentItem(list.GetCurrentItem() + 1)
		} else if event.Rune() == 'k' {
			list.SetCurrentItem(list.GetCurrentItem() - 1)
		}

		return event
	})

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	showpage := false
	modal.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == '?' {
			showpage = !showpage
			pages.HidePage("modal")
		}
		return event
	})

	// Create the main layout.
	layout := tview.NewFlex().
		AddItem(list, 0, 1, false).
		AddItem(pages, 0, 0, false)

	layout.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == '?' {
			showpage = !showpage
			if showpage {
				pages.ShowPage("modal")
			}
		}
		return event
	})

	if err := app.SetRoot(layout, true).SetFocus(list).Run(); err != nil {
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
