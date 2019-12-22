package main

import (
	"github.com/lyquocnam/go-builder/module"
	"github.com/spf13/cobra"
)

func main() {
	downloadCmd := &cobra.Command{
		Use: "download",
		Short:                      "Download template",
		Long:                       "Download template",
		Run: func(cmd *cobra.Command, args []string) {
			module.NewDownloader().Run()
		},
	}

	rootCmd := &cobra.Command{
		Short:                      "Command:",
		Long:                       "Command:",
	}

	rootCmd.AddCommand(downloadCmd)

	rootCmd.Execute()
}