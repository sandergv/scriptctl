package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/sandergv/scriptlab/cli"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
)

func main() {

	// check if folder exist
	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "store.json")
	if _, err := os.Stat(dir); err != nil {
		err := os.Mkdir(dir, os.ModePerm)
		if err != nil {
			fmt.Println("aca", err)
		}
	}
	// if _, err := os.Stat(store); err != nil {
	// 	fmt.Println(err)
	// 	fmt.Println("user is not logged in")
	// }
	store := cli.Store{}

	bstore, _ := os.ReadFile(storeFile)
	_ = json.Unmarshal(bstore, &store)

	client := scriptlabctl.NewClient(store.Token)
	cli.Exec(client)
}
