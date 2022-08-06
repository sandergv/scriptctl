package cli

import (
	"fmt"
	"io/ioutil"

	"github.com/sandergv/scriptctl/pkg/scriptlabctl"
	"github.com/sandergv/scriptctl/pkg/scriptlabctl/types"
)

type RunCMD struct {
	File string   `arg:"-f"`
	Env  []string `arg:"-e"`
	Args []string `arg:"-a"`
}

func (r *RunCMD) handle(client *scriptlabctl.Client) {
	fmt.Println(r.File)
	bcontent, err := ioutil.ReadFile(r.File)
	if err != nil {
		fmt.Println("aca", err)
	}
	details, err := client.RunCode(types.RunCodeOptions{
		Type: "python",
		Env:  r.Env,
		Args: r.Args,
		Code: string(bcontent),
	})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(details)
}
