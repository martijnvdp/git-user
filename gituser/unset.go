package gituser

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "unset active github user",
	Long:  `unset active user`,
	Run: func(cmd *cobra.Command, args []string) {
		unsetuser()
	},
}

func init() {
	rootCmd.AddCommand(unsetCmd)
}

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
