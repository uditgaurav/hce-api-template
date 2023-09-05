package lib

import (
	"bytes"
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
func MonitorChaosExperiment(APIDetails types.APIDetails, mode string) error {

	var err error
	APIDetails.FileName, err = common.CheckFile(APIDetails)
	if err != nil {
		return err
	}
	if err := ApiToMonitorExperiment(APIDetails, mode); err != nil {
		return errors.Errorf("fail to create template file with API to monitor experiment, err: %v,", err)
	}

	delay, _ := strconv.Atoi(APIDetails.Delay)
	timeout, _ := strconv.Atoi(APIDetails.Timeout)

	if delay == 0 {
		delay = 2
	}
	if timeout == 0 {
		timeout = 180
	}

	fmt.Printf("The timeout: %v and delay: %v \n", timeout, delay)

	return retry.
		Times(uint(timeout / delay)).
		Wait(time.Duration(delay) * time.Second).
		Try(func(attempt uint) error {

			var stdout, stderr bytes.Buffer
			cmd := exec.Command("bash", APIDetails.FileName)
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr

			err := cmd.Run()

			if err != nil {
				fmt.Println("Error:", stderr.String())
				return err
			}
			if strings.TrimSpace(stdout.String()) != "Completed" && strings.TrimSpace(stdout.String()) != "Stopped" {
				fmt.Printf("Waiting for experiment completion... CurrentState: %v", stdout.String())
				return errors.Errorf("Waiting for experiment completion... CurrentState: %v", stdout.String())
			}
			fmt.Printf("Experiment completed, CurrentState: %v", stdout.String())
			return nil
		})
}

// ApiToMonitorExperiment will prepare api command to get the workflow status
func ApiToMonitorExperiment(APIDetails types.APIDetails, mode string) error {

	if mode == "intractive" {
		APIDetails = getAPITunablesForExperimentExecution(APIDetails)
	}
	if err := common.ValidateAPITunables(APIDetails); err != nil {
		return err
	}

	cmdOutput := fmt.Sprintf(`
	curl -s --location 'https://app.harness.io/gateway/chaos/manager/api/query?accountIdentifier=%v' \
	--header 'x-api-key: %v' \
	--header 'Content-Type: application/json' \
	--data '{"query":"query ListWorkflowRun(\n  $identifiers: IdentifiersRequest!,\n  $request: ListWorkflowRunRequest!\n) {\n  listWorkflowRun(\n    identifiers: $identifiers,\n    request: $request\n  ) {\n    totalNoOfWorkflowRuns\n    workflowRuns {\n      identifiers {\n          orgIdentifier\n          projectIdentifier\n          accountIdentifier\n      }\n      workflowRunID\n      workflowID\n      weightages {\n        experimentName\n        weightage\n      }\n      updatedAt\n      createdAt\n      infra {\n        infraID\n        infraNamespace\n        infraScope\n        isActive\n        isInfraConfirmed\n      }\n      workflowName\n      workflowManifest\n      phase\n      resiliencyScore\n      experimentsPassed\n      experimentsFailed\n      experimentsAwaited\n      experimentsStopped\n      experimentsNa\n      totalExperiments\n      executionData\n      isRemoved\n      updatedBy {\n        userID\n        username\n      }\n      createdBy {\n        username\n        userID\n      }\n    }\n  }\n}","variables":{"identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"},"request":{"notifyIDs":["%v"]}}}' --compressed | jq -r '.data.listWorkflowRun.workflowRuns[0].phase'`, APIDetails.AccoundID, APIDetails.ApiKey, APIDetails.AccoundID, APIDetails.ProjectID, APIDetails.NotifyID)
	if err := common.WriteCmdToFile(APIDetails.FileName, cmdOutput); err != nil {
		return err
	}
	fmt.Println("The file containing the API command is created successfully")

	return nil
}

// getAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func getAPITunablesForExperimentExecution(APIDetails types.APIDetails) types.APIDetails {

	fmt.Print("Provide the account id: ")
	fmt.Scanf("%s", &APIDetails.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &APIDetails.ProjectID)
	fmt.Print("Provide the Workflow ID: ")
	fmt.Scanf("%s", &APIDetails.WorkflowID)
	fmt.Print("Provide the NotifyID: ")
	fmt.Scanf("%s", &APIDetails.NotifyID)
	fmt.Print("Provide the api key: ")
	fmt.Scanf("%s", &APIDetails.ApiKey)
	fmt.Print("Provide the File Name for API [Default is hce-api.sh]: ")
	fmt.Scanf("%s", &APIDetails.FileName)
	fmt.Print("Provide the delay[Default 2]: ")
	fmt.Scanf("%s", &APIDetails.Delay)
	fmt.Print("Provide the timeout [Default 180]: ")
	fmt.Scanf("%s", &APIDetails.Timeout)

	if APIDetails.Delay == "" {
		APIDetails.Delay = "2"
	}
	if APIDetails.Timeout == "" {
		APIDetails.Timeout = "180"
	}

	return APIDetails
}
