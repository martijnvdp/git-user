# git-user
git switch user tool
 testing & learning cli

switches git user , using git config
and copies the specified key file to id_rsa file in the .ssh profile

## Usage ##
```
git-user a cli util to switch between git user profiles
 For example:
 add users:
 git-user add
 
 list users:
 git-user -l

 switch user:
 git-user -u username
 git-user --user username
 
 show current git config:
 git-user status

Usage:
  git-user [flags]
  git-user [command]

Available Commands:
  add         add users to the git-user config
  help        Help about any command
  list        List users
  status      Show current git config

Flags:
      --config string   config file (default is $HOME/.git-user.yaml)
  -h, --help            help for git-user
  -l, --list            list users from the config file.
  -t, --toggle          Help message for toggle
  -u, --user string     user profile

Use "git-user [command] --help" for more information about a command.

```
## refs ##
```
https://github.com/spf13/cobra
https://github.com/spf13/viper
```
