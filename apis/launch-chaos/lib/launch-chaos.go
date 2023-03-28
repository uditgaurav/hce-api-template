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
func LaunchChaos(apiDetials types.APIDetials, mode string) error {

	if err := ApiToLanchExperiment(apiDetials, mode); err != nil {
		return errors.Errorf("fail to create template file with API to launch chaos experiment, err: %v,", err)
	}

	err := os.Chmod(apiDetials.FileName, 0755)
	if err != nil {
		return err
	}

	cmd := exec.Command("bash", apiDetials.FileName)

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	
	fmt.Println(string(output))
	return nil
}

// ApiToLanchExperiment will prepare api command to get the workflow status
func ApiToLanchExperiment(ApiDetials types.APIDetials, mode string) error {

	if mode == "intractive" {
		ApiDetials = common.GetAPITunablesForExperimentExecution(ApiDetials)
	}
	if err := common.ValidateAPITunables(ApiDetials); err != nil {
		return err
	}

	cmdOutput := fmt.Sprintf(
		`
	curl -s --location 'https://app.harness.io/gateway/chaos/manager/api/query?accountIdentifier=%v' \
    --header 'x-api-key: %v' \
    --header 'Content-Type: application/json' \
    --data '{"query":"mutation RunChaosExperiment(\n  $workflowID: String!,\n  $identifiers: IdentifiersRequest!\n) {\n  runChaosExperiment(\n    workflowID: $workflowID,\n    identifiers: $identifiers\n  ) {\n    notifyID\n  }\n}","variables":{"workflowID":"%v","identifiers":{"orgIdentifier":"default","accountIdentifier":"%v","projectIdentifier":"%v"}}}' --compressed`, ApiDetials.AccoundID, ApiDetials.ApiKey, ApiDetials.WorkflowID, ApiDetials.AccoundID, ApiDetials.ProjectID)
	if err := common.WriteCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}
	return nil
}
