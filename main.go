package main

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	prompt := promptui.Select{
		Label: "Select Branch",
		Items: getBranches(),
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You choose %q\n", result)
}

func getBranches() []string {
	// execute git branch --list and return the list of branches
	branches, err := exec.Command("git", "branch", "--list").Output()
	if err != nil {
		fmt.Printf("Failed to get branches %v\n", err)
		return []string{}
	}
	return strings.Split(strings.TrimSpace(string(branches)), "\n")
}
