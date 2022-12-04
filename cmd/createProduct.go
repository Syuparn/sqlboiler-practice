package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/syuparn/sqlboilerpractice/di"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

// createProductCmd represents the createProduct command
var createProductCmd = &cobra.Command{
	Use:   "createProduct",
	Short: "Create a new product",
	Long:  `This command creates a new product.`,
	Run:   createProduct,
}

func createProduct(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	price, _ := cmd.Flags().GetInt64("price")
	categoryID, _ := cmd.Flags().GetString("categoryid")
	in := &usecase.CreateProductInputData{
		Name:       name,
		Price:      price,
		CategoryID: categoryID,
	}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.CreateProductInputPort) {
		_, perr := p.Handle(ctx, in)
		if perr != nil {
			fmt.Fprintln(os.Stderr, perr.Error())
			os.Exit(1)
		}
	})

	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(createProductCmd)

	createProductCmd.Flags().String("name", "", "name of product")
	createProductCmd.Flags().Int64("price", 0, "price of product")
	createProductCmd.Flags().String("categoryid", "", "product category id of product")
}
