package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pgen = &cobra.Command{
	Use: "pgen",

	Run: genPass,
}

func init() {
	rootCmd.AddCommand(pgen)

	pgen.Flags().StringP("social", "s", "None", "Social network password")

}

func genPass(cmd *cobra.Command, args []string) {

	social, _ := cmd.Flags().GetString("social")

	fmt.Println("Gen pass for", social)

}
