# API Template For Pipelines

Welcome to HCE's GraphQL API template documentation

This contains the APIs to perform a chaos experiment from an automated pipeline

## Table Of Containt

1. [API to Lauch Chaos Experiment](https://uditgaurav.github.io/cv/#api-to-launch-chaos-experiment)
2. API to Monitor Chaos Experiment
3. API to Validate Resilience Score

## Pre-requisites

- **Installed HCE**: You should have a HCE installed from where you can prepare and run experiments from API calls.

- **Prepare Workflow**: We can run a workflow using APIs in any automated pipeline but for that we need to pre create a workflow with right tunables and attributes from HCE. This step will involve:
  - Select right set of experiments you want to include in the workflow
  - Provide desired tunables in the experiment. At any point of time, you can change the tunables and save it - this won't impact the overall api calls, infact this is the to update the tunables if you want to do so.


## API to Launch Chaos Experiment

This contains the API to trigger the Chaos Experiment.

### Tunables 
- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

The above tunables are mandatory to provide. You need to replace it in the given API call.

<h3><u>Looking for details on ACCESS_KEY and ACCESS_ID?</u></h3>

<li> On HCE dashboard click on "Settings" and switch to "Access Key" tab.

You will get this screen:

<img alt="settings-image" src="https://user-images.githubusercontent.com/35391335/212264846-3ea0401c-5ab7-4da5-bdb6-8559e1cb9712.png">

<li> Click on "Create Access Key" if you have lost the older one

<br>


<h3><u>Looking for details on PROJECT_ID and WORKFLOW_ID?</u></h3>
<br>

For Project ID:

<li> On HCE dashboard click on "Project" (as shown in point-1 on the image) and copy the "Project ID". You can also get the project ID from the URL.

Checkout this screen:

<img alt="projectid-img" src="https://user-images.githubusercontent.com/35391335/212269753-9023352a-cf21-49df-9097-de4c23ae3766.png">

For Workflow ID:

<li> Click on "three dots" on the workflow and Navigate to "View Manifest" option. You will get a screen like this:
<img alt="workflow-id-img" src="https://user-images.githubusercontent.com/35391335/212271135-b1e7999e-4c12-409c-80a0-0978610aacbb.png">
It will give you the workflow id for the target workflow.

<br><br>
<code>
curl '[HCE_ENDPOINT]/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"[ACCESS_ID]","access_key":"[ACCESS_KEY]"}' [HCE_ENDPOINT]/auth/login/ctl | grep -o '"access_token":"[^"]*' | cut -d'"' -f4)" -H 'Origin: [HCE_ENDPOINT]/api/' --data-binary '{"query":"mutation reRunChaosWorkFlow($workflowID: String!, $projectID: String!) {reRunChaosWorkFlow(workflowID: $workflowID, projectID: $projectID)}","variables":{"workflowID":"[WORKFLOW_ID]","projectID":"[PROJECT_ID]"}}' --compressed
</code>

<br><br>


Replace the tunables (along with `[]`) in the above query template to make it usable. For any issues refer the [HCE docs](https://developer.harness.io/docs/chaos-engineering).
