curl 'http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091//api/query' \
	-H 'Accept-Encoding: gzip, deflate, br' \
	-H 'Content-Type: application/json' \
	-H 'Accept: application/json' \
	-H 'Connection: keep-alive' \
	-H 'DNT: 1' \
	-H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"adminNCWQu","access_key":"nEdGNDDrTFHyCnl"}' http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091//auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" \
	-H 'Origin: http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091//api/' \
	--data-binary '{"query":"mutation reRunChaosWorkFlow($workflowID: String!, $projectID: String!) {reRunChaosWorkFlow(workflowID: $workflowID, projectID: $projectID)}","variables":{"workflowID":"f4581780-efaf-4155-956e-6c379f24394b","projectID":"abceb5f4-4268-4467-9818-ad6e3b6bfd78"}}' --compressed