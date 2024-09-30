package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Udehlee/issue-tracka/issue"
)

func ReadInput() (string, error) {
	r := bufio.NewReader(os.Stdin)

	input, err := r.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input, nil
}

func EnterIssue() (string, string) {
	fmt.Println("Enter the title of the issue:")
	title, _ := ReadInput()

	fmt.Println("Enter the text of the issue:")
	text, _ := ReadInput()

	return title, text
}

func main() {

	input := flag.String("command", "", "Specify the command: create, list, open, or add-comment")
	flag.Parse()

	is := issue.NewMemory()

	switch *input {
	case "create":
		title, text := EnterIssue()

		issu, err := is.Create(title, text)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		if err = is.Save(issu); err != nil {
			fmt.Println("Error saving issue:", err)
			return
		}

		issueJson := is.JSON(issu)
		fmt.Println("Issue created successfully:", issueJson)

	case "list":
		list, err := is.List()
		if err != nil {
			fmt.Println("Failed to list issues:", err)
		}

		listJSON := is.JSON(list)
		fmt.Println(" all issues", listJSON)

	case "open":
		if len(os.Args) <= 2 {
			fmt.Println("Please provide an issue ID to open")
			return
		}

		Id := os.Args[2]

		openIssue, err := is.Open(Id)
		if err != nil {
			fmt.Println("Failed to open issue:", err)
			return
		}

		openJSON := is.JSON(openIssue)
		fmt.Println("issue found", openJSON)

	case "add-comment":
		if len(os.Args) <= 4 {
			fmt.Println("Please provide issue ID, commenter name, and comment text")
			return
		}

		var issueId, name, text string
		args := []*string{&issueId, &name, &text}

		for i, arg := range os.Args[1:4] {
			*args[i] = arg
		}

		err := is.AddComment(issueId, name, text)
		if err != nil {
			fmt.Println("there is error", err)
			return
		}
		fmt.Println("added successfully")

	default:
		fmt.Println("Invalid command. Please use one of the following: create, list, open, add-comment")
	}

}
