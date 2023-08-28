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
		internal.GetToken(),
		internal.DeleteToken(),
		internal.GetAllTokens(),
	)

	// todo: Add encryption
	// todo: Feat encrypt option

	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
