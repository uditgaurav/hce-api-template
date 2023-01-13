# API Template For Pipelines

Welcome to HCE's GraphQL API template documentation

This contains the APIs to perform a chaos experiment from an automated pipeline

## Table Of Content

1. [API to Launch Chaos Experiment](https://uditgaurav.github.io/cv/#api-to-launch-chaos-experiment)
2. [API to Monitor Chaos Experiment]()
3. [API to Validate Resilience Score]()

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

<h3>Looking for details on ACCESS_KEY and ACCESS_ID?</h3>

<li> On the HCE dashboard click on "Settings" and switch to the "Access Key" tab.

You will get this screen:

<img alt="settings-image" src="https://user-images.githubusercontent.com/35391335/212264846-3ea0401c-5ab7-4da5-bdb6-8559e1cb9712.png">

<li> Click on "Create Access Key" if you have lost the older one

<br>


<h3>Looking for details on PROJECT_ID and WORKFLOW_ID?</h3>
<br>

For Project ID:

<li> On the HCE dashboard click on "Project" (as shown in point-1 on the image) and copy the "Project ID". You can also get the project ID from the URL.

Checkout this screen:

<img alt="projectid-img" src="https://user-images.githubusercontent.com/35391335/212269753-9023352a-cf21-49df-9097-de4c23ae3766.png">

For Workflow ID:

<li> Click on "three dots" on the workflow and Navigate to the "View Manifest" option. You will get a screen like this:

<img alt="workflow-id-img" src="https://user-images.githubusercontent.com/35391335/212271135-b1e7999e-4c12-409c-80a0-0978610aacbb.png">

It will give you the workflow id for the target workflow.

Now use the given API call to launch chaos with all the tunables mentioned above.

<br><br>

<table>
  <tr>
    <td>
      <code>
      curl '&lt;HCE_ENDPOINT&gt;/api/query' -H 'Accept-Encoding: gzip, deflate, br' -H 'Content-Type: application/json' -H 'Accept: application/json' -H 'Connection: keep-alive' -H 'DNT: 1' -H "Authorization: $(curl -s -H "Content-Type: application/json" -d '{"access_id":"&lt;ACCESS_ID&gt;","access_key":"&lt;ACCESS_KEY&gt;"}' &lt;HCE_ENDPOINT&gt;/auth/login/ctl | grep -o '"access_token":"&lt;^"&gt;*' | cut -d'"' -f4)" -H 'Origin: &lt;HCE_ENDPOINT&gt;/api/' --data-binary '{"query":"mutation reRunChaosWorkFlow($workflowID: String!, $projectID: String!) {reRunChaosWorkFlow(workflowID: $workflowID, projectID: $projectID)}","variables":{"workflowID":"&lt;WORKFLOW_ID&gt;","projectID":"&lt;PROJECT_ID&gt;"}}' --compressed
      </code>
    </td>
  </tr>
</table>
<br><br>


Replace the tunables (along with '[]') in the above query template to make it usable. For any issues refer to the <a href="https://developer.harness.io/docs/chaos-engineering">HCE docs</a>.
