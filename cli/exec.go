package cli

import "context"

type ExecCMD struct {
	List *ListExecCMD `arg:"subcommand:list"`
}

func (e *ExecCMD) handle(ctx context.Context) error {

	p := getParserFromContext(ctx)

	switch {
	case e.List != nil:
		err := e.List.handle(ctx)
		if err != nil {
			p.FailSubcommand(err.Error(), "exec", "list")
		}
	}

	return nil
}

type ListExecCMD struct {
}

func (l *ListExecCMD) handle(ctx context.Context) error {

	client := getClientFromContext(ctx)

	execs, err := client.GetExecList()
	if err != nil {
		return err
	}

	header := []string{"ID", "ENV", "SCRIPT"}

	data := [][]string{}

	for _, v := range execs {
		data = append(data, []string{v.ID, v.ExecEnv.Name, v.Script.Name})
	}

	showTable(header, data)

	return nil
}
