package gituser

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "git-user",
	Short: "switch between github accounts",
	Long: `CLI tool to switch between github accounts

	Usage:
	  - list
	  	list github accounts
	  - add
		add github account
	  - remove
		revoce github account
	  - set
	    set active github account `,
	Run: func(cmd *cobra.Command, args []string) {
		a, err := cmd.Flags().GetBool("add")
		l, err := cmd.Flags().GetBool("list")
		s, err := cmd.Flags().GetBool("set")
		if a && err == nil {
			adduser()
		}
		if l && err == nil {
			listusers()
		}
		if s && err == nil {
			setuser()
		}
		fmt.Println("use one of the command line arguments")

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.my-calc.yaml)")
	rootCmd.Flags().BoolP("list", "l", false, "List github users")
	rootCmd.Flags().BoolP("add", "a", false, "add github user")
	rootCmd.Flags().BoolP("set", "s", false, "set active github user")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(home)
		viper.SetConfigName(".git-users")
	}

	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
