package gituser

import (
	"fmt"
	"os/exec"
)

func unsetuser() {
	cmd := exec.Command("git", "config", "--unset", "user.name")
	_, err := cmd.Output()
	cmd = exec.Command("git", "config", "--unset", "user.email")
	_, err = cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}
