package cmd

import (
	"os"

	"github.com/Phantas0s/gocket/internal"
	"github.com/Phantas0s/gocket/internal/platform"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var consumerKey, browser, sort string
var count int
var tui bool

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

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Pocket consumer key (required)")
	listCmd.PersistentFlags().StringVarP(&browser, "browser", "b", "", "Browser to open URLs (included the auth URL)")
	listCmd.PersistentFlags().StringVarP(
		&sort,
		"sort",
		"s",
		"newest",
		"Sort by 'newest' (default), 'oldest', 'title' or 'url'",
	)
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 10, "Number of results (0 for all, default 10)")
	listCmd.PersistentFlags().BoolVarP(&tui, "tui", "t", false, "Display the TUI")

	listCmd.MarkFlagRequired("key")

	viper.BindPFlag("key", listCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("browser", listCmd.PersistentFlags().Lookup("browser"))
	viper.BindPFlag("count", listCmd.PersistentFlags().Lookup("count"))
	viper.BindPFlag("sort", listCmd.PersistentFlags().Lookup("sort"))
	viper.BindPFlag("output", listCmd.PersistentFlags().Lookup("output"))
}

func runList() {
	list := internal.List(consumerKey, browser, count, sort)
	if tui {
		tui := internal.TUI{Instance: &platform.Tview{}}
		tui.Display(list)
	} else {
		// TODO put that in a proper render function
		for _, v := range list {
			os.Stdout.WriteString(v.URL + "\n")
			// os.Stdout.WriteString(v.URL + " ")
		}
	}
}
