package internal

import (
	"log"

	"github.com/amirhnajafiz/tokenizer/pkg/stdout"

	"github.com/spf13/cobra"
)

// CCommands is a struct that holds tokenizer main commands.
type CCommands struct {
	Stdout *stdout.Stdout
}

// SetToken into the conf file.
func (c CCommands) SetToken() *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Set token",
		Long:  "Set a new token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 2 {
				log.Fatal(ErrParams)
			}

			if err := Set(args[0], args[1]); err != nil {
				log.Fatalln(err)
			}
		},
	}
}

// GetToken from the conf file.
func (c CCommands) GetToken() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get token",
		Long:  "Get an existing token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Fatal(ErrParams)
			}

			value, err := Get(args[0])
			if err != nil {
				log.Fatalln(err)
			}

			c.Stdout.Print(value)
		},
	}
}

// DeleteToken from the conf file.
func (c CCommands) DeleteToken() *cobra.Command {
	return &cobra.Command{
		Use:   "del",
		Short: "Delete token",
		Long:  "Delete an existing token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal(ErrParams)
			}

			if err := Remove(args[0]); err != nil {
				log.Fatalln(err)
			}
		},
	}
}

// GetAllTokens from the conf file.
func (c CCommands) GetAllTokens() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "All tokens",
		Long:  "Get all existing tokens",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				log.Fatal(ErrParams)
			}

			keys, err := GetKeys()
			if err != nil {
				log.Fatalln(err)
			}

			c.Stdout.Print(keys...)
		},
	}
}
