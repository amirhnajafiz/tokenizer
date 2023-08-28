package main

import (
	"log"

	"github.com/amirhnajafiz/tokenizer/internal"

	"github.com/spf13/cobra"
)

func main() {
	root := cobra.Command{}

	root.AddCommand(
		internal.SetToken(),
	)

	// todo: Add set token
	// todo: Add view tokens
	// todo: Remove token
	// todo: Update token
	// todo: Add get token
	// todo: Add encryption
	// todo: Feat encrypt option

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
