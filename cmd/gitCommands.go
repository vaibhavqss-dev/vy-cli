package cmd

import (
	"bytes"
	"fmt"
	"os/exec"
)

func Commit(message string) string {
	gitInit()
	// check if all the changes committed or not!
	cmd := exec.Command("git", "add", ".")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Sprintf("commit error: %w\nOutput: %s\nError: %s",
		err, stdout.String(), stderr.String())
	}
	
	cmd = exec.Command("git", "commit", "-m", message)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return fmt.Sprintf("commit error: %w\nOutput: %s\nError: %s",
		err, stdout.String(), stderr.String())
	}
	return "changes has been Staged and Committed"
}

// this will initialize the git repository
func gitInit() {
	_ = exec.Command("git", "init")
}