package main

import (
	"fmt"
	"os"

	"github.com/uditgaurav/hce-api-template/apis"
	"github.com/uditgaurav/hce-api-template/cmd"
	"github.com/uditgaurav/hce-api-template/types"
)

// Main function
func main() {

	var api int
	apiDetials := types.APIDetials{}

	mode := apis.CheckMode()

	switch mode {
	case "non-intractive":
		cmd.Execute()
	default:

		options := `
		Provide the index number to create a file with the API command from the given options.
		For Example to create API that launches experiment give 1 as input.
	
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

			if err := apis.ApiToLanchExperiment(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to launch experiment, err: %v,", err)
				os.Exit(1)
			}

		case 2:
			if err := apis.ApiToMonitorExperiment(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to monitor experiment, err: %v,", err)
				os.Exit(1)
			}
		case 3:
			if err := apis.ApiToValidateResilienceScore(apiDetials, mode); err != nil {
				fmt.Printf("fail to create template file with API to monitor experiment, err: %v,", err)
				os.Exit(1)
			}

		default:

			fmt.Println("Not a valid option, please select from the given opetion")
			os.Exit(0)
		}
	}
}
