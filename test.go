package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	p := filepath.Base("test/test.py")
	fmt.Println(p)
}
