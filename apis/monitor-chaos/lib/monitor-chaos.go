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

	return retry.
		Times(uint(timeout / delay)).
		Wait(time.Duration(delay) * time.Second).
		Try(func(attempt uint) error {

			cmd := exec.Command("bash", "-c", "./%v", apiDetials.FileName)
			stdout, err := cmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return err
			}

			if strings.TrimSpace(string(stdout)) != "Succeeded" {
				return errors.Errorf("Waiting for experiment completion... CurrentState: %v", string(stdout))
			}
			fmt.Printf("Experiment completed, CurrentState: %v", string(stdout))
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

	cmdOutput := fmt.Sprintf(`curl '%v/api/query' \
	-H 'Accept-Encoding: gzip, deflate, br' \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-H 'Connection: keep-alive' \
	-H 'DNT: 1' \
	-H "Authorization: $(curl -s -H "Content-Type: application/json" \
	-d '{"access_id":"%v","access_key":"%v"}' %v/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" \
	-H 'Origin: %v' \
	--data-binary '{"query":"query ( $request: ListWorkflowRunsRequest!) {\n listWorkflowRuns( request: $request) {\n  totalNoOfWorkflowRuns\n  workflowRuns {\n   workflowID\n   phase\n   executionData\n  } \n }\n}","variables":{"request":{"projectID":"%v","workflowIDs":["%v"]}}}' \
	--compressed | jq -r '.data.listWorkflowRuns.workflowRuns[0].phase'`, ApiDetials.HCEEndpoint, ApiDetials.AccessID, ApiDetials.AccessKey, ApiDetials.HCEEndpoint, ApiDetials.HCEEndpoint, ApiDetials.ProjectID, ApiDetials.WorkflowID)

	if err := common.WriteCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}
	fmt.Println("The file containing the API command is created successfully")

	return nil
}

// getAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func getAPITunablesForExperimentExecution(ApiDetials types.APIDetials) types.APIDetials {

	fmt.Print("Provide the HCE endpoint: ")
	fmt.Scanf("%s", &ApiDetials.HCEEndpoint)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &ApiDetials.ProjectID)
	fmt.Print("Provide the Workflow ID: ")
	fmt.Scanf("%s", &ApiDetials.WorkflowID)
	fmt.Print("Provide the HCE Access Key: ")
	fmt.Scanf("%s", &ApiDetials.AccessKey)
	fmt.Print("Provide the HCE Access ID: ")
	fmt.Scanf("%s", &ApiDetials.AccessID)
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
