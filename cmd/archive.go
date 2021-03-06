package cmd

import (
	"os"

	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
)

// TODO encapsulate option variables into value objects.
var a bool

func archiveCmd() *cobra.Command {
	listArchiveCmd := &cobra.Command{
		Use:   "archive",
		Short: "List your archive",
		// TODO write some help
		Run: func(cmd *cobra.Command, args []string) {
			runArchive()
		},
	}
	listArchiveCmd.Flags().BoolVarP(&a, "add", "a", false, "Add the listed articles (with confirmation).")

	return listArchiveCmd
}

func runArchive() {
	pocket := internal.CreatePocket(consumerKey)
	list := pocket.ListArchive(count, order, search)
	if tui {
		tui := internal.TUI{Pocket: pocket}
		tui.ListArchive(list, noconfirm)
	} else {
		IDs := []int{}
		for _, v := range list {
			IDs = append(IDs, v.ID)
			if title {
				os.Stdout.WriteString(v.Title + "\n")
			}
			os.Stdout.WriteString(v.URL + "\n")
		}

		if a {
			if noconfirm || prompt("Do you really want to add all the articles listed?") {
				pocket.Unarchive(IDs)
			}
		}
	}
}
