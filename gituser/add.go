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
	Users []Userdata `mapstructure:"users"`
}

func adduser() *Userdata {
	var user Userdata
	fmt.Println("add user")
	fmt.Println("Enter username: ")
	_, err := fmt.Scanln(&user.Name)
	fmt.Println("Enter e-mail: ")
	_, err = fmt.Scanln(&user.Email)
	fmt.Println("Enter token: ")
	_, err = fmt.Scanln(&user.Token)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return &user
}

func Adduser() {
	var git_users Gitusers
	var err error
	var input string
	for addanother := true; addanother != false; {
		git_users.Users = append(git_users.Users, *adduser())
		println("Add another (yes/no)")
		fmt.Scan(&input)
		if input != "yes" && input != "y" {
			addanother = false
		}
	}
	viper.Set("users", git_users)
	viper.WriteConfig()

	if err != nil {
		error.Error(err)
	}
}
