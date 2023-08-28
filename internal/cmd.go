package internal

import (
	"fmt"

	"github.com/spf13/cobra"
)

// SetToken into conf file
func SetToken() *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "set token",
		Long:  "set a new token",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
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
			fmt.Println(args)
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
			fmt.Println(args)
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
			fmt.Println(args)
		},
	}
}
