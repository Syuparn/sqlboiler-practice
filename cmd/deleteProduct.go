package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/syuparn/sqlboilerpractice/di"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

// deleteProductCmd represents the deleteProduct command
var deleteProductCmd = &cobra.Command{
	Use:   "deleteProduct",
	Short: "Delete the specified product",
	Long:  `This command deletes the specified product.`,
	Run:   deleteProduct,
}

func deleteProduct(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	in := &usecase.DeleteProductInputData{
		ID: id,
	}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.DeleteProductInputPort) {
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
	rootCmd.AddCommand(deleteProductCmd)

	deleteProductCmd.Flags().String("id", "", "id of product")
}
