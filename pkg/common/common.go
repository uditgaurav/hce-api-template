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
func ValidateAPITunables(ApiDetials types.APIDetials) error {

	if strings.TrimSpace(ApiDetials.AccoundID) == "" {
		return errors.Errorf("Account ID can't be empty, please provide a valid account id")
	}
	if strings.TrimSpace(ApiDetials.ProjectID) == "" {
		return errors.Errorf("ProjectID can't be empty. %v", VariableNotFoundError)

	}
	if strings.TrimSpace(ApiDetials.NotifyID) == "" {
		return errors.Errorf("NotifyID can't be empty %v", VariableNotFoundError)

	}
	if strings.TrimSpace(ApiDetials.ApiKey) == "" {
		return errors.Errorf("AccessKey can't be empty %v", VariableNotFoundError)

	}

	return nil
}

// GetAPITunablesForExperimentExecution will get the values to prepare api command in interactive mode
func GetAPITunablesForExperimentExecution(ApiDetials types.APIDetials) types.APIDetials {

	fmt.Print("Provide the HCE AccoundID: ")
	fmt.Scanf("%s", &ApiDetials.AccoundID)
	fmt.Print("Provide the Project ID: ")
	fmt.Scanf("%s", &ApiDetials.ProjectID)
	fmt.Print("Provide the Workflow ID: ")
	fmt.Scanf("%s", &ApiDetials.WorkflowID)
	fmt.Print("Provide the HCE ApiKey: ")
	fmt.Scanf("%s", &ApiDetials.ApiKey)
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
