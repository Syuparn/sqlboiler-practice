package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/syuparn/sqlboilerpractice/di"
	"github.com/syuparn/sqlboilerpractice/usecase"
)

// createCategoryCmd represents the createCategory command
var createCategoryCmd = &cobra.Command{
	Use:   "createCategory",
	Short: "Create a new product category",
	Long:  `This command creates a new product category.`,
	Run:   createCategory,
}

func createCategory(cmd *cobra.Command, args []string) {
	name, _ := cmd.Flags().GetString("name")
	in := &usecase.CreateCategoryInputData{
		Name: name,
	}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.CreateCategoryInputPort) {
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
	rootCmd.AddCommand(createCategoryCmd)

	createCategoryCmd.Flags().String("name", "", "name of category")
}
