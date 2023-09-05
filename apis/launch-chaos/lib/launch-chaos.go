package lib

import (
	"fmt"
	"os"

	"os/exec"

	"github.com/pkg/errors"
	"github.com/uditgaurav/hce-api-template/pkg/common"
	"github.com/uditgaurav/hce-api-template/pkg/types"
)

// LaunchChaos will prepare the api command to get re-run a given workflow
func LaunchChaos(APIDetails types.APIDetails, mode string) error {

	var err error
	APIDetails.FileName, err = common.CheckFile(APIDetails)
	if err != nil {
		return err
	}
	if err := ApiToLanchExperiment(APIDetails, mode); err != nil {
		return errors.Errorf("fail to create template file with API to launch chaos experiment, err: %v,", err)
	}

	err = os.Chmod(APIDetails.FileName, 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("bash", APIDetails.FileName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}

	if APIDetails.Output != "" {
		if err := os.WriteFile(APIDetails.Output, output, 0644); err != nil {
			fmt.Println("Failed to write to output file:", err)
			return err
		}
	} else {
		fmt.Println(string(output))
	}

	return nil
}

// ApiToLanchExperiment will prepare api command to get the workflow status
func ApiToLanchExperiment(APIDetails types.APIDetails, mode string) error {

	if mode == "intractive" {
		APIDetails = common.GetAPITunablesForExperimentExecution(APIDetails)
	}
	if err := common.ValidateAPITunables(APIDetails); err != nil {
		return err
	}

	cmdOutput := fmt.Sprintf(
		`
	curl -s --location 'https://app.harness.io/gateway/chaos/manager/api/query?accountIdentifier=%v' \
    --header 'x-api-key: %v' \
    --header 'Content-Type: application/json' \
    --data '{"query":"mutation RunChaosExperiment(\n  $workflowID: String!,\n  $identifiers: IdentifiersRequest!\n) {\n  runChaosExperiment(\n    workflowID: $workflowID,\n    identifiers: $identifiers\n  ) {\n    notifyID\n  }\n}","variables":{"workflowID":"%v","identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"}}}' --compressed`, APIDetails.AccoundID, APIDetails.ApiKey, APIDetails.WorkflowID, APIDetails.AccoundID, APIDetails.ProjectID)
	if err := common.WriteCmdToFile(APIDetails.FileName, cmdOutput); err != nil {
		return err
	}
	return nil
}
