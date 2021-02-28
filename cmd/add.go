package cmd

import (
	"fmt"
	"os"

	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a pocket article",
	// TODO write some help (?)
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		add(args)
	},
}

func add(URLs []string) {
	pocket := internal.CreatePocket(consumerKey)

	for _, v := range URLs {
		pocket.Add(v)
		os.Stdout.WriteString(fmt.Sprintf("%s has been added\n", v))
	}
}
