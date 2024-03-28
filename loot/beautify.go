package loot

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

//requires jsbeautifier
//pip install jsbeautifier

const (
	beautifyCommand = "js-beautify"
)

func beautfyExists() bool {
	_, err := exec.LookPath(beautifyCommand)
	return err == nil
}

func beautifyFile(file string) error {
	doIfVerbose(func() {
		fmt.Printf("beautifying %s\n", filepath.Base(file))
	})
	args := []string{"-r", file}
	cmd := exec.Command(beautifyCommand, args...)
	err := cmd.Run()
	if err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return fmt.Errorf("Command '%s' failed with exit status %d\n", beautifyCommand, exiterr.ExitCode())
		}
		return err
	}
	doIfVerbose(func() {
		fmt.Printf("%s is beautiful\n", filepath.Base(file))
	})
	return nil
}
