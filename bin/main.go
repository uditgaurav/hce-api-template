package main

import (
	"fmt"
	"os"

	"github.com/chaosnative/hce-api-template/apis"
	"github.com/uditgaurav/hce-api-template/apis"
)

// Main function
func main() {

	var api int

	options := `
	Provide the index number to create a file with the API command from the given options.
	For Example to create API that launch experiment give 1 as input.

	[1] API to Launch Chaos Experiment
	[2] API to Monitor Chaos Experiment
	[3] API to Validate Resilience Score

	`
	fmt.Println(options)
	fmt.Scanf("Select from the given options %d", &api)

	switch api {

	case 1:

	if err := apis.ApiToLanchExperiment(); err !=nil {
		fmt.Errorf("fail to create template file with API to launch experiment, err: %v",err)
	}

	case 2:
	case 3:

	default:

		fmt.Println("Not a valid option, please select from the given opetion")
		os.Exit(0)
		
	}

}
