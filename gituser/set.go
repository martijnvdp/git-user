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
			cmd := exec.Command("git", "config", "--local", "user.name", usr.Name)
			_, err := cmd.Output()
			cmd = exec.Command("git", "config", "--local", "user.email", usr.Email)
			_, err = cmd.Output()
			cmd = exec.Command("git", "config", "--global", "user.name", usr.Name)
			_, err = cmd.Output()
			cmd = exec.Command("git", "config", "--global", "user.email", usr.Email)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		}
	}
}
