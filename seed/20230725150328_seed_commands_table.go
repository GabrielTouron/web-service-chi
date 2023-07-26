package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	"github.com/dustinkirkland/golang-petname"
	"github.com/pressly/goose/v3"
	"math/rand"
)

var (
	word      = flag.Int("words", 1, "The number of words in the pet name")
	separator = flag.String("separator", "-", "The separator between words in the pet name")
)

func init() {
	goose.AddMigrationContext(upSeedCommandsTable, downSeedCommandsTable)
}

func upSeedCommandsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	// Seed commands table with 1000 commands
	for i := 0; i < 1500; i++ {
		randomInt := rand.Intn(1000)
		randomIntString := fmt.Sprintf("%03d", randomInt)

		query := "INSERT INTO commands(name, command) VALUES($1, $2)"

		if _, err := tx.Exec(query, petname.Generate(*word, *separator), randomIntString); err != nil {
			return err
		}
	}

	return nil
}

func downSeedCommandsTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec("DELETE FROM commands")
	if err != nil {
		return err
	}

	return nil
}
