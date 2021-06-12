package cmd

import (
	"os"

	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
)

// TODO add all these options into value object.
var order, search, filter, tag string
var count int
var tui, archive, delete, noconfirm, title bool

func listCmd() *cobra.Command {
	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List your Pocket pages",
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
		"order by 'newest', 'oldest', 'title', or 'url'",
	)
	listCmd.PersistentFlags().StringVarP(&search, "search", "s", "", "Search by title and URL")
	listCmd.PersistentFlags().StringVarP(&filter, "filter", "f", "article", "filter by type ('article', 'video', 'image')")
	listCmd.PersistentFlags().StringVarP(&tag, "tag", "", "tag", "filter by tag")
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 0, "Number of results (0 for all)")

	listCmd.PersistentFlags().BoolVarP(&tui, "tui", "", false, "Display the results in a TUI")
	listCmd.PersistentFlags().BoolVarP(&title, "title", "t", false, "Display the titles")
	listCmd.Flags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles (with confirmation)")
	listCmd.PersistentFlags().BoolVarP(&delete, "delete", "d", false, "Delete the listed articles (with confirmation)")
	listCmd.PersistentFlags().BoolVarP(&noconfirm, "noconfirm", "", false, "Don't ask for any confirmation")

	return listCmd
}

func runList() {
	pocket := internal.CreatePocket(consumerKey)
	list := pocket.List(count, order, search, filter, tag)
	if tui {
		tui := internal.TUI{Pocket: pocket}
		tui.List(list, noconfirm)
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
			if noconfirm || prompt("Do you really want to ARCHIVE all the articles listed?") {
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
