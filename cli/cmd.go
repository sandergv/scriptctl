package cli

import (
	"context"
	"fmt"
)

type CmdCMD struct {
	Name string   `arg:"positional,required" help:""`
	Args []string `arg:"positional" help:""`
}

func (c *CmdCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	details, err := client.RunCommand(c.Name, c.Args)
	if err != nil {
		fmt.Println(err)
		return err
	}

	for _, o := range details.Output {
		fmt.Println(o)
	}

	if details.ExitCode != 0 {
		fmt.Println(details.Error)
	}

	return nil
}
