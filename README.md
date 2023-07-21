# Comprehensive Documentation

## Problem Approach:
The problem is to create a GoLang script that scans a Git repository for embedded AWS IAM (Identity and Access Management) keys. IAM keys are sensitive credentials that provide access to AWS services, and embedding them in code can pose a significant security risk. The script's goal is to identify potential occurrences of IAM keys in the repository's code, including both the latest code and historical commits in all branches.

To achieve this, the script uses the GoLang Git library "go-git" to interact with the Git repository. It walks through the repository's working directory and subdirectories, scanning each file for patterns that resemble IAM keys using regular expressions. The script then prints any potential IAM keys found in the files.

The script includes the following steps:

1. function isValidAWSKey to validate IAM keys by making a basic AWS API call.
2. function scanFileForKeys to scan each file for potential IAM keys using regular expressions.
3. function scanGitRepo to scan the entire Git repository, checking out each branch, and scanning all files in the working directory and its subdirectories.

Main function to receive the repository path as an argument and call the scanGitRepo function.
Solution:
The provided GoLang script uses regular expressions to search for potential IAM keys in all files of a Git repository. It will not validate the IAM keys with AWS but will only print potential matches. The goal is to identify and prompt the user to review any potential occurrences of IAM keys found in the code and remove them if necessary.

## Running the Code:
Here are the steps to run the GoLang script to scan a Git repository for embedded AWS IAM keys:

1. Install the required packages using go get:

```bash
go get -u github.com/go-git/go-git/v5
```
2. Save the provided GoLang script in a file named scan_aws_keys.go.

3. Build the GoLang script:

```bash
go build scan_aws_keys.go
```
4. Clone the sample repository containing potential IAM keys:
```bash
git clone https://github.com/abhishek-pingsafe/Devops-Node
```
5. Run the script, providing the path to the cloned repository as an argument:
```bash
Copy code
./scan_aws_keys Devops-Node
```
The script will scan the entire repository and print any potential IAM keys found in the files.