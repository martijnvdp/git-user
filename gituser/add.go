package gituser

import (
	"fmt"

	"github.com/spf13/viper"
)

type gituserdata struct {
	userName  string `mapstructure:"user_name"`
	userEmail string `mapstructure:"user_email"`
	userToken string `mapstructure:"user_token"`
}

type gitusers struct {
	users      gituserdata
	configfile string `mapstructure:"config_file"`
}

func Adduser() {
	var userdata gituserdata
	fmt.Println("add user")
	fmt.Println("Enter username: ")
	_, err := fmt.Scanln(&userdata.userName)
	fmt.Println("Enter username: ")
	_, err = fmt.Scanln(&userdata.userEmail)
	fmt.Println("Enter username: ")
	_, err = fmt.Scanln(&userdata.userToken)

	if userdata.userName != "" {
		viper.SetDefault("users.user_name", userdata.userName)
		viper.SetDefault("users.user_email", userdata.userEmail)
		viper.SetDefault("users.user_token", userdata.userToken)
		viper.WriteConfig()
	}

	if err != nil {
		error.Error(err)
	}
}
