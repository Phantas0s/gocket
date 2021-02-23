package platform

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Tview struct{}

func (v *Tview) List(urls []string, titles []string) {
	app := tview.NewApplication()
	list := tview.NewList()

	list.SetSelectedTextColor(tcell.ColorRed)
	list.SetSelectedBackgroundColor(tcell.ColorBlack)
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

	f := func(mod *tview.Modal, action string) func(int, string, string, rune) {
		return func(i int, main string, sec string, r rune) {
			mod.SetText(fmt.Sprintf("Are you sure you want to delete \"%s\"?", main))
		}
	}

	for i, v := range titles {
		url := urls[i]
		list.AddItem(v, url, rune(i), func() {
			openBrowser(url)
		})

	}
	pages.AddPage("list", list, true, true)

	delete := tview.NewModal().AddButtons([]string{"Yes", "No"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			pages.SendToFront("list")
			pages.SendToBack("delete")
			pages.HidePage("delete")
		}).SetFocus(0)
	list.SetChangedFunc(f(delete, "delete"))
	pages.AddPage("delete", delete, false, false)

	list.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'j' {
			list.SetCurrentItem(list.GetCurrentItem() + 1)
		} else if event.Rune() == 'k' {
			list.SetCurrentItem(list.GetCurrentItem() - 1)
		} else if event.Rune() == 'x' {
			pages.SendToBack("list")
			pages.SendToFront("delete")
			pages.ShowPage("delete")
		}

		return event
	})

	txt := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	txt.SetBorder(true).
		SetBorderColor(tcell.ColorWhite).
		SetBackgroundColor(tcell.ColorDefault)

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
