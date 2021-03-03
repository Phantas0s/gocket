package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/Phantas0s/gocket/internal"
	"github.com/Phantas0s/gocket/internal/platform"
	"github.com/spf13/cobra"
)

var order, search string
var count int
var tui, archive, delete, noconfirm, title bool

func ListCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List your pocket articles",
		// TODO write some help
		Run: func(cmd *cobra.Command, args []string) {
			runList()
		},
	}

	listCmd.Flags().StringVarP(
		&order,
		"order",
		"o",
		"newest",
		"order by 'newest' (default), 'oldest', 'title', or 'url'.",
	)
	listCmd.Flags().StringVarP(&search, "search", "s", "", "Only list items with title or URL matching the search.")
	listCmd.Flags().IntVarP(&count, "count", "c", 10, "Number of results (0 for all, default 10).")

	listCmd.Flags().BoolVarP(&tui, "tui", "", false, "Display the TUI.")
	listCmd.Flags().BoolVarP(&title, "title", "t", false, "Display the title the line before the URL.")
	listCmd.Flags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles.")
	listCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete the listed articles.")
	listCmd.Flags().BoolVarP(&noconfirm, "noconfirm", "", false, "Don't ask for any confirmation.")

	return listCmd
}

func runList() {
	pocket := internal.CreatePocket(consumerKey)
	list := pocket.List(count, order, search)
	if tui {
		tv := platform.Tview{
			IDs:    make([]int, len(list)),
			URLs:   make([]string, len(list)),
			Titles: make([]string, len(list)),
		}
		for k, v := range list {
			tv.IDs[k] = v.ID
			tv.URLs[k] = v.URL
			tv.Titles[k] = v.Title
		}

		tui := internal.TUI{Viewer: &tv, Pocket: pocket}
		tui.List(noconfirm)
	} else {
		IDs := []int{}
		for _, v := range list {
			IDs = append(IDs, v.ID)
			if title {
				os.Stdout.WriteString(v.Title + "\n")
			}
			os.Stdout.WriteString(v.URL + "\n")
		}

		if archive {
			if noconfirm || prompt("Do you really want to archive all the articles listed?") {
				pocket.Archive(IDs)
			}
		}

		if delete {
			if noconfirm || prompt("Do you really want to DELETE all the articles listed?") {
				pocket.Delete(IDs)
			}
		}
	}
}

func prompt(message string) bool {
	os.Stdout.WriteString(message + " (y/n)")
	reader := bufio.NewReader(os.Stdin)
	i, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	if strings.Trim(string(i), "\n") == "y" {
		return true
	} else {
		os.Stdout.WriteString("Aborted.")
		return false
	}
}
