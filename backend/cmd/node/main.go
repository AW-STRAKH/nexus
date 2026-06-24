package main

import (
	"context"
	"flag"
	"log"

	"github.com/awatansh/nexus/internal/app"
	"github.com/awatansh/nexus/internal/identity"
	"github.com/awatansh/nexus/internal/storage/sqlite"
	quictransport "github.com/awatansh/nexus/internal/transport/quic"
)

func main() {

	port := flag.String(
		"port",
		"9000",
		"listening port",
	)

	connect := flag.String(
		"connect",
		"",
		"peer address to connect to",
	)

	flag.Parse()

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

	quicConfig :=
		quictransport.DefaultConfig()

	quicConfig.ListenAddress =
		":" + *port

	transport :=
		quictransport.NewTransport(
			quicConfig,
		)

	node := app.NewNode(
		identityService,
		transport,
		*connect,
	)

	if err := node.Start(
		context.Background(),
	); err != nil {

		log.Fatal(err)
	}

	select {}
}
