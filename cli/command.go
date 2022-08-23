package cli

import (
	"context"
	"fmt"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type CommandCMD struct {
	Create *CreateCommandCMD `arg:"subcommand:create"`
	List   *ListCommandCMD   `arg:"subcommand:list"`
}

func (c *CommandCMD) handle(ctx context.Context) error {

	p := getParserFromContext(ctx)

	switch {
	case c.Create != nil:
		err := c.Create.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "command", "create")
		}
	case c.List != nil:
		err := c.List.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "command", "list")
		}
	}
	return nil
}

type CreateCommandCMD struct {
	Name        string   `arg:"positional"` // unique
	Script      string   `arg:"positional"`
	Description string   `arg:"-d"`
	Namespace   string   `arg:"-n"`
	Env         []string `arg:"-e"`
}

func (c *CreateCommandCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	id, err := client.CreateCommand(types.CreateCommandRequest{
		Name:        c.Name,
		Description: c.Description,
		Namespace:   c.Namespace,
		Env:         c.Env,
		ScriptID:    c.Script,
	})
	if err != nil {
		return err
	}

	fmt.Println("Command ID:", id)

	return nil
}

type ListCommandCMD struct {
}

func (l *ListCommandCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	commands, err := client.GetCommandList()
	if err != nil {
		return err
	}

	header := []string{"ID", "NAME", "DESCRIPTION", "SCRIPT", "NAMESPACE"}

	data := [][]string{}

	for _, d := range commands {
		data = append(data, []string{d.ID, d.Name, d.Description, d.Script.Name, d.Namespace})
	}

	showTable(header, data)

	return nil
}
