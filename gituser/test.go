package gituser

import (
	"fmt"
	"os/exec"
)

func Test() {
	exit := ""
	_, err := exec.Command("ssh", "-T", "git@github.com").Output()
	if err != nil {
		exit = err.Error()
	}
	fmt.Println("Testing connection with github:")
	if exit == "exit status 1" {
		fmt.Println("succes")
	} else {
		fmt.Println(err.Error())
	}
}
