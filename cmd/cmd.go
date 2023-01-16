package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/uditgaurav/hce-api-template/apis"
	"github.com/uditgaurav/hce-api-template/types"
)

var LaunchChaos = &cobra.Command{
	Use:     "generate",
	Short:   "Launches Chaos Experiment Workflow",
	Long:    "Launches Chaos Experiment Workflow",
	Example: "./hce-api generate --api launch-experiment --hce-endpoint=http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/ --project-id abceb5f4-4268-4467-9818-ad6e3b6bfd78 --workflow-id f4581780-efaf-4155-956e-6c379f24394b --access-key nEdGNDDrTFHyCnl --access-id adminNCWQu --file-name hce-api.sh",
	Run: func(cmd *cobra.Command, args []string) {

		apiDetials := types.APIDetials{}
		mode := apis.CheckMode()

		api, err := cmd.Flags().GetString("api")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		switch api {
		case "launch-experiment":

			apiDetials.HCEEndpoint, err = cmd.Flags().GetString("hce-endpoint")
			apiDetials.ProjectID, err = cmd.Flags().GetString("project-id")
			apiDetials.WorkflowID, err = cmd.Flags().GetString("workflow-id")
			apiDetials.AccessKey, err = cmd.Flags().GetString("access-key")
			apiDetials.AccessID, err = cmd.Flags().GetString("access-id")
			apiDetials.FileName, err = cmd.Flags().GetString("file-name")
			apiDetials.API, err = cmd.Flags().GetString("api")

			if err := apis.ApiToLanchExperiment(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to launch experiment, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)

		case "monitor-experiment":
			apiDetials.HCEEndpoint, err = cmd.Flags().GetString("hce-endpoint")
			apiDetials.ProjectID, err = cmd.Flags().GetString("project-id")
			apiDetials.WorkflowID, err = cmd.Flags().GetString("workflow-id")
			apiDetials.AccessKey, err = cmd.Flags().GetString("access-key")
			apiDetials.AccessID, err = cmd.Flags().GetString("access-id")
			apiDetials.FileName, err = cmd.Flags().GetString("file-name")
			apiDetials.API, err = cmd.Flags().GetString("api")

			if err := apis.ApiToMonitorExperiment(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to monitor experiment, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)

		case "validate-resilience-score":
			apiDetials.HCEEndpoint, err = cmd.Flags().GetString("hce-endpoint")
			apiDetials.ProjectID, err = cmd.Flags().GetString("project-id")
			apiDetials.WorkflowID, err = cmd.Flags().GetString("workflow-id")
			apiDetials.AccessKey, err = cmd.Flags().GetString("access-key")
			apiDetials.AccessID, err = cmd.Flags().GetString("access-id")
			apiDetials.FileName, err = cmd.Flags().GetString("file-name")
			apiDetials.API, err = cmd.Flags().GetString("api")

			if err := apis.ApiToValidateResilienceScore(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to validate resilience score of the workflow, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)
		}

	},
}

var rootCmd = &cobra.Command{
	Use:   "hce-api",
	Short: "CLI to generate api",
}

func init() {

	rootCmd.AddCommand(LaunchChaos)

	LaunchChaos.Flags().String("api", "", "Set the name of target api")
	LaunchChaos.Flags().String("hce-endpoint", "", "Set the hce-endpoint for the given HCE")
	LaunchChaos.Flags().String("project-id", "", "Set the hce project id")
	LaunchChaos.Flags().String("workflow-id", "", "Set the workflow id")
	LaunchChaos.Flags().String("access-key", "", "Set the access key")
	LaunchChaos.Flags().String("access-id", "", "Set the access id")
	LaunchChaos.Flags().StringP("file-name", "f", "", "The target file name which contains the API command")
}
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
