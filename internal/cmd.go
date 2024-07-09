package internal

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// SetToken into conf file
func SetToken() *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Set token",
		Long:  "Set a new token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 2 {
				log.Fatal(ErrParams)
			}

			if err := Set(args[0], args[1]); err != nil {
				log.Fatalln(err)
			}
		},
	}
}

// GetToken from conf file
func GetToken(export bool, path string) *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get token",
		Long:  "Get an existing token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal(ErrParams)
			}

			value, err := Get(args[0])
			if err != nil {
				log.Fatalln(err)
			}

			if export {
				if err := exportToFile(path, value); err != nil {
					log.Fatalln(err)
				}
			} else {
				fmt.Println(value)
			}
		},
	}
}

// DeleteToken from conf file
func DeleteToken() *cobra.Command {
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

// GetAllTokens from conf file
func GetAllTokens(export bool, path string) *cobra.Command {
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

			if export {
				if err := exportToFile(path, keys...); err != nil {
					log.Fatalln(err)
				}
			} else {
				for _, key := range keys {
					fmt.Println(key)
				}
			}
		},
	}
}
