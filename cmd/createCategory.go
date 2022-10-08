package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// createCategoryCmd represents the createCategory command
var createCategoryCmd = &cobra.Command{
	Use:   "createCategory",
	Short: "Create a new product category",
	Long:  `This command creates a new product category.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createCategory called")
	},
}

func init() {
	rootCmd.AddCommand(createCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
