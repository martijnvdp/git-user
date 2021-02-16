package gituser

import (
	"fmt"
	"os/exec"
)

func Status() {

	out, err := exec.Command("git", "config", "--list").Output()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("Current git config:")
	fmt.Println(string(out))
}
