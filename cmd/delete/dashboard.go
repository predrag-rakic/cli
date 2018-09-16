package cmd_delete

import (
	"fmt"

	client "github.com/semaphoreci/cli/api/client"

	"github.com/semaphoreci/cli/cmd/utils"
	"github.com/spf13/cobra"
)

var DeleteDashboardCmd = &cobra.Command{
	Use:     "dashboard [NAME]",
	Short:   "Delete a dashboard.",
	Long:    ``,
	Aliases: []string{"dashboards", "dash"},
	Args:    cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		name := args[0]

		c := client.NewDashboardV1AlphaApi()

		err := c.DeleteDashboard(name)

		utils.Check(err)

		fmt.Printf("Dashboard '%s' deleted.\n", name)
	},
}
