package lib

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"os/exec"

	"github.com/pkg/errors"
	"github.com/uditgaurav/hce-api-template/pkg/common"
	"github.com/uditgaurav/hce-api-template/pkg/types"
	"github.com/uditgaurav/hce-api-template/pkg/utils/retry"
)

// MonitorChaosExperimenrt will prepare the api command to get the experiment status and monitor it for the timeout duration
func MonitorChaosExperiment(apiDetials types.APIDetials, mode string) error {

	if err := ApiToMonitorExperiment(apiDetials, mode); err != nil {
		return errors.Errorf("fail to create template file with API to monitor experiment, err: %v,", err)
	}

	delay, _ := strconv.Atoi(apiDetials.Delay)
	timeout, _ := strconv.Atoi(apiDetials.Timeout)

	fmt.Printf("The timeout: %v and delay: %v \n", timeout, delay)

	return retry.
		Times(uint(timeout / delay)).
		Wait(time.Duration(delay) * time.Second).
		Try(func(attempt uint) error {

			cmd := exec.Command("bash", apiDetials.FileName)

			output, err := cmd.CombinedOutput()
			if err != nil {
				fmt.Println("Error:", err)
				return err
			}
			if strings.TrimSpace(string(output)) != "Completed" {
				fmt.Printf("Waiting for experiment completion... CurrentState: %v", string(output))
				return errors.Errorf("Waiting for experiment completion... CurrentState: %v", string(output))
			}
			fmt.Printf("Experiment completed, CurrentState: %v", string(output))
			return nil
		})
}

// ApiToMonitorExperiment will prepare api command to get the workflow status
func ApiToMonitorExperiment(ApiDetials types.APIDetials, mode string) error {

	if mode == "intractive" {
		ApiDetials = getAPITunablesForExperimentExecution(ApiDetials)
	}
	if err := common.ValidateAPITunables(ApiDetials); err != nil {
		return err
	}

	cmdOutput := fmt.Sprintf(`
	curl -s --location 'https://app.harness.io/gateway/chaos/manager/api/query?accountIdentifier=%v' \
	--header 'x-api-key: %v' \
	--header 'Content-Type: application/json' \
	--data '{"query":"query ListWorkflowRun(\n  $identifiers: IdentifiersRequest!,\n  $request: ListWorkflowRunRequest!\n) {\n  listWorkflowRun(\n    identifiers: $identifiers,\n    request: $request\n  ) {\n    totalNoOfWorkflowRuns\n    workflowRuns {\n      identifiers {\n          orgIdentifier\n          projectIdentifier\n          accountIdentifier\n      }\n      workflowRunID\n      workflowID\n      weightages {\n        experimentName\n        weightage\n      }\n      updatedAt\n      createdAt\n      infra {\n        infraID\n        infraNamespace\n        infraScope\n        isActive\n        isInfraConfirmed\n      }\n      workflowName\n      workflowManifest\n      phase\n      resiliencyScore\n      experimentsPassed\n      experimentsFailed\n      experimentsAwaited\n      experimentsStopped\n      experimentsNa\n      totalExperiments\n      executionData\n      isRemoved\n      updatedBy {\n        userID\n        username\n      }\n      createdBy {\n        username\n        userID\n      }\n    }\n  }\n}","variables":{"identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"},"request":{"notifyIDs":["%v"]}}}' --compressed | jq -r '.data.listWorkflowRun.workflowRuns[0].phase'`, ApiDetials.AccoundID, ApiDetials.ApiKey, ApiDetials.AccoundID, ApiDetials.ProjectID, ApiDetials.NotifyID)
	if err := common.WriteCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}
	fmt.Println("The file containing the API command is created successfully")

	return nil
}

// getAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func getAPITunablesForExperimentExecution(ApiDetials types.APIDetials) types.APIDetials {

	fmt.Print("Provide the account id: ")
	fmt.Scanf("%s", &ApiDetials.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &ApiDetials.ProjectID)
	fmt.Print("Provide the Workflow ID: ")
	fmt.Scanf("%s", &ApiDetials.WorkflowID)
	fmt.Print("Provide the NotifyID: ")
	fmt.Scanf("%s", &ApiDetials.NotifyID)
	fmt.Print("Provide the api key: ")
	fmt.Scanf("%s", &ApiDetials.ApiKey)
	fmt.Print("Provide the File Name for API [Default is hce-api.sh]: ")
	fmt.Scanf("%s", &ApiDetials.FileName)
	fmt.Print("Provide the delay[Default 2]: ")
	fmt.Scanf("%s", &ApiDetials.Delay)
	fmt.Print("Provide the timeout [Default 180]: ")
	fmt.Scanf("%s", &ApiDetials.Timeout)

	if ApiDetials.Delay == "" {
		ApiDetials.Delay = "2"
	}
	if ApiDetials.Timeout == "" {
		ApiDetials.Timeout = "180"
	}

	return ApiDetials
}
