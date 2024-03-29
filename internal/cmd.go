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
		Short: "set token",
		Long:  "set a new token",
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
func GetToken() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "get token",
		Long:  "get an existing token",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				log.Fatal(ErrParams)
			}

			value, err := Get(args[0])
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println()
			fmt.Println(value)
		},
	}
}

// DeleteToken from conf file
func DeleteToken() *cobra.Command {
	return &cobra.Command{
		Use:   "del",
		Short: "delete token",
		Long:  "delete an existing token",
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
func GetAllTokens() *cobra.Command {
	return &cobra.Command{
		Use:   "all",
		Short: "all tokens",
		Long:  "get all existing tokens",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 0 {
				log.Fatal(ErrParams)
			}

			keys, err := GetKeys()
			if err != nil {
				log.Fatalln(err)
			}

			fmt.Println()

			for _, key := range keys {
				fmt.Println(key)
			}
		},
	}
}
