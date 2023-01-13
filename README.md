# API Template For Automation

Welcome to HCE's GraphQL API template documentation

This contains the APIs to perform a chaos experiment from an automated pipeline

## Table Of Content

1. [API to Launch Chaos Experiment](https://uditgaurav.github.io/hce-api-template//#api-to-launch-chaos-experiment)
2. [API to Monitor Chaos Experiment](https://uditgaurav.github.io/hce-api-template//#api-to-monitor-chaos-experiment)
3. [API to Validate Resilience Score](https://uditgaurav.github.io/hce-api-template//#api-to-validate-resilience-score)

## Pre-requisites

- **Installed HCE**: You should have an HCE installed from where you can prepare and run experiments from API calls.

- **Prepare Workflow**: We can run a workflow using APIs in any automated pipeline but for that, we need to pre-create a workflow with the right tunables and attributes from HCE. This step will involve:
  - Select the right set of experiments you want to include in the workflow
  - Provide desired tunables in the experiment. At any point in time, you can change the tunables and save it - this won't impact the overall API calls, infact this is the to update the tunables if you want to do so.


## API to Launch Chaos Experiment

This contains the API to trigger the Chaos Experiment.

### Tunables 
- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

The above tunables are mandatory to provide. You need to replace it in the given API call.

### Looking for details on ACCESS_KEY and ACCESS_ID?

- On the HCE dashboard click on "Settings" and switch to the "Access Key" tab.

You will get this screen:

![settings-image](https://user-images.githubusercontent.com/35391335/212264846-3ea0401c-5ab7-4da5-bdb6-8559e1cb9712.png)

- Click on "Create Access Key" if you have lost the older one

### Looking for details on PROJECT_ID and WORKFLOW_ID?

For Project ID:

- On the HCE dashboard click on "Project" (as shown in point-1 on the image) and copy the "Project ID". You can also get the project ID from the HCE URL at this place.

Checkout this screen:

![projectid-img](https://user-images.githubusercontent.com/35391335/212269753-9023352a-cf21-49df-9097-de4c23ae3766.png)

For Workflow ID:

- Click on "three dots" on the workflow and Navigate to the "View Manifest" option. You will get a screen like this:

![workflow-id-img](https://user-images.githubusercontent.com/35391335/212271135-b1e7999e-4c12-409c-80a0-0978610aacbb.png)

It will give you the workflow id for the target workflow.

Now use the given API call to launch chaos with all the tunables mentioned above.

```

curl '<HCE_ENDPOINT>/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: 
keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"<ACCESS_ID>","access_key":"<ACCESS_KEY>"}' 
<HCE_ENDPOINT>/auth/login/ctl | grep -o '"access_token":"<^">*' | cut -d'"' -f4)" -H 'Origin: <HCE_ENDPOINT>/api/' --data-binary '{"query":"mutation 
reRunChaosWorkFlow($workflowID: String!, $projectID: String!) {reRunChaosWorkFlow(workflowID: $workflowID, projectID: $projectID)}","variables":
{"workflowID":"<WORKFLOW_ID>","projectID":"<PROJECT_ID>"}}' --compressed

```

Replace the tunables (along with '[]') in the above query template to make it usable. For any issues refer to the [HCE docs](https://developer.harness.io/docs/chaos-engineering)

## API to Monitor Chaos Experiment

This contains the API to monitor the Chaos Experiment that is this API will help us to wait for the workflow completion.

### Tunables 

- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

- Please refer [step 1](https://uditgaurav.github.io/hce-api-template//#api-to-launch-chaos-experiment) to know how can we get the values of different tunables.

```

curl '<HCE_ENDPOINT>/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: 
keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"<ACCESS_ID>","access_key":"<ACCESS_KEY>"}' 
<HCE_ENDPOINT>/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" -H 'Origin: <HCE_ENDPOINT>' --data-binary '{"query":"query ( $request: 
ListWorkflowRunsRequest!) {\n listWorkflowRuns( request: $request) {\n  totalNoOfWorkflowRuns\n  workflowRuns {\n   workflowID\n   phase\n   executionData\n  }
\n }\n}","variables":{"request":{"projectID":"<PROJECT_ID>","workflowIDs":["<WORKFLOW_ID>"]}}}' --compressed | jq -r '.data.listWorkflowRuns.workflowRuns[0].phase'

```

Replace the tunables (along with '[]') in the above query template to make it usable.

#### A sample shell script to monitor Chaos Experiment

- In this sample script we will wait for the workflow completion with the delay of 2 seconds and 150 retries, you can adjust these values based on total chaos duration.

```bash
#!/bin/sh

$delay=2
$retry=150

for i in range {0..$retry}; do
    res=$(curl '<HCE_ENDPOINT>/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"<ACCESS_ID>","access_key":"<ACCESS_KEY>"}' <HCE_ENDPOINT>/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" -H 'Origin: <HCE_ENDPOINT>' --data-binary '{"query":"query ( $request: ListWorkflowRunsRequest!) {\n listWorkflowRuns( request: $request) {\n  totalNoOfWorkflowRuns\n  workflowRuns {\n   workflowID\n   phase\n   executionData\n  }\n }\n}","variables":{"request":{"projectID":"<PROJECT_ID>","workflowIDs":["<WORKFLOW_ID>"]}}}' --compressed | jq -r '.data.listWorkflowRuns.workflowRuns[0].phase')
    if [ "$res" == "Succeeded" ]; then
        echo "Experiment completed, CurrentState: $res"
        exit 0
    fi
    sleep $delay
    echo "Waiting for experiment completion... CurrentState: $res"
done

echo "[Error]: Timeout the workflows is not completed with delay: $delay and retry: $retry, CurrentState: $res"

exit 1
```

(Replace the tunables (along with '[]') in the above query template to make it usable)

## API to Validate Resilience Score

This contains the API to get the resilience score for a workflow run and validate against the expected probe success percentage.

### Tunables 

- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

- Please refer [step 1](https://uditgaurav.github.io/hce-api-template//#api-to-launch-chaos-experiment) to know how can we get the values of different tunables.

```

curl '<HCE_ENDPOINT>/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: 
keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"<ACCESS_ID>","access_key":"<ACCESS_KEY>"}' 
<HCE_ENDPOINT>/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" -H 'Origin: <HCE_ENDPOINT>' --data-binary '{"query":"query ( $request: 
ListWorkflowRunsRequest!) {\n listWorkflowRuns( request: $request) {\n  totalNoOfWorkflowRuns\n  workflowRuns {\n   workflowID\n   phase\n   executionData\n  }
\n }\n}","variables":{"request":{"projectID":"<PROJECT_ID>","workflowIDs":["<WORKFLOW_ID>"]}}}' --compressed | jq -r '.data.listWorkflowRuns.workflowRuns[0].
executionData' |jq -r '.nodes'|  jq 'map(select(has("chaosData"))) | .[].chaosData.probeSuccessPercentage'

```

(Replace the tunables (along with '[]') in the above query template to make it usable)

#### A sample shell script to validate resiliency score

- In this sample script you will get the probe success percentage for the last workflow run, you can make use of it to compare with the expected probe success percentage.

```bash
#!/bin/sh

$expectedProbeSuccessPercentage=100

res=$(curl '<HCE_ENDPOINT>/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"<ACCESS_ID>","access_key":"<ACCESS_KEY>"}' <HCE_ENDPOINT>/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" -H 'Origin: <HCE_ENDPOINT>' --data-binary '{"query":"query ( $request: ListWorkflowRunsRequest!) {\n listWorkflowRuns( request: $request) {\n  totalNoOfWorkflowRuns\n  workflowRuns {\n   workflowID\n   phase\n   executionData\n  }\n }\n}","variables":{"request":{"projectID":"<PROJECT_ID>","workflowIDs":["<WORKFLOW_ID>"]}}}' --compressed | jq -r '.data.listWorkflowRuns.workflowRuns[0].executionData' |jq -r '.nodes'|  jq 'map(select(has("chaosData"))) | .[].chaosData.probeSuccessPercentage')
if [ "$res" != "$expectedProbeSuccessPercentage" ]; then
    echo "The probe success percentage is: $res, expected probe sucess percentage: $expectedProbeSuccessPercentage"
    exit 1
fi

echo "The probe success percentage is equal to expected probe success percentage"
exit 0
```
(Replace the tunables (along with '[]') in the above query template to make it usable)
