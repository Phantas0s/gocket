package cmd

import (
	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add URL...",
	Short: "Add URLs to Pocket",
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func add(URLs []string) {
	pocket := internal.CreatePocket(consumerKey)

	for _, v := range URLs {
		pocket.Add(v)
	}
}
