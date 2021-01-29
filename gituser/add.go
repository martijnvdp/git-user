package gituser

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add github user",
	Long:  `add github user`,
	Run: func(cmd *cobra.Command, args []string) {
		adduser()
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func adduser() {
	fmt.Println("add user")
	fmt.Println("Enter username: ")
	var username string
	_, err := fmt.Scanln(&username)

	if err != nil {
		error.Error(err)
	}
}
