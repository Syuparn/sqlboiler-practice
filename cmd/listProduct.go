package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/samber/lo"
	"github.com/spf13/cobra"

	"github.com/syuparn/sqlboilerpractice/di"
	"github.com/syuparn/sqlboilerpractice/domain"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

// listProductCmd represents the listProduct command
var listProductCmd = &cobra.Command{
	Use:   "listProduct",
	Short: "List products",
	Long:  `This command lists products.`,
	Run:   listProduct,
}

func listProduct(cmd *cobra.Command, args []string) {
	in := &usecase.ListProductInputData{}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.ListProductInputPort) {
		out, perr := p.Handle(ctx, in)
		if perr != nil {
			fmt.Fprintln(os.Stderr, perr.Error())
			os.Exit(1)
		}

		lo.ForEach(out.Products, func(c *domain.Product, _ int) {
			fmt.Printf("id: %s, name: %s, price: %d, category_id: %s\n", c.ID, c.Name, c.Price, c.CategoryID)
		})
	})

	if err != nil {
		panic(err)
	}
}

func init() {
	rootCmd.AddCommand(listProductCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listProductCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listProductCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
