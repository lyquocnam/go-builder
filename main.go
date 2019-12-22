package main

import (
	"github.com/lyquocnam/go-builder/module"
	"github.com/spf13/cobra"
	"log"
)

func main() {
	rootCmd := &cobra.Command{
		Short:                      "Command:",
		Long:                       "Command:",
	}

	rootCmd.AddCommand(downloadCommand())
	rootCmd.AddCommand(dockerBuildCommand())
	rootCmd.AddCommand(dockerRunCommand())

	rootCmd.Execute()
}


func downloadCommand() *cobra.Command {
	downloadCmd := &cobra.Command{
		Use: "download",
		Short:                      "Download template",
		Long:                       "Download template",
		Run: func(cmd *cobra.Command, args []string) {
			forceOverride, err := cmd.Flags().GetBool("force")
			if err != nil {
				log.Fatalln(err)
			}

			logger := module.NewLogger()
			module.NewDownloader(logger).Run(forceOverride)
		},
	}

	downloadCmd.Flags().BoolP("force", "f", false,"force override written file")

	return downloadCmd
}

func dockerBuildCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "build",
		Short:                      "Build app from docker template",
		Long:                       "Build app from docker template",
		Run: func(cmd *cobra.Command, args []string) {
			tagName, err := cmd.Flags().GetString("tag")
			if err != nil {
				log.Fatalln(err)
			}

			mode, err := cmd.Flags().GetString("mode")
			if err != nil {
				log.Fatalln(err)
			}
			logger := module.NewLogger()
			docker := module.NewDocker(logger)
			if mode == "development" {
				docker.BuildDev(tagName)
			} else if mode == "production" {
				docker.BuildProd(tagName)
			}
		},
	}

	cmd.Flags().StringP("tag", "t", "","docker tag name")
	cmd.Flags().StringP("mode", "m", "development","environment mode, default: development")

	return cmd
}


func dockerRunCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "run",
		Short:                      "Run app from docker",
		Long:                       "Run app from docker",
		Run: func(cmd *cobra.Command, args []string) {
			serviceName, err := cmd.Flags().GetString("name")
			if err != nil {
				log.Fatalln(err)
			}

			logger := module.NewLogger()
			module.NewDocker(logger).Run(serviceName)
		},
	}

	cmd.Flags().StringP("name", "n", "","docker service name")

	return cmd
}