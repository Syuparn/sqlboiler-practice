package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteProductCmd represents the deleteProduct command
var deleteProductCmd = &cobra.Command{
	Use:   "deleteProduct",
	Short: "Delete the specified product",
	Long:  `This command deletes the specified product.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteProduct called")
	},
}

func init() {
	rootCmd.AddCommand(deleteProductCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteProductCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteProductCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
