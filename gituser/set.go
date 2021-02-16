package gituser

import (
	"fmt"
	"os/exec"

	"github.com/spf13/viper"
)

func Setuser(username string) {
	var users []Userdata
	viper.UnmarshalKey("users.Users", &users)
	for _, usr := range users {
		if usr.Name == username {
			cmd, err := exec.Command("git", "config", "--local", "user.name", usr.Name).Output()
			fmt.Println(string(cmd))
			cmd, err = exec.Command("git", "config", "--local", "user.email", usr.Email).Output()
			fmt.Println(string(cmd))
			cmd, err = exec.Command("git", "config", "--global", "user.name", usr.Name).Output()
			fmt.Println(string(cmd))
			cmd, err = exec.Command("git", "config", "--global", "user.email", usr.Email).Output()
			fmt.Println(string(cmd))
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		}
	}
}
