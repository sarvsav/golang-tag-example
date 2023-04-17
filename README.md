# golang-tag-example

An example repo that demonstrates how tagging works in golang and github

## What is this repo about?

This repo is a demonstration on how to use stable or custom tags while using
the `go get ...` command. This is useful when you want to use a specific version
of a package in your project.

## How to use this repo? Follow the below instructions

Start by switching to the v0.0.1 branch.

```bash
git checkout -b v0.0.1
```

We will start by initializing a new go module here. It can be done using the
following command:

```bash
go mod init github.com/your_username/golang-tag-example
```

Once the module is initialized, we notice that there is one file present in the
current directory with named `go.mod` Or you can check the branch v0.0.1 to see
the file.

Switch back to main branch and merge your changes.

```bash
git checkout main
git merge v0.0.1
```

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

And, performed the below steps:

```bash
## Add and commit the changes
git add main.go 
git commit -m "Update main.go"

## Created the stable tag
git tag stable

## Checkout and merge to main branch
git checkout main
git merge v1.0.0

## Push everything to remote
git push origin --all
git push origin --tags
```

Now, let's try to install the package using the following command:

```bash
go install github.com/sarvsav/golang-tag-example@stable
```

Now, run the binary and verify the output:

```bash
$ golang-tag-example.exe 
Hello, stable
```

Hope, it helps you to understand how to use tags in golang.

## Adding Dockerfile

Create a Dockerfile in project root that contains information on how to build
the project.

## Building the image

We need to build the image using the following command:

```bash
docker build -t golang-tag-example .
```

## Running the image

To verify the image, we can run the following command that starts the container
and prints the output on the screen.

```bash
docker run --name learningtherightway-example golang-tag-example
```

## Distributing the image using container registry

The docker image now needs to be build using CI system and stored in a container
registry for distribution. There are different container registries available
like docker hub, ghcr, etc. We will use GitHub Container Registry for this.

## GitHub Actions for creating image

We will use GitHub Actions to build the image whenever there is a new commit to
`main` branch or new tag is created. We will use the `docker/build-push-action`
in this example.

The action file will be stored in `.github/workflows/docker.yml`.

## Using go-feature-flag

Install the go module using the following command:

```bash
go get github.com/thomaspoignant/go-feature-flag
```

There is already a repo exists containing the featureflag configuration. Here
are the repo details: [codingtherightway-ff](https://github.com/sarvsav/codingtherightway-ff)

## USing go-feature-flag in the code

```go
// Start with initializing the SDK
 err := ffclient.Init(ffclient.Config{
  PollingInterval: 10 * time.Second,
  Logger:          log.New(os.Stdout, "", 0),
  Context:         context.Background(),
  Retriever: &githubretriever.Retriever{
   RepositorySlug: "sarvsav/codingtherightway-ff",
   Branch:         "main",
   FilePath:       "config/flag-config.yaml",
   Timeout:        3 * time.Second,
  },
 })
```

Creating the users to test the feature flag:

```go
 // create users
 // Anonymous user
 user1 := ffuser.NewAnonymousUser("aea2fdc1-b9a0-417a-b707-0c9083de68e3")
 // User with custom attributes
 user2 := ffuser.NewUserBuilder("785a14bf-d2c5-4caa-9c70-2bbc4e3732a5").
  AddCustom("guest", true).Build()
```

Below code with fetch the flag value for the users

```go
 user1HasAccessToNewAdmin, err := ffclient.BoolVariation("new-admin-access", user1, false)
 user2HasAccessToNewAdmin, err := ffclient.BoolVariation("flag-only-for-admin", user2, false)
```

## Verifying the result

As the custom key `admin` is not set for user2, hence the feature flag will be
disabled for user2.

```bash
$ go run main.go 
Hello, stable
[2023-04-17T17:45:25+02:00] flag new-admin-access added
[2023-04-17T17:45:25+02:00] flag flag-only-for-admin added
user1 has access to the new admin
```
