package lib

import (
	"fmt"
	"strconv"

	"os/exec"

	"github.com/pkg/errors"
	"github.com/uditgaurav/hce-api-template/pkg/common"
	"github.com/uditgaurav/hce-api-template/pkg/types"
)

// ValidateResilienceScore will prepare the api command to get the resilience score and validate it with expected score.
func ValidateResilienceScore(apiDetials types.APIDetials, mode string) error {

	if err := PrepareResilienceScoreCMD(apiDetials, mode); err != nil {
		return errors.Errorf("fail to create template file with API to validate resilience score for a experiment, err: %v,", err)
	}

	cmd := exec.Command("bash", "-c", "./%v", apiDetials.FileName)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	output, _ := strconv.Atoi(string(stdout))

	if output != apiDetials.ExpectedResilienceScore {
		return errors.Errorf("resilience score check is failed, err: %v", err)
	}
	return nil
}

// PrepareResilienceScoreCMD will prepare a command to get the RR of a workflow
func PrepareResilienceScoreCMD(ApiDetials types.APIDetials, mode string) error {

	if mode == "intractive" {
		ApiDetials = getAPITunablesForExperimentExecution(ApiDetials)
	}
	if err := common.ValidateAPITunables(ApiDetials); err != nil {
		return err
	}

	// cmdOutput := fmt.Sprintf(
	// 	`curl '%v/api/query' \
	// 	-H 'Accept-Encoding: gzip, deflate, br' \
	// 	-H 'Content-Type: application/json' \
	// 	-H 'Accept: application/json' \
	// 	-H 'Connection: keep-alive' \
	// 	-H 'DNT: 1' \
	// 	-H  "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"%v","access_key":"%v"}' %v/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)"  \
	// 	-H 'Origin: %v' \
	// 	--data-binary '{"query":"query ( $request: ListWorkflowRunsRequest!) {\n  listWorkflowRuns( request: $request) {\n  totalNoOfWorkflowRuns\n  workflowRuns {\n    workflowID\n   phase\n   executionData\n  }\n }\n}","variables":{"request":{"projectID":"%v","workflowIDs":["%v"]}}}' --compressed \
	// 	| jq -r '.data.listWorkflowRuns.workflowRuns[0].executionData' |jq -r '.nodes'|  jq 'map(select(has("chaosData"))) | .[].chaosData.probeSuccessPercentage'
	// 	`, ApiDetials.HCEEndpoint, ApiDetials.AccessID, ApiDetials.AccessKey, ApiDetials.HCEEndpoint, ApiDetials.HCEEndpoint, ApiDetials.ProjectID, ApiDetials.WorkflowID)

	cmdOutput := fmt.Sprintf(`curl %v`, ApiDetials.AccoundID)

	if err := common.WriteCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}
	fmt.Println("The file containing the API command is created successfully")

	return nil
}

// getAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func getAPITunablesForExperimentExecution(ApiDetials types.APIDetials) types.APIDetials {

	fmt.Print("Provide the AccoundID: ")
	fmt.Scanf("%s", &ApiDetials.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &ApiDetials.ProjectID)
	fmt.Print("Provide the Workflow ID: ")
	fmt.Scanf("%s", &ApiDetials.WorkflowID)
	fmt.Print("Provide the HCE Api Key: ")
	fmt.Scanf("%s", &ApiDetials.ApiKey)
	fmt.Print("Provide the File Name for API [Default is hce-api.sh]: ")
	fmt.Scanf("%s", &ApiDetials.FileName)
	fmt.Print("Provide the Expected Resilience Score: ")
	fmt.Scanf("%d", &ApiDetials.ExpectedResilienceScore)

	return ApiDetials
}
