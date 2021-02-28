package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/adrg/xdg"
)

var consumerKey string

var rootCmd = &cobra.Command{
	Use:   "gocket",
	Short: "Pocket in the shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Pocket consumer key (required).")
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))
	rootCmd.AddCommand(addCmd)
}

func Execute() {
	verifyKey()
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func verifyKey() {
	if viper.Get("key") == "" {
		os.Stderr.WriteString(fmt.Sprintf(`
ERROR: You need a pocket consumer key.
You can create an application with a key at: https://getpocket.com/developer/apps/
You can then use the option -k to specify the key.
You can also write "key: 123_consumer_key" in the file "%s".`,
			filepath.Join(xdg.ConfigHome, "gocket/config.yml"),
		))
		rootCmd.Help()
		os.Exit(1)
	}
}
