package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/philippgille/apiomat-go/aomc"
	"github.com/philippgille/apiomat-go/aoms"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints the version of the ApiOmat server",
	Long: `version is a command for printing the version of the ApiOmat server.

For example:
aom version --baseUrl "https://apiomat.yourcompany.com/yambas/rest"`,
	Run: func(cmd *cobra.Command, args []string) {
		client := aomc.NewDefaultClient(baseUrl, username, password, aoms.System(system))

		version, err := client.GetVersion()
		if err != nil {
			logFatal(err, "An error occurred during fetching the version of the ApiOmat server")
		}

		fmt.Printf("ApiOmat version: %v", version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
