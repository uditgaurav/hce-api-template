package lib

import (
	"fmt"
	"strings"

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

	cmd := exec.Command("bash", "-c", "./%v", apiDetials.FileName)
	stdout, err := cmd.Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	fmt.Println(strings.TrimSpace(string(stdout)))
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
		`curl '%v/api/query' \
	-H 'Accept-Encoding: gzip, deflate, br' \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-H 'Connection: keep-alive' \
	-H 'DNT: 1' \
	-H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"%v","access_key":"%v"}' %v/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" \
	-H 'Origin: %v/api/' \
	--data-binary '{"query":"mutation reRunChaosWorkFlow($workflowID: String!, $projectID: String!) {reRunChaosWorkFlow(workflowID: $workflowID, projectID: $projectID)}","variables":{"workflowID":"%v","projectID":"%v"}}' --compressed`, ApiDetials.HCEEndpoint, ApiDetials.AccessID, ApiDetials.AccessKey, ApiDetials.HCEEndpoint, ApiDetials.HCEEndpoint, ApiDetials.WorkflowID, ApiDetials.ProjectID)

	if err := common.WriteCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}

	fmt.Println("The file containing the API command is created successfully")

	return nil
}
