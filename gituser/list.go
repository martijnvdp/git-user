package gituser

import (
	"fmt"

	"github.com/spf13/viper"
)

func Listusers() {
	var users []Userdata
	viper.UnmarshalKey("users.Users", &users)
	fmt.Println("list users")
	for _, usr := range users {
		fmt.Println(" account: ", usr.Name)
		fmt.Println(" mail: ", usr.Email)
		fmt.Println(" key_file: ", usr.Keyfilename)
	}
}
