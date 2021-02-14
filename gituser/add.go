package gituser

import (
	"fmt"

	"github.com/spf13/viper"
)

type Userdata struct {
	Name  string `mapstructure:"name"`
	Email string `mapstructure:"email"`
	Token string `mapstructure:"token"`
}

type Gitusers struct {
	Users      []Userdata `mapstructure:"users"`
	Configfile string     `mapstructure:"config_file"`
}

func Writeuser(usr Gitusers) {
	for _, u := range usr.Users {

		if u.Name != "" {
			viper.Set("users.user_name", u.Name)
			viper.Set("users.user_email", u.Email)
			viper.Set("users.user_token", u.Token)
			viper.WriteConfig()
		}

	}
}

func Adduser() {
	var git_users Gitusers
	var user Userdata
	fmt.Println("add user")
	fmt.Println("Enter username: ")
	_, err := fmt.Scanln(&user.Name)
	fmt.Println("Enter e-mail: ")
	_, err = fmt.Scanln(&user.Email)
	fmt.Println("Enter token: ")
	_, err = fmt.Scanln(&user.Token)

	git_users.Users = append(git_users.Users, user)
	Writeuser(git_users)

	if err != nil {
		error.Error(err)
	}
}
