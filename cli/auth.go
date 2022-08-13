package cli

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"syscall"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
	"golang.org/x/term"
)

type AuthCMD struct {
	Login *LoginCMD `arg:"subcommand:login"`
}

func (a *AuthCMD) handle(client *scriptlabctl.Client) {
	switch {
	case a.Login != nil:
		a.Login.handle(client)

	}
}

type LoginCMD struct {
}

func (l *LoginCMD) handle(client *scriptlabctl.Client) {
	// fmt.Print("Username: ")
	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "store.json")
	reader := bufio.NewReader(os.Stdin)

	username := ""
	password := ""

	for {
		fmt.Print("Username: ")
		username, _ = reader.ReadString('\n')
		username = strings.Trim(username, " \n")
		if username == "" {
			fmt.Println("Username can't be empty")
			continue
		}
		if strings.Contains(username, " ") {
			fmt.Println("Username can't contain spaces")
			continue
		}
		break
	}

	//
	fmt.Print("Password: ")
	bytePassword, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		fmt.Println(err)
		return
	}
	password = string(bytePassword)
	_ = password

	// token, _, err :=
	fmt.Println()
	token, _, err := client.Login(username, password)
	if err != nil {
		fmt.Println(err)
	}
	store := Store{
		Token: token,
	}

	bstore, err := json.Marshal(store)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(storeFile, bstore, os.ModePerm)
	if err != nil {
		panic(err)
	}
	fmt.Println("Login Succeeded")
}
