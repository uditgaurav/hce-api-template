package lib

import (
	"fmt"
	"os"
	"strings"

	"os/exec"

	"github.com/pkg/errors"
	"github.com/uditgaurav/hce-api-template/pkg/common"
	"github.com/uditgaurav/hce-api-template/pkg/types"
)

// ValidateResilienceScore will prepare the API command to get the resilience score and validate it with the expected score.
func ValidateResilienceScore(apiDetails types.APIDetails, mode string) error {

	var err error
	apiDetails.FileName, err = common.CheckFile(apiDetails)
	if err != nil {
		return err
	}
	// Stage 1: Prepare Resilience Score Command
	if err := PrepareResilienceScoreCMD(apiDetails, mode); err != nil {
		return errors.Errorf("Stage 1: Failed to create template file with API to validate resilience score for an experiment. Error: %v", err)
	}

	// Stage 2: Execute the Command
	cmd := exec.Command("bash", apiDetails.FileName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Stage 2: Command execution failed. Error: %v, Output: %s\n", err, output)
		return err
	}

	// Stage 3: Write Output to File
	if apiDetails.Output != "" {

		if err := os.WriteFile(apiDetails.Output, []byte(GetResilienceScoreOutput(apiDetails)), 0644); err != nil {
			fmt.Printf("Stage 3: Failed to write initial output to file. Error: %v\n", err)
			return err
		}

		// Stage 4: Execute the Output File
		cmd := exec.Command("bash", apiDetails.Output)
		completeOutput, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Printf("Stage 4: Failed to execute output file. Error: %v, Output: %s\n", err, completeOutput)
			return err
		}

		// Stage 5: Write Complete Output to File
		if err := os.WriteFile(apiDetails.Output, []byte(completeOutput), 0644); err != nil {
			fmt.Printf("Stage 5: Failed to write complete output to file. Error: %v\n", err)
			return err
		}
	}

	fmt.Println(strings.TrimSpace(string(output)))
	return nil
}

// PrepareResilienceScoreCMD will prepare a command to get the RR of a workflow
func PrepareResilienceScoreCMD(APIDetails types.APIDetails, mode string) error {

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
		--data '{"query":"query ListWorkflowRun(\n  $identifiers: IdentifiersRequest!,\n  $request: ListWorkflowRunRequest!\n) {\n  listWorkflowRun(\n    identifiers: $identifiers,\n    request: $request\n  ) {\n    totalNoOfWorkflowRuns\n    workflowRuns {\n      identifiers {\n          orgIdentifier\n          projectIdentifier\n          accountIdentifier\n      }\n      workflowRunID\n      workflowID\n      weightages {\n        experimentName\n        weightage\n      }\n      updatedAt\n      createdAt\n      infra {\n        infraID\n        infraNamespace\n        infraScope\n        isActive\n        isInfraConfirmed\n      }\n      workflowName\n      workflowManifest\n      phase\n      resiliencyScore\n      experimentsPassed\n      experimentsFailed\n      experimentsAwaited\n      experimentsStopped\n      experimentsNa\n      totalExperiments\n      executionData\n      isRemoved\n      updatedBy {\n        userID\n        username\n      }\n      createdBy {\n        username\n        userID\n      }\n    }\n  }\n}","variables":{"identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"},"request":{"notifyIDs":["%v"]}}}' --compressed | jq -r '.data.listWorkflowRun.workflowRuns[0].executionData' |jq -r '.nodes'|  jq 'map(select(has("chaosData"))) | .[].chaosData.probeSuccessPercentage'`, APIDetails.AccoundID, APIDetails.ApiKey, APIDetails.AccoundID, APIDetails.ProjectID, APIDetails.NotifyID)

	if err := common.WriteCmdToFile(APIDetails.FileName, cmdOutput); err != nil {
		return err
	}
	return nil
}

// GetResilienceScoreoutput will prepare a command to get the RR of a workflow
func GetResilienceScoreOutput(APIDetails types.APIDetails) string {

	cmdOutput := fmt.Sprintf(`
		curl -s --location 'https://app.harness.io/gateway/chaos/manager/api/query?accountIdentifier=%v' \
		--header 'x-api-key: %v' \
		--header 'Content-Type: application/json' \
		--data '{"query":"query ListWorkflowRun(\n  $identifiers: IdentifiersRequest!,\n  $request: ListWorkflowRunRequest!\n) {\n  listWorkflowRun(\n    identifiers: $identifiers,\n    request: $request\n  ) {\n    totalNoOfWorkflowRuns\n    workflowRuns {\n      identifiers {\n          orgIdentifier\n          projectIdentifier\n          accountIdentifier\n      }\n      workflowRunID\n      workflowID\n      weightages {\n        experimentName\n        weightage\n      }\n      updatedAt\n      createdAt\n      infra {\n        infraID\n        infraNamespace\n        infraScope\n        isActive\n        isInfraConfirmed\n      }\n      workflowName\n      workflowManifest\n      phase\n      resiliencyScore\n      experimentsPassed\n      experimentsFailed\n      experimentsAwaited\n      experimentsStopped\n      experimentsNa\n      totalExperiments\n      executionData\n      isRemoved\n      updatedBy {\n        userID\n        username\n      }\n      createdBy {\n        username\n        userID\n      }\n    }\n  }\n}","variables":{"identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"},"request":{"notifyIDs":["%v"]}}}' --compressed`, APIDetails.AccoundID, APIDetails.ApiKey, APIDetails.AccoundID, APIDetails.ProjectID, APIDetails.NotifyID)

	return cmdOutput
}

// getAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func getAPITunablesForExperimentExecution(APIDetails types.APIDetails) types.APIDetails {

	fmt.Print("Provide the AccoundID: ")
	fmt.Scanf("%s", &APIDetails.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &APIDetails.ProjectID)
	fmt.Print("Provide the NotifyID ID: ")
	fmt.Scanf("%s", &APIDetails.NotifyID)
	fmt.Print("Provide the HCE Api Key: ")
	fmt.Scanf("%s", &APIDetails.ApiKey)
	fmt.Print("Provide the File Name for API [Default is hce-api.sh]: ")
	fmt.Scanf("%s", &APIDetails.FileName)
	fmt.Print("Provide the Expected Resilience Score: ")
	fmt.Scanf("%d", &APIDetails.ExpectedResilienceScore)

	return APIDetails
}
