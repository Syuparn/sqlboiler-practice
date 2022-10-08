package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCategoryCmd represents the listCategory command
var listCategoryCmd = &cobra.Command{
	Use:   "listCategory",
	Short: "list product categories",
	Long:  `This command lists product categories.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listCategory called")
	},
}

func init() {
	rootCmd.AddCommand(listCategoryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCategoryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCategoryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
