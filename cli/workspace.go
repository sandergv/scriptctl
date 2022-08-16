package cli

import (
	"bufio"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"regexp"
	"strings"
	"syscall"

	"golang.org/x/term"
)

type WorkspaceCMD struct {
	Login   *LoginWorkspaceCMD   `arg:"subcommand:login"`
	List    *ListWorkspaceCMD    `arg:"subcommand:list"`
	Use     *UseWorkspaceCMD     `arg:"subcommand:use"`
	Inspect *InspectWorkspaceCMD `arg:"subcommand:inspect"`
}

func (w *WorkspaceCMD) handle(ctx context.Context) error {

	p := getParserFromContext(ctx)

	switch {
	case w.Login != nil:
		err := w.Login.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "workspace", "login")
		}
	case w.List != nil:
		err := w.List.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "workspace", "list")
		}
	case w.Use != nil:
		err := w.Use.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "workspace", "use")
		}
	case w.Inspect != nil:
		err := w.Inspect.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "workspace", "inspect")
		}
	}

	return nil
}

type LoginWorkspaceCMD struct {
	Username  string `arg:"-u" help:"Username configured for the workspace"`
	Password  string `arg:"-p" help:"Password configured for the workspace"`
	PassStdin bool   `arg:"--password-stdin" help:"Take password from stdin"`
	Name      string `arg:"positional,required"`
	// Description string `arg`
	Server string `arg:"positional" help:"Server url, if no url is provided it will try to conect to http://localhost:6892"` // host url
}

func (l *LoginWorkspaceCMD) handle(ctx context.Context) error {

	if l.Server == "" {
		l.Server = "http://localhost:6892"
	}

	valid, _ := regexp.MatchString("^[a-z-_]{3,12}", l.Name)
	if !valid {
		return errors.New("invalid workspace name")
	}

	// check if folder exist
	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "config.json")

	cfg := Config{
		Workspaces: map[string]WorkspaceDetails{},
	}

	if _, err := os.Stat(storeFile); os.IsNotExist(err) {

		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			return err
		}
	} else {
		bconfig, _ := os.ReadFile(storeFile)
		_ = json.Unmarshal(bconfig, &cfg)
	}

	if _, ok := cfg.Workspaces[l.Name]; ok {
		return errors.New("name is already registered")
	}

	client := getClientFromContext(ctx)

	reader := bufio.NewReader(os.Stdin)

	uFlag := false
	pFlag := false

	for _, f := range os.Args {
		if f == "-u" || f == "--username" {
			uFlag = true
		} else if f == "-p" || f == "--password" {
			pFlag = true
		}
	}

	if !uFlag {
		fmt.Print("Username: ")
		username, _ := reader.ReadString('\n')
		l.Username = strings.Trim(username, " \n")
	}
	if !pFlag || l.PassStdin {
		fmt.Print("Password: ")
		bytePassword, err := term.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println(err)
			return err
		}
		l.Password = string(bytePassword)

	}

	// token, _, err :=
	fmt.Println()
	details, err := client.Login(l.Server, l.Username, l.Password)
	if err != nil {
		fmt.Println(err)
		return err
	}

	cfg.Workspace = l.Name
	cfg.Workspaces[l.Name] = WorkspaceDetails{
		ID:        details.WorkspaceID,
		Name:      l.Name,
		Host:      l.Server,
		Token:     details.Token,
		ExpiresAt: details.ExpiresAt,
	}

	b, _ := json.Marshal(cfg)
	_ = os.WriteFile(storeFile, b, os.ModePerm)

	fmt.Println("Login Succeeded")

	return nil
}

type ListWorkspaceCMD struct {
}

func (l *ListWorkspaceCMD) handle(ctx context.Context) error {

	// check if folder exist
	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "config.json")

	cfg := Config{
		Workspaces: map[string]WorkspaceDetails{},
	}

	if _, err := os.Stat(storeFile); err == nil {
		bconfig, _ := os.ReadFile(storeFile)
		_ = json.Unmarshal(bconfig, &cfg)
	}

	header := []string{"NAME", "ENDPOINT"}
	data := [][]string{}

	for _, w := range cfg.Workspaces {
		ws := w.Name
		if cfg.Workspace == w.Name {
			ws = w.Name + " *"
		}
		data = append(data, []string{ws, w.Host})
	}

	showTable(header, data)

	return nil
}

type UseWorkspaceCMD struct {
	Name string `arg:"positional,required"`
}

func (u *UseWorkspaceCMD) handle(ctx context.Context) error {

	cfg := getConfig()

	if cfg.Workspace == u.Name {
		return errors.New("namespace already in use")
	}

	if _, ok := cfg.Workspaces[u.Name]; !ok {
		return errors.New("workspace does not exist")
	}

	//
	cfg.Workspace = u.Name

	dir := path.Join(os.Getenv("HOME"), ".scriptlab")
	storeFile := path.Join(dir, "config.json")
	b, _ := json.Marshal(cfg)
	_ = os.WriteFile(storeFile, b, os.ModePerm)

	return nil
}

type InspectWorkspaceCMD struct {
	Name string `arg:"positional,required"`
}

func (i *InspectWorkspaceCMD) handle(ctx context.Context) error {

	cfg := getConfig()

	b, _ := json.MarshalIndent(cfg.Workspaces[i.Name], "", "  ")

	fmt.Println(string(b))

	return nil
}
