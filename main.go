package main

import (
	"fmt"
	"os/exec"
)

func main() {
	cmd := exec.Command("cmd", "/c", "shutdown", "-r")

	if err := cmd.Start(); err != nil {
		fmt.Println("Error:", err)
	}
}
