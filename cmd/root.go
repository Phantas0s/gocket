package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/adrg/xdg"
)

var consumerKey string

func rootCmd(v *viper.Viper) *cobra.Command {
	return &cobra.Command{
		Use:   "gocket",
		Short: "Pocket in the shell",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			bindFlagToConfig(cmd, v)
			verifyKey(cmd)
		},
	}
}

func initConfig() *viper.Viper {
	v := viper.New()
	v.AddConfigPath(filepath.Join(xdg.ConfigHome, "gocket"))
	v.AddConfigPath(".")
	v.SetConfigName("config")

	v.AutomaticEnv()
	v.ReadInConfig()

	return v
}

func Execute() {
	rootCmd := rootCmd(initConfig())
	rootCmd.AddCommand(ListCmd())
	rootCmd.PersistentFlags().StringVarP(&consumerKey, "key", "k", "", "Pocket consumer key (required).")
	rootCmd.AddCommand(addCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func verifyKey(cmd *cobra.Command) {
	if consumerKey == "" {
		os.Stderr.WriteString(fmt.Sprintf(`
ERROR: You need a pocket consumer key.
You can create an application with a key at: https://getpocket.com/developer/apps/
You can then use the option -k to specify the key.
You can also write "key: 123_consumer_key" in the file "%s".`,
			filepath.Join(xdg.ConfigHome, "gocket/config.yml"),
		))
		cmd.Help()
		os.Exit(1)
	}
}

func bindFlagToConfig(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})
}
