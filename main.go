package main

import (
	"log"

	"github.com/amirhnajafiz/tokenizer/internal"
	"github.com/amirhnajafiz/tokenizer/pkg/stdout"

	"github.com/spf13/cobra"
)

func main() {
	// init root cobra command and stdout instance
	root := cobra.Command{}
	std := stdout.Stdout{}

	// create cmd flags
	root.PersistentFlags().StringVarP(&std.Path, "file", "f", "out.txt", "a file path to export the content into it")
	root.PersistentFlags().BoolVarP(&std.File, "output", "o", false, "if set to true, it will export the output data into a file (works on get and all commands)")
	root.PersistentFlags().BoolVarP(&std.Clipboard, "copy", "c", false, "if set to true, it will export the output data into clipboard (works on get command)")

	// create file in set
	if !internal.CheckFile() {
		_ = internal.InitFile()
	}

	// create the internal.ccommands
	cc := internal.CCommands{}

	// bind cobra commands to cc commands
	root.AddCommand(
		cc.SetToken(),
		cc.GetToken(),
		cc.DeleteToken(),
		cc.GetAllTokens(),
	)

	// execute root cobra command
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
