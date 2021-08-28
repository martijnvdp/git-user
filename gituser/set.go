package gituser

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/viper"
)

func Copyfile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func Setuser(username string) {
	var users []Userdata
	viper.UnmarshalKey("users.Users", &users)
	for _, usr := range users {
		if usr.Name == username {
			_, err := exec.Command("git", "config", "--local", "user.name", usr.Name).Output()
			_, err = exec.Command("git", "config", "--local", "user.email", usr.Email).Output()
			_, err = exec.Command("git", "config", "--global", "user.name", usr.Name).Output()
			_, err = exec.Command("git", "config", "--global", "user.email", usr.Email).Output()

			src := os.Getenv("userprofile") + "\\.ssh\\" + usr.Keyfilename
			dest := os.Getenv("userprofile") + "\\.ssh\\" + "id_rsa"
			e := Copyfile(src, dest)
			if e != nil {
				log.Fatal(e)
			}
			if err == nil {
				fmt.Println("switched to user: ", username)
			} else {
				fmt.Println(err.Error())
				return
			}
		}
	}
}
