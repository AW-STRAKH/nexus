package main

import (
	"context"
	"log"

	"github.com/awatansh/nexus/internal/app"
	"github.com/awatansh/nexus/internal/identity"
	"github.com/awatansh/nexus/internal/storage/sqlite"
)

func main() {

	db, err := sqlite.Open(
		"./data/nexus.db",
	)

	if err != nil {
		log.Fatal(err)
	}

	if err := sqlite.Migrate(db); err != nil {
		log.Fatal(err)
	}

	identityStore :=
		sqlite.NewIdentityStore(db)

	identityService :=
		identity.NewService(
			identityStore,
		)

	node := app.NewNode(
		identityService,
	)

	if err := node.Start(
		context.Background(),
	); err != nil {

		log.Fatal(err)
	}
}
