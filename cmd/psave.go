/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// psaveCmd represents the psave command
var psaveCmd = &cobra.Command{
	Use:   "psave",
	Short: "Store your password on a text encrypted file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("psave called")
	},
}

func init() {
	rootCmd.AddCommand(psaveCmd)

}
