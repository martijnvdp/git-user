package gituser

import (
	"fmt"
	"os/exec"
)

func setuser() {
	fmt.Println("Enter username: ")
	var username string
	var email string
	_, err := fmt.Scanln(&username)
	fmt.Println("Enter email: ")
	_, err = fmt.Scanln(&email)

	if err != nil {
		error.Error(err)
	} else {
		cmd := exec.Command("git", "config", "--local", "user.name", username)
		_, err := cmd.Output()
		cmd = exec.Command("git", "config", "--local", "user.email", email)
		_, err = cmd.Output()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
