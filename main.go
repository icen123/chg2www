package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args

	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))

	needChangePath := dir + "/"
	if len(args) >= 2 {
		if strings.HasPrefix(args[1], "/") {
			needChangePath = args[1]
		} else  {
			needChangePath += args[1]
		}
	}

	command := "chown www:www " + needChangePath + " -R"
	cmd := exec.Command("/bin/bash", "-c", command)

	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	command = "chmod g+w " + needChangePath + " -R"
	cmd = exec.Command("/bin/bash", "-c", command)
	_, err = cmd.Output()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}
