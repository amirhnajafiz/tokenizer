package internal

import (
	"fmt"
	"github.com/spf13/cobra"
)

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
