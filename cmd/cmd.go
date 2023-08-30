package cmd

import (
	"encoding/json"
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
	Example: "./hce-api-saas generate --api launch-experiment --account-id=cTU1lRSWS2KWNV9phKvuOA --project-id ChaosTestinProd2 --workflow-id f4581780-efaf-4155-956e-6c379f24394b --api-key nEdGNDDrTFHyCnl --file-name hce-api.sh",
	Run: func(cmd *cobra.Command, args []string) {

		apiDetials := types.APIDetials{}
		mode := common.CheckMode()
		var api string

		configFile, _ := cmd.Flags().GetString("config")
		if configFile != "" {
			file, err := os.Open(configFile)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			defer file.Close()

			config := []JSONConfig{}
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&config)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)
				os.Exit(1)
			}

			// Assuming you want to use the first element in the JSON array
			if len(config) > 0 {
				apiDetials.ApiKey = config[0].ApiKey
				apiDetials.AccoundID = config[0].AccountId
				apiDetials.ProjectID = config[0].ProjectId
				apiDetials.WorkflowID = config[0].WorkflowId
				api = config[0].API
				apiDetials.FileName = config[0].FileName
				apiDetials.Output = config[0].Output
				apiDetials.NotifyID = config[0].NotifyID
			}
		} else {

			var err error
			api, err = cmd.Flags().GetString("api")
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			apiDetials.FileName, err = cmd.Flags().GetString("file-name")
			if err != nil {
				fmt.Println(err)
			}
			apiDetials.AccoundID, err = cmd.Flags().GetString("account-id")
			if err != nil {
				fmt.Println(err)
			}
			apiDetials.ProjectID, err = cmd.Flags().GetString("project-id")
			if err != nil {
				fmt.Println(err)
			}
			apiDetials.WorkflowID, err = cmd.Flags().GetString("workflow-id")
			if err != nil {
				fmt.Println(err)
			}
			apiDetials.ApiKey, err = cmd.Flags().GetString("api-key")
			if err != nil {
				fmt.Println(err)
			}
			apiDetials.API, err = cmd.Flags().GetString("api")
			if err != nil {
				fmt.Println(err)
			}
		}

		switch api {
		case "launch-experiment":

			if err := launchChaos.LaunchChaos(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to launch experiment, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)

		case "monitor-experiment":
			apiDetials.Delay, _ = cmd.Flags().GetString("delay")
			apiDetials.Timeout, _ = cmd.Flags().GetString("timeout")
			apiDetials.NotifyID, _ = cmd.Flags().GetString("notifyID")

			if err := monitorChaos.MonitorChaosExperiment(apiDetials, mode); err != nil {
				fmt.Printf("monitor chaos failed, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)

		case "validate-resilience-score":

			// apiDetials.NotifyID, _ = cmd.Flags().GetString("notifyID")

			if err := validateRR.ValidateResilienceScore(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to validate resilience score of the workflow, err: %v,", err)
				os.Exit(1)
			}
			os.Exit(0)
		}

	},
}

var rootCmd = &cobra.Command{
	Use:   "hce-api-saas",
	Short: "CLI to generate api",
}

func init() {

	rootCmd.AddCommand(LaunchChaos)

	LaunchChaos.Flags().String("api", "", "Set the name of target api")
	LaunchChaos.Flags().String("project-id", "", "Set the hce project id")
	LaunchChaos.Flags().String("account-id", "", "Set the account id")
	LaunchChaos.Flags().String("workflow-id", "abc", "Set the workflow id")
	LaunchChaos.Flags().String("notifyID", "abc", "Set the notifyID id")
	LaunchChaos.Flags().String("api-key", "", "Set the api key")
	LaunchChaos.Flags().String("file-name", "", "The target file name which contains the API command")
	LaunchChaos.Flags().String("delay", "2", "The delay provided for multiple iteration")
	LaunchChaos.Flags().String("timeout", "180", "The timeout provided for multiple iteration")
	LaunchChaos.Flags().String("config", "", "Path to the JSON config file")

}
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

type JSONConfig struct {
	ApiKey     string `json:"apiKey"`
	AccountId  string `json:"accountId"`
	ProjectId  string `json:"projectId"`
	WorkflowId string `json:"workflowId"`
	API        string `json:"api"`
	FileName   string `json:"fileName"`
	Output     string `json:"output"`
	NotifyID   string `json:"notifyId"`
}
