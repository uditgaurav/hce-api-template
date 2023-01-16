package apis

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/uditgaurav/hce-api-template/types"
)

const (
	VariableNotFoundError = "Don't know how to derive this tunable? Visit now, Visit now https://uditgaurav.github.io/hce-api-template//#derive-tunables"
)

func ApiToLanchExperiment(ApiDetials types.APIDetials, mode string) error {

	ApiDetials = getAPITunablesToLaunchExperiment(ApiDetials, mode)

	if ApiDetials.FileName == "" {
		ApiDetials.FileName = "hce-api.sh"
	}
	if err := validateAPITunables(ApiDetials); err != nil {
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

	if err := writeCmdToFile(ApiDetials.FileName, cmdOutput); err != nil {
		return err
	}

	return nil
}

func writeCmdToFile(fileName, cmd string) error {

	f, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer f.Close()

	_, err2 := f.WriteString(cmd)

	if err2 != nil {
		return err2
	}

	return nil
}

func getAPITunablesToLaunchExperiment(ApiDetials types.APIDetials, mode string) types.APIDetials {

	if mode == "intractive" {

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
		fmt.Print("Provide the File Name for api [Default is hce-api.sh]: ")
		fmt.Scanf("%d", &ApiDetials.FileName)
	}

	return ApiDetials
}

func validateAPITunables(ApiDetials types.APIDetials) error {

	if strings.TrimSpace(ApiDetials.HCEEndpoint) == "" {
		return errors.Errorf("HCE endpoint can't be empty, please provide a valid endpoint value")
	}
	if strings.TrimSpace(ApiDetials.ProjectID) == "" {
		return errors.Errorf("ProjectID can't be empty. %v", VariableNotFoundError)

	}
	if strings.TrimSpace(ApiDetials.WorkflowID) == "" {
		return errors.Errorf("WorkflowID can't be empty %v", VariableNotFoundError)

	}
	if strings.TrimSpace(ApiDetials.AccessKey) == "" {
		return errors.Errorf("AccessKey can't be empty %v", VariableNotFoundError)

	}
	if strings.TrimSpace(ApiDetials.AccessID) == "" {
		return errors.Errorf("AccessID can't be empty %v", VariableNotFoundError)
	}

	return nil
}

func CheckMode() string {
	var mode string

	if len(os.Args) > 0 {
		mode = "non-intractive"
	} else {
		mode = "intractive"
	}
	return mode
}
