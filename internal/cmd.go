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

			key, _ := Code(args[0])
			value, _ := Code(args[1])

			if err := Set(key, value); err != nil {
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

			key, _ := Code(args[0])

			value, err := Get(key)
			if err != nil {
				log.Fatalln(err)
			}

			tmp, _ := DeCode(value)

			fmt.Println(tmp)
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

			key, _ := Code(args[0])

			if err := Remove(key); err != nil {
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

			for _, key := range keys {
				tmp, _ := DeCode(key)

				fmt.Println(tmp)
			}
		},
	}
}
