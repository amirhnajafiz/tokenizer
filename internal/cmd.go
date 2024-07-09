package internal

import (
	"fmt"
	"log"
	"os/exec"

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
			keys, err := GetKeys()
			if err != nil {
				log.Fatalln(err)
			}

			c.Stdout.Print(keys...)
		},
	}
}

// Backup creates a config file with current data, all keys and values init.
func (c CCommands) Backup() *cobra.Command {
	return &cobra.Command{
		Use:   "backup",
		Short: "Backup current tokens",
		Long:  "Backup decrypted current tokens into a file",
		Run: func(cmd *cobra.Command, args []string) {
			sudoer := exec.Command("/bin/sh", "-c", "sudo ls")
			if _, err := sudoer.Output(); err != nil {
				log.Fatalf("failed to run backup command: %v\n", err)
			}

			mapping := make([]string, 0)

			keys, err := GetKeys()
			if err != nil {
				log.Fatalln(err)
			}

			for _, key := range keys {
				if value, err := Get(key); err == nil {
					mapping = append(mapping, fmt.Sprintf("%s = %s", key, value))
				}
			}

			c.Stdout.Print(mapping...)
		},
	}
}
