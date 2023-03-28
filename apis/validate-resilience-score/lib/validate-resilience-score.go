package lib

import (
	"fmt"
	"strings"

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

	cmd := exec.Command("bash", apiDetials.FileName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	fmt.Println(strings.TrimSpace(string(output)))
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

	cmdOutput := fmt.Sprintf(`
		curl -s --location 'https://app.harness.io/gateway/chaos/manager/api/query?accountIdentifier=%v' \
		--header 'x-api-key: %v' \
		--header 'Content-Type: application/json' \
		--data '{"query":"query ListWorkflowRun(\n  $identifiers: IdentifiersRequest!,\n  $request: ListWorkflowRunRequest!\n) {\n  listWorkflowRun(\n    identifiers: $identifiers,\n    request: $request\n  ) {\n    totalNoOfWorkflowRuns\n    workflowRuns {\n      identifiers {\n          orgIdentifier\n          projectIdentifier\n          accountIdentifier\n      }\n      workflowRunID\n      workflowID\n      weightages {\n        experimentName\n        weightage\n      }\n      updatedAt\n      createdAt\n      infra {\n        infraID\n        infraNamespace\n        infraScope\n        isActive\n        isInfraConfirmed\n      }\n      workflowName\n      workflowManifest\n      phase\n      resiliencyScore\n      experimentsPassed\n      experimentsFailed\n      experimentsAwaited\n      experimentsStopped\n      experimentsNa\n      totalExperiments\n      executionData\n      isRemoved\n      updatedBy {\n        userID\n        username\n      }\n      createdBy {\n        username\n        userID\n      }\n    }\n  }\n}","variables":{"identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"},"request":{"notifyIDs":["%v"]}}}' --compressed | jq -r '.data.listWorkflowRun.workflowRuns[0].executionData' |jq -r '.nodes'|  jq 'map(select(has("chaosData"))) | .[].chaosData.probeSuccessPercentage'`,ApiDetials.AccoundID, ApiDetials.ApiKey, ApiDetials.AccoundID, ApiDetials.ProjectID, ApiDetials.NotifyID)

	if err := common.WriteCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}
	return nil
}

// getAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func getAPITunablesForExperimentExecution(ApiDetials types.APIDetials) types.APIDetials {

	fmt.Print("Provide the AccoundID: ")
	fmt.Scanf("%s", &ApiDetials.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &ApiDetials.ProjectID)
	fmt.Print("Provide the NotifyID ID: ")
	fmt.Scanf("%s", &ApiDetials.NotifyID)
	fmt.Print("Provide the HCE Api Key: ")
	fmt.Scanf("%s", &ApiDetials.ApiKey)
	fmt.Print("Provide the File Name for API [Default is hce-api.sh]: ")
	fmt.Scanf("%s", &ApiDetials.FileName)
	fmt.Print("Provide the Expected Resilience Score: ")
	fmt.Scanf("%d", &ApiDetials.ExpectedResilienceScore)

	return ApiDetials
}
