package cmd

import (
	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var consumerKey, browser string
var count int

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all pocket articles",
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
	listCmd.PersistentFlags().StringVarP(&browser, "browser", "b", "", "Broswer to open the authorization URL")
	listCmd.PersistentFlags().IntVarP(&count, "count", "c", 0, "Number of results (0 for all)")

	listCmd.MarkFlagRequired("key")

	viper.BindPFlag("key", listCmd.PersistentFlags().Lookup("key"))
	viper.BindPFlag("browser", listCmd.PersistentFlags().Lookup("browser"))
	viper.BindPFlag("count", listCmd.PersistentFlags().Lookup("count"))
}

func runList() {
	internal.DisplayList(consumerKey, browser, count)
}
