package cli

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type ActionCMD struct {
	Create *CreateActionCMD `arg:"subcommand:create"`
	List   *ListActionCMD   `arg:"subcommand:list"`
}

func (a *ActionCMD) handle(ctx context.Context) error {

	p := getParserFromContext(ctx)

	switch {
	case a.Create != nil:
		err := a.Create.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "action", "create")
		}
	case a.List != nil:
		err := a.List.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "action", "list")
		}
	}

	return nil
}

type CreateActionCMD struct {
	Name string `arg:"positional,required"`
	Exec string `arg:"positional,required"`
}

func (c *CreateActionCMD) handle(ctx context.Context) error {

	valid, _ := regexp.MatchString("^[a-z-_]{3,12}", c.Name)

	if !valid {
		return errors.New("invalid action name")
	}

	client := getClientFromContext(ctx)

	id, err := client.CreateAction(types.CreateActionRequest{
		Name:   c.Name,
		ExecID: c.Exec,
	})
	if err != nil {
		return err
	}
	fmt.Println("Action ID:", id)

	return nil

}

type ListActionCMD struct {
}

func (l *ListActionCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	acts, err := client.GetActionList()
	if err != nil {
		return err
	}

	headers := []string{"ID", "NAME", "SCRIPT", "EXEC ID"}

	data := [][]string{}

	for _, a := range acts {
		data = append(data, []string{a.ID, a.Name, a.Script.Name, a.ExecID})
	}

	showTable(headers, data)

	return nil
}
