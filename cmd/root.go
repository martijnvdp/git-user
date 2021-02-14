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

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

type gituserdata struct {
	userName  string `mapstructure:"user_name"`
	userEmail string `mapstructure:"user_email"`
	userToken string `mapstructure:"user_token"`
}

type gitusers struct {
	users      gituserdata
	configfile string `mapstructure:"config_file"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "git-user",
	Short: "git-user a cli util to switch between git user profiles",
	Long: `git-user a cli util to switch between git user profiles
 For example:

git-user username`,
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
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.git-user.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	var git_users gitusers
	var git_userdata gituserdata
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
	viper.SetEnvPrefix("GITUSER_") //not working
	viper.AllowEmptyEnv(true)
	viper.AutomaticEnv()
	viper.Unmarshal(&git_users)
	viper.Unmarshal(&git_userdata)
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
