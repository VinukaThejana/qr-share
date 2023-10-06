package cmd

import (
	"os"

	"github.com/VinukaThejana/qr-share/controller"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "qr",
	Short: "Share within the local network",
	Long: `
Share within the local network

EXAMPLES
  qr
    serve the files within the current working directory over the local network

  qr filename
    server the file of the filename over the local network

CORE COMMANDS
  dir
    serve the contents of the custom path over the local network
`,
	Run: func(_ *cobra.Command, args []string) {
		if len(args) == 0 {
			controller.Serve("./")
		} else {
			controller.Serve(args[0])
		}
	},
}

// Execute the root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
