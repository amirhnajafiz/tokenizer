package main

import (
	"log"

	"github.com/amirhnajafiz/tokenizer/internal"

	"github.com/spf13/cobra"
)

var (
	exportPath string
	exportFlag bool
)

func main() {
	root := cobra.Command{}

	// create cmd flags
	root.Flags().StringVarP(&exportPath, "path", "p", "out.txt", "a file path to export the content into it")
	root.Flags().BoolVarP(&exportFlag, "export", "e", false, "if set to true, it will export the output data into a file (works on get and all commands)")

	// create file in set
	if !internal.CheckFile() {
		_ = internal.InitFile()
	}

	// create the internal.ccommands
	cc := internal.CCommands{}

	// bind cobra commands to cc commands
	root.AddCommand(
		cc.SetToken(),
		cc.GetToken(exportFlag, exportPath),
		cc.DeleteToken(),
		cc.GetAllTokens(exportFlag, exportPath),
	)

	// execute root cobra command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
