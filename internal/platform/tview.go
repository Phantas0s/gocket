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

	mapping(app)

	for i, v := range titles {
		url := urls[i]
		list.AddItem(v, url, rune(i), func() {
			openBrowser(url)
		})

	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}

func help(app *tview.Application) {
	app.NewModal().SetText("Help")
}

func mapping(app *tview.Application) {
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'x' {
			help(app)
		} else if event.Rune() == '?' {
		}
		return event
	})

	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		return event
	})
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
