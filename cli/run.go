package cli

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/sandergv/scriptlab/pkg/scriptlabctl"
	"github.com/sandergv/scriptlab/pkg/scriptlabctl/types"
)

type RunCMD struct {
	File string   `arg:"-f"`
	Env  []string `arg:"-e"`
	Exec string

	Args []string `arg:"-a"`

	Out      bool `arg:"--show-output"`
	Logs     bool `arg:"--show-logs"`
	Response bool `arg:"--show-response"`

	ResultPath string `arg:"--result-path" placeholder:"PATH"`
}

func (r *RunCMD) handle(client *scriptlabctl.Client) {
	// fmt.Println(r.File)
	// fmt.Println(r.ResultPath)
	if r.File != "" {
		details, err := runCode(client, r.File, r.Env, r.Args)
		if err != nil {
			fmt.Println(err)
		}
		showDetails(details, r.Out, r.Logs, r.Response)
		if r.ResultPath != "" {
			storeResult(r.ResultPath, details)
		}
	}

}

func runCode(c *scriptlabctl.Client, fp string, env []string, args []string) (types.RunDetails, error) {
	bcontent, err := ioutil.ReadFile(fp)
	if err != nil {
		fmt.Println("aca", err)
	}
	details, err := c.RunCode(types.RunCodeOptions{
		Type: "python",
		Env:  env,
		Args: env,
		Code: string(bcontent),
	})
	return details, err
}

func showDetails(details types.RunDetails, out bool, logs bool, response bool) {
	fmt.Println("Exit Code:", details.ExitCode)
	if details.ExitCode != 0 {
		fmt.Println("err", details.Error)
		return
	}

	if out {
		if len(details.Output) > 0 {
			fmt.Println("Output:")
			for _, v := range details.Output {
				fmt.Println(v)
			}
		} else {
			fmt.Println("no output")
		}
	}

	if logs {
		fmt.Println("Logs:")
		for _, v := range details.Logs {
			fmt.Println(v)
		}
	}

	if response {
		fmt.Println("Response:")
		fmt.Println(details.Response)
	}
}

func storeResult(rp string, details types.RunDetails) {
	if _, err := os.Stat(rp); err != nil {
		fmt.Println(err)
		return
	}

	b, err := json.MarshalIndent(details, "", "  ")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ioutil.WriteFile(path.Join(rp, "result.json"), b, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
}
