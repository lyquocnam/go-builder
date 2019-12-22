package main

import (
	"github.com/lyquocnam/go-builder/module"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	downloadCmd := &cobra.Command{
		Use: "download",
		Short:                      "Download template",
		Long:                       "Download template",
		Run: func(cmd *cobra.Command, args []string) {
			forceOverride, err := cmd.Flags().GetBool("force")
			if err != nil {
				log.Fatalln(err)
			}

			module.NewDownloader().Run(forceOverride)
		},
	}

	downloadCmd.Flags().BoolP("force", "f", false,"force override written file")

	rootCmd := &cobra.Command{
		Short:                      "Command:",
		Long:                       "Command:",
	}

	rootCmd.AddCommand(downloadCmd)

	rootCmd.Execute()
}