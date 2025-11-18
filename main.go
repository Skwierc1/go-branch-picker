package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/manifoldco/promptui"
)

func main() {
	allOptions := getBranches()
	prompt := promptui.Select{
		Label: "Select Branch or Action",
		Items: allOptions,
		Size:  10,
		Searcher: func(input string, index int) bool {
			return branchSearcher(allOptions, input, index)
		},
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Failed to select branch or action %v\n", err)
		return
	}

	if result == "Exit" {
		fmt.Println("Exiting...")
		return
	}

	if result == "Create New Branch" {
		fmt.Println("Creating new branch...")
		createNewBranch()
		return
	}

	fmt.Printf("Checking out %s\n", result)
	cmd := exec.Command("git", "checkout", result)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed to checkout branch '%s': %v\n", result, err)
		if len(output) > 0 {
			fmt.Fprintf(os.Stderr, "%s\n", string(output))
		}
		os.Exit(1)
	}
	if len(output) > 0 {
		fmt.Printf("%s\n", string(output))
	}
}

func createNewBranch() {
	prompt := promptui.Prompt{
		Label: "Enter New Branch Name",
	}
	name, err := prompt.Run()
	if err != nil {
		fmt.Printf("Failed to create new branch %v\n", err)
		return
	}

	exec.Command("git", "checkout", "-b", name).Run()
}

func getBranches() []string {
	// execute git branch --list and return the list of branches
	branches, err := exec.Command("git", "branch", "--list").Output()
	if err != nil {
		fmt.Printf("Failed to get branches %v\n", err)
		return []string{}
	}

	branchLines := strings.Split(strings.TrimSpace(string(branches)), "\n")
	allBranches := make([]string, 0, len(branchLines))
	for _, line := range branchLines {
		// Trim whitespace and remove asterisk prefix (for current branch)
		branchName := strings.TrimSpace(line)
		branchName = strings.TrimPrefix(branchName, "* ")
		if branchName != "" {
			allBranches = append(allBranches, branchName)
		}
	}

	defaultOptions := []string{"Exit", "Create New Branch"}
	selectOptions := make([]string, len(allBranches)+len(defaultOptions))
	copy(selectOptions, defaultOptions)
	copy(selectOptions[len(defaultOptions):], allBranches)
	return selectOptions
}

func branchSearcher(allOptions []string, input string, index int) bool {
	branch := allOptions[index]
	name := strings.Replace(strings.ToLower(branch), " ", "", -1)
	input = strings.Replace(strings.ToLower(input), " ", "", -1)

	return strings.Contains(name, input)
}
