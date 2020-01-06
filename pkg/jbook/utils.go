package jbook

import "os/exec"

func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

// CheckTools checks that the required executables ("jrnl", "mdbook") are avaliable and on the path.
func CheckTools() bool {
	return commandExists("jrnl") && commandExists("mdbook")
}
