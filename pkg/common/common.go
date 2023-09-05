package common

import (
	"fmt"
	"os"
	"strings"

	"github.com/pkg/errors"
	"github.com/uditgaurav/hce-api-template/pkg/types"
)

const (
	VariableNotFoundError = "Don't know how to derive this tunable? Visit now, Visit now https://uditgaurav.github.io/hce-api-template//#derive-tunables"
)

// WriteCmdToFile will write an api command on a given file
func WriteCmdToFile(fileName, cmd string) error {

	if strings.TrimSpace(fileName) == "" {
		fileName = "hce-api.sh"
	}

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

// ValidateAPITunables will validate the inputs used to prepare the api command
func ValidateAPITunables(APIDetails types.APIDetails) error {

	if strings.TrimSpace(APIDetails.AccoundID) == "" {
		return errors.Errorf("Account ID can't be empty, please provide a valid account id")
	}
	if strings.TrimSpace(APIDetails.ProjectID) == "" {
		return errors.Errorf("ProjectID can't be empty. %v", VariableNotFoundError)

	}
	if strings.TrimSpace(APIDetails.WorkflowID) == "" {
		return errors.Errorf("WorkflowID can't be empty %v", VariableNotFoundError)
	}
	if strings.TrimSpace(APIDetails.ApiKey) == "" {
		return errors.Errorf("AccessKey can't be empty %v", VariableNotFoundError)

	}

	return nil
}

// GetAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func GetAPITunablesForExperimentExecution(APIDetails types.APIDetails) types.APIDetails {

	fmt.Print("Provide the HCE AccoundID: ")
	fmt.Scanf("%s", &APIDetails.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &APIDetails.ProjectID)
	fmt.Print("Provide the Workflow ID: ")
	fmt.Scanf("%s", &APIDetails.WorkflowID)
	fmt.Print("Provide the HCE ApiKey: ")
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

// CheckMode will derive the mode user has selected
func CheckMode() string {
	var mode string

	if len(os.Args) > 1 {
		mode = "non-intractive"
	} else {
		mode = "intractive"
	}
	return mode
}

func CheckFile(APIDetails types.APIDetails) (string, error) {

	if APIDetails.FileName == "" {
		APIDetails.FileName = "hce-api.sh"
	}
	_, err := os.Stat(APIDetails.FileName)

	if os.IsNotExist(err) {
		_, err := os.Create(APIDetails.FileName)
		if err != nil {
			return APIDetails.FileName, err
		}
	}
	return APIDetails.FileName, nil
}
