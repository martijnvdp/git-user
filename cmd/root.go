/*
Copyright © 2021 M van der Ploeg

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/martijnxd/git-user/gituser"
	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-user",
	Short: "git-user a cli util to switch between git user profiles",
	Long: `git-user a cli util to switch between git user profiles
 For example:
 add users:
 git-user add
 
 list users:
 git-user -l

 switch user:
 git-user -u username
 git-user --user username
 
 show current git config:
 git-user status

 test current git config:
 git-user test
 
 `,
	Run: func(cmd *cobra.Command, args []string) {
		l, _ := cmd.Flags().GetBool("list")
		u, _ := cmd.Flags().GetString("user")
		if u != "" {
			gituser.Setuser(u)
		}
		if l {
			gituser.Listusers()
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	var l bool
	var u string
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-user.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&l, "list", "l", false, "list users from the config file.")
	rootCmd.Flags().StringVarP(&u, "user", "u", "", "user profile")

}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var git_users gituser.Gitusers
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		configName := "config.gituser"
		configType := "yml"
		viper.AddConfigPath(home)
		viper.SetConfigName(configName)
		viper.SetConfigType(configType)
		cfgFile = filepath.Join(home, configName+"."+configType)
	}
	viper.AllowEmptyEnv(true)
	viper.Unmarshal(&git_users)
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	} else {
		_, err := os.Stat(cfgFile)
		if !os.IsExist(err) {
			if _, err := os.Create(cfgFile); err != nil {
			}
		}
		if err := viper.SafeWriteConfig(); err != nil {
		}
	}
}
