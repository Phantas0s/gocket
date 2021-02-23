package cmd

import (
	"bufio"
	"os"
	"strings"

	"github.com/Phantas0s/gocket/internal"
	"github.com/Phantas0s/gocket/internal/platform"
	"github.com/spf13/cobra"
)

var consumerKey, order, search string
var count int
var tui, archive, delete, noprompt, title bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Pocket consumer key (required).")
	listCmd.PersistentFlags().StringVarP(
		&order,
		"order",
		"o",
		"newest",
		"order by 'newest' (default), 'oldest', 'title', or 'url'.",
	)
	listCmd.PersistentFlags().StringVarP(&search, "search", "s", "", "Only list items with title or URL matching the search.")
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 10, "Number of results (0 for all, default 10).")

	listCmd.PersistentFlags().BoolVarP(&tui, "tui", "", false, "Display the TUI.")
	listCmd.PersistentFlags().BoolVarP(&title, "title", "t", false, "Display the title the line before the URL.")
	listCmd.PersistentFlags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles.")
	listCmd.PersistentFlags().BoolVarP(&delete, "delete", "d", false, "Delete the listed articles.")

	listCmd.MarkFlagRequired("key")

	// Cobra shenanigans
	// viper.BindPFlag("key", listCmd.PersistentFlags().Lookup("key"))
	// viper.BindPFlag("count", listCmd.PersistentFlags().Lookup("count"))
	// viper.BindPFlag("order", listCmd.PersistentFlags().Lookup("order"))
	// viper.BindPFlag("output", listCmd.PersistentFlags().Lookup("output"))
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your pocket articles",
	// TODO write some help
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		runList()
	},
}

func runList() {
	pocket := internal.CreatePocket(consumerKey)
	list := pocket.List(count, order, search)
	if tui {
		tui := internal.TUI{Instance: &platform.Tview{}}
		tui.List(list)
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
			if noprompt || prompt("Do you really want to archive all the articles listed?") {
				pocket.Archive(IDs)
			}
		}

		if delete {
			if noprompt || prompt("Do you really want to DELETE all the articles listed?") {
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
