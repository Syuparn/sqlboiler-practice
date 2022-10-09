package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/syuparn/sqlboilerpractice/di"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

// deleteCategoryCmd represents the deleteCategory command
var deleteCategoryCmd = &cobra.Command{
	Use:   "deleteCategory",
	Short: "Delete the specified product category",
	Long:  `This command deletes the specified product category.`,
	Run:   deleteCategory,
}

func deleteCategory(cmd *cobra.Command, args []string) {
	id, _ := cmd.Flags().GetString("id")
	in := &usecase.DeleteCategoryInputData{
		ID: id,
	}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.DeleteCategoryInputPort) {
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
	rootCmd.AddCommand(deleteCategoryCmd)

	deleteCategoryCmd.Flags().String("id", "", "id of category")
}
