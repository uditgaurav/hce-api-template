package cmd

import (
	"fmt"
	"os"

	launchChaos "github.com/uditgaurav/hce-api-template/apis/launch-chaos/lib"
	monitorChaos "github.com/uditgaurav/hce-api-template/apis/monitor-chaos/lib"
	validateRR "github.com/uditgaurav/hce-api-template/apis/validate-resilience-score/lib"

	"github.com/spf13/cobra"
	"github.com/uditgaurav/hce-api-template/pkg/common"
	"github.com/uditgaurav/hce-api-template/pkg/types"
)

var LaunchChaos = &cobra.Command{
	Use:     "generate",
	Short:   "Launches Chaos Experiment Workflow",
	Long:    "Launches Chaos Experiment Workflow",
	Example: "./hce-api generate --api launch-experiment --hce-endpoint=http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/ --project-id abceb5f4-4268-4467-9818-ad6e3b6bfd78 --workflow-id f4581780-efaf-4155-956e-6c379f24394b --access-key nEdGNDDrTFHyCnl --access-id adminNCWQu --file-name hce-api.sh",
	Run: func(cmd *cobra.Command, args []string) {

		apiDetials := types.APIDetials{}
		mode := common.CheckMode()

		api, err := cmd.Flags().GetString("api")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		apiDetials.FileName, err = cmd.Flags().GetString("file-name")
		apiDetials.HCEEndpoint, err = cmd.Flags().GetString("hce-endpoint")
		apiDetials.ProjectID, err = cmd.Flags().GetString("project-id")
		apiDetials.WorkflowID, err = cmd.Flags().GetString("workflow-id")
		apiDetials.AccessKey, err = cmd.Flags().GetString("access-key")
		apiDetials.AccessID, err = cmd.Flags().GetString("access-id")
		apiDetials.API, err = cmd.Flags().GetString("api")

		switch api {
		case "launch-experiment":

			if err := launchChaos.LaunchChaos(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to launch experiment, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)

		case "monitor-experiment":

			apiDetials.Delay, err = cmd.Flags().GetString("delay")
			apiDetials.Timeout, err = cmd.Flags().GetString("timeout")

			if err := monitorChaos.MonitorChaosExperiment(apiDetials, mode); err != nil {
				fmt.Printf("monitor chaos failed, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)

		case "validate-resilience-score":

			apiDetials.ExpectedResilienceScore, err = cmd.Flags().GetInt("100")

			if err := validateRR.ValidateResilienceScore(apiDetials, mode); err != nil {
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
	LaunchChaos.Flags().String("file-name", "", "The target file name which contains the API command")
	LaunchChaos.Flags().String("delay", "2", "The delay provided for multiple iteration")
	LaunchChaos.Flags().String("retries", "180", "The timeout provided for multiple iteration")

}
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}
