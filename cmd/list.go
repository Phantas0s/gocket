package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Phantas0s/gocket/internal"
	"github.com/Phantas0s/gocket/internal/platform"
	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var order, search string
var count int
var tui, archive, delete, noconfirm, title bool

func init() {
	listCmd.Flags().StringVarP(
		&order,
		"order",
		"o",
		"newest",
		"order by 'newest' (default), 'oldest', 'title', or 'url'.",
	)
	listCmd.Flags().StringVarP(&search, "search", "s", "", "Only list items with title or URL matching the search.")
	listCmd.Flags().IntVarP(&count, "count", "c", 10, "Number of results (0 for all, default 10).")

	listCmd.Flags().BoolVarP(&tui, "tui", "", false, "Display the TUI.")
	listCmd.Flags().BoolVarP(&title, "title", "t", false, "Display the title the line before the URL.")
	listCmd.Flags().BoolVarP(&archive, "archive", "a", false, "Archive the listed articles.")
	listCmd.Flags().BoolVarP(&delete, "delete", "d", false, "Delete the listed articles.")
	listCmd.Flags().BoolVarP(&noconfirm, "noconfirm", "", false, "Don't ask for any confirmation.")
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
	if viper.Get("key") == nil {
		fmt.Println("ERROR: You need a pocket consumer key.")
		fmt.Println("You can create an application with a key at: https://getpocket.com/developer/apps/")
		fmt.Sprintf(
			"Use the key flag -k to specify the key or write it in the file %s",
			// TODO define the config path at one place (DRY)
			filepath.Join(xdg.ConfigHome, "gocket"),
		)
		rootCmd.Help()
	}

	pocket := internal.CreatePocket(viper.Get("key").(string))
	list := pocket.List(count, order, search)
	if tui {
		tui := internal.TUI{Instance: &platform.Tview{}, Pocket: pocket}
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
			if noconfirm || prompt("Do you really want to archive all the articles listed?") {
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
