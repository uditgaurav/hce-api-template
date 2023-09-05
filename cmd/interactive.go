package cmd

import (
	"fmt"
	"os"

	launchChaos "github.com/uditgaurav/hce-api-template/apis/launch-chaos/lib"
	monitorChaos "github.com/uditgaurav/hce-api-template/apis/monitor-chaos/lib"
	validateRR "github.com/uditgaurav/hce-api-template/apis/validate-resilience-score/lib"

	"github.com/uditgaurav/hce-api-template/pkg/types"
)

// ExecuteIntractive will run the command in intractive mode
func ExecuteIntractive() {
	var api int
	APIDetails := types.APIDetails{}

	options := `
	Provide the index number to create a file with the API command from the given options.
	For Example to create an API that launches an experiment give 1 as input.

	[1] API to Launch Chaos Experiment
	[2] API to Monitor Chaos Experiment
	[3] API to Validate Resilience Score

	`
	fmt.Println(options)
	fmt.Print("Select from the given options: ")
	fmt.Scanf("%d", &api)

	fmt.Print("\nTo know more about any of the tunables visit https://uditgaurav.github.io/hce-api-template//#derive-tunables\n\n")

	switch api {

	case 1:

		if err := launchChaos.LaunchChaos(APIDetails, "intractive"); err != nil {
			fmt.Printf("fail to create template file with API to launch experiment, err: %v,", err)
			os.Exit(1)
		}

	case 2:
		if err := monitorChaos.MonitorChaosExperiment(APIDetails, "intractive"); err != nil {
			fmt.Printf("manitor chaos experiment failed, err: %v,", err)
			os.Exit(1)
		}
	case 3:
		if err := validateRR.ValidateResilienceScore(APIDetails, "intractive"); err != nil {
			fmt.Printf("fail to create template file with API to monitor experiment, err: %v,", err)
			os.Exit(1)
		}

	default:

		fmt.Println("Not a valid option, please select from the given opetion")
		os.Exit(0)
	}
}
