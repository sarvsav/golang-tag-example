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

Checkout to new branch named `v0.0.2` using the following command:

```bash
git checkout -b v0.0.2
```

Create a new file named `main.go` and add the following code:

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, v0.0.2")
}
```

Now, we have tagged the branch v0.0.2 with tag `alpha`. Now, let's try to
install the package using the following command:

```bash
go install github.com/sarvsav/golang-tag-example@alpha
```

BONUS: To push all the branches and tags for the remote repo, use the following
command:

```bash
git push origin --all ## To push all the branches
git push origin --tags ## To push all the tags
```

Now, on running the below commands:
```bash
$ go install github.com/sarvsav/golang-tag-example@alpha
go: downloading github.com/sarvsav/golang-tag-example v0.0.0-20230413191708-b8f379cb72fb

$ golang-tag-example.exe 
Hello, v0.0.2
```

We can see that the package is installed and the output is as expected from
v0.0.2.

Now, let's repeat the same process and create a stable revision.

```bash
git checkout -b v1.0.0 ## Checkout to stable branch
```

Update the following code to `main.go`:

```go
fmt.Println("Hello, stable")
```
