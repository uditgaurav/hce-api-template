package apis

import (
	"fmt"
	"os"
)

func ApiToLanchExperiment()error{

	_, err := os.Getwd()
	if err != nil {
		return err
	}

	cmd:= `curl '<HCE_ENDPOINT>/api/query' \
	-H 'Accept-Encoding: gzip, deflate, br' \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-H 'Connection: keep-alive' \
	-H 'DNT: 1' \
	-H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"<ACCESS_ID>","access_key":"<ACCESS_KEY>"}' \
	<HCE_ENDPOINT>/auth/login/ctl | grep -o '"access_token":"<^">*' | cut -d'"' -f4)" \
	-H 'Origin: <HCE_ENDPOINT>/api/' \
	--data-binary '{"query":"mutation reRunChaosWorkFlow($workflowID: String!, $projectID: String!) {reRunChaosWorkFlow(workflowID: $workflowID, projectID: $projectID)}","variables":{"workflowID":"<WORKFLOW_ID>","projectID":"<PROJECT_ID>"}}' --compressed`

	fmt.Println(cmd)


	return nil
}