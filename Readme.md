## issue-tracka

This is a simple command line project for tracking your issues.It is similar to github issues. It is written in Go.

# How to run the program
```sh
go run main.go -command=[options]
```

# Options
- ```sh create ``` Create a new issue
- ```sh list ``` list all issues present
- ``` sh open ``` Open a issue using its Id
- ```sh add-comment ``` add-comment adds comments to existing issue

# Usage

- Create a new issue
```sh
go run main.go -command=create
```

- list all issues
```sh
go run main.go -command=list 
```

- Open an issue with a specific Id
```sh
go run main.go -command=open "bb982b"
```

- add-comment to an issue to a specific Id
```sh
go run main.go -command=open "bb982b"
```

# Keep in mind
- when you run the ```sh go run main.go -command=create ``` , you will be asked to enter title and text of issue.

# Testing
- To run test

```sh
go test -v
```