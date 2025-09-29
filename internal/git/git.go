package git

import (
	"fmt"
	"os/exec"
)

func IsGitInstalled() error {
	_, err := exec.LookPath("git")
	if err != nil {
		return fmt.Errorf("git command not found in PATH: please install Git and ensure it's in your PATH")
	}
	return nil
}

func HasGitRepo() error {
	cmd := exec.Command("git", "rev-parse", "--git-dir")
	err := cmd.Run()

	if err != nil {
		return err
	}

	return nil
}
