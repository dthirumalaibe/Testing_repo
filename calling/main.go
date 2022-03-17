package main

import (
	"Testing/Public_repo/gpoll"
	"fmt"
	"log"

	"github.com/eddieowens/gpoll"
)

func main() {
	poller, err := gpoll.NewPoller(gpoll.PollConfig{
		Git: gpoll.GitConfig{
			Auth: gpoll.GitAuthConfig{
				// Uses the SSH key from my local directory.
				SshKey: "/Users/admin/go/src/Testing/Public_repo/ssh_key/thiru_id",
			},
			// The target remote.
			Remote: "git@github.com:dthirumalaibe/Testing_repo.git",
		},
		// OnUpdate: func(change gpoll.GitChange) {
		FileChangeFilter: func(change gpoll.GitChange) {
			switch change.ChangeType {
			case gpoll.ChangeTypeDelete:
				fmt.Printf("%s was deleted", change.Filename)
			case gpoll.ChangeTypeUpdate:
				fmt.Printf("%s was updated", change.Filename)
			case gpoll.ChangeTypeCreate:
				fmt.Printf("%s was created", change.Filename)
			}
		},
	})

	if err != nil {
		panic(err)
	}

	// Will poll the repo until poller.Stop() is called.
	log.Fatal(poller.Start())
}
