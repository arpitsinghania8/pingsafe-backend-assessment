package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func scanFileForKeys(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error opening file: %s", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Use a regular expression to find potential IAM keys in each line of the file
		re := regexp.MustCompile(`(?i)(?P<access_key_id>[A-Z0-9]{20})(?P<secret_access_key>[a-zA-Z0-9/+=]{40})`)
		matches := re.FindAllStringSubmatch(line, -1)

		for _, match := range matches {
			accessKeyID := match[1]
			secretAccessKey := match[2]

			fmt.Printf("Potential AWS IAM Key found in file: %s\nAccess Key ID: %s\nSecret Access Key: %s\n", filePath, accessKeyID, secretAccessKey)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading file: %s", err)
	}
}

func scanGitRepo(repoPath string) {
	r, err := git.PlainOpen(repoPath)
	if err != nil {
		log.Fatalf("Error opening Git repository: %s", err)
	}

	refIter, err := r.Branches()
	if err != nil {
		log.Fatalf("Error iterating through branches: %s", err)
	}

	err = refIter.ForEach(func(ref *plumbing.Reference) error {
		// Check out each branch
		if ref.Name().IsBranch() {
			branchName := ref.Name().Short()
			fmt.Println("Scanning branch:", branchName)

			w, err := r.Worktree()
			if err != nil {
				return err
			}

			err = w.Checkout(&git.CheckoutOptions{
				Branch: ref.Name(),
			})
			if err != nil {
				return err
			}

			// Walk through the repository's working directory to find files
			err = filepath.Walk(repoPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				if !info.IsDir() {
					scanFileForKeys(path)
				}
				return nil
			})
			return err
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error: %s", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the path to the Git repository.")
	}

	repoPath := os.Args[1]
	scanGitRepo(repoPath)
}
