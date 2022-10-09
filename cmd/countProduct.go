package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/syuparn/sqlboilerpractice/di"
	"github.com/syuparn/sqlboilerpractice/domain"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

// countProductCmd represents the countProduct command
var countProductCmd = &cobra.Command{
	Use:   "countProduct",
	Short: "Count the number of products in the specified category",
	Long:  `This command counts the number of products in the specified category.`,
	Run:   countProduct,
}

func countProduct(cmd *cobra.Command, args []string) {
	categoryID, _ := cmd.Flags().GetString("categoryid")
	in := &usecase.CountProductInputData{
		CategoryID: categoryID,
	}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.CountProductInputPort) {
		out, perr := p.Handle(ctx, in)
		if perr != nil {
			fmt.Fprintln(os.Stderr, perr.Error())
			os.Exit(1)
		}

		func(s *domain.CategoryStatistics) {
			fmt.Printf("category_id: %s, category_name: %s, num_products: %d\n", s.CategoryID, s.CategoryName, s.NumProducts)
		}(out.CategoryStatistics)
	})

	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(countProductCmd)

	countProductCmd.Flags().String("categoryid", "", "product category id")
}
