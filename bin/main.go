package main

import (
	"github.com/uditgaurav/hce-api-template/cmd"
	"github.com/uditgaurav/hce-api-template/pkg/common"
)

// Main function
func main() {

	mode := common.CheckMode()

	switch mode {
	case "non-intractive":
		cmd.Execute()
	default:
		cmd.ExecuteIntractive()
	}
}
