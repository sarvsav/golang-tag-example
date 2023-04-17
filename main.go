package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/ffuser"
	"github.com/thomaspoignant/go-feature-flag/retriever/githubretriever"
)

func main() {
	fmt.Println("Hello, stable")

	// Init client with a GitHub retriever.
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

	// Check init errors.
	if err != nil {
		log.Fatal(err)
	}
	// defer closing ffclient
	defer ffclient.Close()

	// create users
	user1 := ffuser.NewAnonymousUser("aea2fdc1-b9a0-417a-b707-0c9083de68e3")
	user2 := ffuser.NewUserBuilder("785a14bf-d2c5-4caa-9c70-2bbc4e3732a5").
		AddCustom("guest", true).Build()

	// --- test flag with no rule
	// user1
	user1HasAccessToNewAdmin, err := ffclient.BoolVariation("new-admin-access", user1, false)
	if err != nil {
		// we log the error, but we still have a meaningful value in user1HasAccessToNewAdmin (the default value).
		log.Printf("something went wrong when getting the flag: %v", err)
	}
	if user1HasAccessToNewAdmin {
		fmt.Println("user1 has access to the new admin")
	}

	user2HasAccessToNewAdmin, err := ffclient.BoolVariation("flag-only-for-admin", user2, false)
	if err != nil {
		// we log the error, but we still have a meaningful value in user1HasAccessToNewAdmin (the default value).
		log.Printf("something went wrong when getting the flag: %v", err)
	}
	if user2HasAccessToNewAdmin {
		fmt.Println("user2 has access to the admin only flag")
	}
}
