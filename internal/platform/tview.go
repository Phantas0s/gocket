package platform

import (
	"os/exec"
	"runtime"

	"github.com/rivo/tview"
)

type Tview struct{}

func (v *Tview) Display(urls []string, titles []string) {
	app := tview.NewApplication()
	list := tview.NewList()

	for i, v := range titles {
		list.AddItem(v, urls[i], rune(i), func() {
			openBrowser(urls[i])
		})
	}

	list.AddItem("Quit", "Press to exit", 'q', func() {
		app.Stop()
	})

	if err := app.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
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
