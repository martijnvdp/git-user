package gituser

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list github users",
	Long:  `list github users`,
	Run: func(cmd *cobra.Command, args []string) {
		listusers()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listusers() {
	fmt.Println("list users")
}
