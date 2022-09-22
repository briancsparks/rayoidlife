package cmd

import (
	"fmt"
  "github.com/briancsparks/raylib/raygun"

  "github.com/spf13/cobra"
)

// -------------------------------------------------------------------------------------------------------------------

// twoCmd represents the two command
var twoCmd = &cobra.Command{
	Use:   "two",
	Short: "A two",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("two called")
    raygun.MainTwo()
	},
}

// -------------------------------------------------------------------------------------------------------------------

func init() {
	rootCmd.AddCommand(twoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// twoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// twoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
