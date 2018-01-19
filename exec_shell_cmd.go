package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := "date; ls -al /tmp/"
	out, err := exec.Command("bash", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Failed to execute command: %s\n", cmd)
	}
	fmt.Printf("%s", out)
}
