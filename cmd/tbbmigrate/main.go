package main

import (
	"fmt"
	"os"
	"time"

	"github.com/sm1else-bot/nekocoin-blockchain/database"
)

func main() {
	state, err := database.NewStateFromDisk()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer state.Close()

	block0 := database.NewBlock(
		database.Hash{},
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("Jess", "Jess", 3, ""),
			database.NewTx("Jess", "Jess", 700, "reward"),
		},
	)

	state.AddBlock(block0)
	block0hash, _ := state.Persist()

	block1 := database.NewBlock(
		block0hash,
		uint64(time.Now().Unix()),
		[]database.Tx{
			database.NewTx("Jess", "Harshith", 2000, ""),
			database.NewTx("Jess", "Jess", 100, "reward"),
			database.NewTx("Harshith", "Jess", 1, ""),
			database.NewTx("Harshith", "Prasad", 1000, ""),
			database.NewTx("Harshith", "Jess", 50, ""),
			database.NewTx("Jess", "Jess", 600, "reward"),
		},
	)

	state.AddBlock(block1)
	state.Persist()
}
