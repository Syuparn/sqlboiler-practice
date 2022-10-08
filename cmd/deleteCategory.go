package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deleteCategoryCmd represents the deleteCategory command
var deleteCategoryCmd = &cobra.Command{
	Use:   "deleteCategory",
	Short: "Delete the specified product category",
	Long:  `This command deletes the specified product category.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("deleteCategory called")
	},
}

func init() {
	rootCmd.AddCommand(deleteCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
