package cli

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

type Store struct {
	Token string `json:"token"`
}

type WorkspaceDetails struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Host      string    `json:"host"`
	Token     string    `json:"token"`
	ExpiresAt time.Time `json:"expires_at"`
}

type Config struct {
	Workspace  string                      `json:"workspace"` //selected working host
	Workspaces map[string]WorkspaceDetails `json:"workspaces"`
}

func getConfig() Config {
	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "config.json")

	cfg := Config{
		Workspaces: map[string]WorkspaceDetails{},
	}

	if _, err := os.Stat(storeFile); err == nil {
		bconfig, _ := os.ReadFile(storeFile)
		_ = json.Unmarshal(bconfig, &cfg)
	}
	return cfg
}
