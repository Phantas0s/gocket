package cmd

import (
	"os"

	"github.com/Phantas0s/gocket/internal"
	"github.com/Phantas0s/gocket/internal/platform"
	"github.com/spf13/cobra"
)

// TODO add all these options into value object.
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

	listCmd.AddCommand(archiveCmd())
	listCmd.PersistentFlags().StringVarP(
		&order,
		"order",
		"o",
		"newest",
		"order by 'newest' (default), 'oldest', 'title', or 'url'.",
	)
	listCmd.PersistentFlags().StringVarP(&search, "search", "s", "", "Search by title and URL.")
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 0, "Number of results (0 for all).")

	listCmd.PersistentFlags().BoolVarP(&tui, "tui", "", false, "Display the results in a TUI.")
	listCmd.PersistentFlags().BoolVarP(&title, "title", "t", false, "Display the titles.")
	listCmd.Flags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles (with confirmation).")
	listCmd.PersistentFlags().BoolVarP(&delete, "delete", "d", false, "Delete the listed articles (with confirmation).")
	listCmd.PersistentFlags().BoolVarP(&noconfirm, "noconfirm", "", false, "Don't ask for any confirmation.")

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
				go pocket.Archive(IDs)
			}
		}

		if delete {
			if noconfirm || prompt("Do you really want to DELETE all the articles listed?") {
				go pocket.Delete(IDs)
			}
		}
	}
}
