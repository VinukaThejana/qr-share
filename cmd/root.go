package cmd

import (
	"fmt"
	"os"

	"github.com/VinukaThejana/qr-share/utils"
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
		var path string
		if len(args) == 0 {
			path = "./"
		} else {
			path = args[0]
			if _, err := os.Stat(path); err != nil {
				utils.Text{}.Error(fmt.Sprintf("%s : filename is not valid\n", path))
				return
			}
		}

		fmt.Println(path)
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
