package main

import (
	"fmt"
	"os"

	"github.com/sandergv/scriptctl/cli"
	yaml "gopkg.in/yaml.v3"
)

func main() {
	fn := os.Args[1]

	fmt.Println(fn)

	b, err := os.ReadFile(fn)
	fmt.Println(err)

	t := cli.ConfigFile{}

	_ = yaml.Unmarshal(b, &t)
	fmt.Println(t)

}
