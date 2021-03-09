package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Phantas0s/gocket/internal"
	"github.com/spf13/cobra"
)

var verbose bool

func addCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add URL...",
		Short: "Add URLs to Pocket",
		Run: func(cmd *cobra.Command, args []string) {
			add(args)
		},
	}

	addCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Output success messages.")

	return addCmd
}

func add(URLs []string) {
	if len(URLs) == 0 {
		URLs = readStdin()
	}

	pocket := internal.CreatePocket(consumerKey)
	for _, v := range URLs {
		pocket.Add(v)
		if verbose {
			fmt.Fprintf(os.Stdout, "Added: %s\n", v)
		}
	}
}

func readStdin() (input []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = strings.Split(scanner.Text(), " ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return
}
