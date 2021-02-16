package gituser

import (
	"fmt"
	"os/exec"
)

func setuser(username string) {
	git_users := Getusers()
	for _, usr := range git_users.Users {
		if usr.Name == username {
			cmd := exec.Command("git", "config", "--Global", "user.name", usr.Name)
			_, err := cmd.Output()
			cmd = exec.Command("git", "config", "--Global", "user.email", usr.Email)
			_, err = cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}

		}
	}
}
