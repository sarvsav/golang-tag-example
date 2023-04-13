# golang-tag-example

An example repo that demonstrates how tagging works in golang and github

## What is this repo about?

This repo is a demonstration on how to use stable or custom tags while using
the `go get ...` command. This is useful when you want to use a specific version
of a package in your project.

## How to use this repo? Follow the below instructions:

We will start by initializing a new go module here. It can be done using the
following command:

```bash
go mod init github.com/your_username/golang-tag-example
```

Once the module is initialized, we notice that there is one file present in the
current directory with named `go.mod` Or you can check the branch v0.0.1 to see
the file.

Create a new file named `main.go` and add the following code:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, v0.0.2")
}
```
