package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// countProductCmd represents the countProduct command
var countProductCmd = &cobra.Command{
	Use:   "countProduct",
	Short: "Count the number of products in each category",
	Long:  `This command counts the number of products in each category.`,
	Run:   countProduct,
}

func countProduct(cmd *cobra.Command, args []string) {
	fmt.Println("countProduct called")
}

func init() {
	rootCmd.AddCommand(countProductCmd)
}
