package cmd

import (
	"bufio"
	"os"

	"github.com/Phantas0s/gocket/internal"
	"github.com/Phantas0s/gocket/internal/platform"
	"github.com/spf13/cobra"
)

var consumerKey, sort string
var count int
var tui, archive bool

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Pocket consumer key (required)")
	listCmd.PersistentFlags().StringVarP(
		&sort,
		"sort",
		"s",
		"newest",
		"Sort by 'newest' (default), 'oldest', 'title' or 'url'",
	)
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 10, "Number of results (0 for all, default 10)")
	listCmd.PersistentFlags().BoolVarP(&tui, "tui", "t", false, "Display the TUI")
	listCmd.PersistentFlags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles")
	// TODO
	listCmd.PersistentFlags().BoolVarP(&archive, "noprompt", "n", false, "Doesn't ask you anything (DANGEROUS)")

	listCmd.MarkFlagRequired("key")

	// Cobra shenanigans
	// viper.BindPFlag("key", listCmd.PersistentFlags().Lookup("key"))
	// viper.BindPFlag("count", listCmd.PersistentFlags().Lookup("count"))
	// viper.BindPFlag("sort", listCmd.PersistentFlags().Lookup("sort"))
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
	list := pocket.List(count, sort)
	if tui {
		tui := internal.TUI{Instance: &platform.Tview{}}
		tui.Display(list)
	} else {
		IDs := []int{}
		for _, v := range list {
			os.Stdout.WriteString(v.URL + "\n")
			IDs = append(IDs, v.ID)
			// os.Stdout.WriteString(v.URL + " ")
		}
		if archive {
			os.Stdout.WriteString("Are you sure you want to archive all the articles above? (y/n)")
			reader := bufio.NewReader(os.Stdin)
			i, err := reader.ReadString('\n')
			if err != nil {
				panic(err)
			}
			if string(i) == "y\n" {
				pocket.Archive(IDs)
			} else {
				os.Stdout.WriteString("Aborted.")
			}
		}
	}
}
