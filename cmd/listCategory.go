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

// listCategoryCmd represents the listCategory command
var listCategoryCmd = &cobra.Command{
	Use:   "listCategory",
	Short: "list product categories",
	Long:  `This command lists product categories.`,
	Run:   listCategory,
}

func listCategory(cmd *cobra.Command, args []string) {
	in := &usecase.ListCategoryInputData{}

	c := di.NewContainer()
	ctx := context.Background()
	err := c.Invoke(func(p usecase.ListCategoryInputPort) {
		out, perr := p.Handle(ctx, in)
		if perr != nil {
			fmt.Fprintln(os.Stderr, perr.Error())
			os.Exit(1)
		}

		lo.ForEach(out.Categories, func(c *domain.Category, _ int) {
			fmt.Printf("id: %s, name: %s\n", c.ID, c.Name)
		})
	})

	if err != nil {
		panic(err)
	}
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
