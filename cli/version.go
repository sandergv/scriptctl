package cli

import "fmt"

type VersionCMD struct {
}

func (v *VersionCMD) handle() {
	fmt.Println("scriptctl v0.0.1")
}
