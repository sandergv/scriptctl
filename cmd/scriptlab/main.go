package main

import (
	"encoding/json"
	"os"
	"path"

	"github.com/sandergv/scriptlab/cli"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
)

func main() {

	// check if folder exist
	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "config.json")

	cfg := cli.Config{
		Workspaces: map[string]cli.WorkspaceDetails{},
	}

	if _, err := os.Stat(storeFile); err != nil {
		// fmt.Println(err)
		_ = os.Mkdir(dir, os.ModePerm)
		b, _ := json.Marshal(cfg)
		_ = os.WriteFile(storeFile, b, os.ModePerm)
	} else {
		bconfig, _ := os.ReadFile(storeFile)
		_ = json.Unmarshal(bconfig, &cfg)
	}

	url := ""
	token := ""
	if a, ok := cfg.Workspaces[cfg.Workspace]; ok {
		url = a.Host
		token = a.Token
	}

	client := scriptlabctl.NewClient(scriptlabctl.ClientOptions{
		Url:   url,
		Token: token,
	})
	cli.Exec(client)
}
