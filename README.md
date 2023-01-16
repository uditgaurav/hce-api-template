# API Template For Automation

Welcome to HCE's GraphQL API template documentation

This contains the APIs templates to perform a chaos experiment in an automated way.

## Table Of Content

1. [API to Launch Chaos Experiment](https://uditgaurav.github.io/hce-api-template//#api-to-launch-chaos-experiment)
2. [API to Monitor Chaos Experiment](https://uditgaurav.github.io/hce-api-template//#api-to-monitor-chaos-experiment)
3. [API to Validate Resilience Score](https://uditgaurav.github.io/hce-api-template//#api-to-validate-resilience-score)

## Pre-requisites

- **Installed HCE**: You should have an HCE installed from where you can prepare and run experiments from API calls.

- **Prepare Workflow**: We can run a workflow using APIs in any automated pipeline but for that, we need to pre-create a workflow with the right tunables and attributes from HCE. This step will involve:
  - Select the right set of experiments you want to include in the workflow
  - Provide desired tunables in the experiment. At any point in time, you can change the tunables and save it - this won't impact the overall API calls, infact this is the to update the tunables if you want to do so.

- **Download hce-api**: Follow the below mentioned steps to prepare API that launches chaos experiment.

- Download `hce-api` binary using this command, replace `<ACRH>` with you system arch (supports `amd64` and `arm64`).

```bash
wget https://github.com/uditgaurav/hce-api-template/releases/download/0.1.0/hce-api-<ARCH> -O hce-api
chmod +x hce-api
```

If you have don't have `wget` then you can also use `curl` command:

```bash
curl -L https://github.com/uditgaurav/hce-api-template/releases/download/0.1.0/hce-api-<ARCH> -o hce-api
chmod +x hce-api
```

## Derive Tunables

- Derive the tunables for API calls.

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

## API to Launch Chaos Experiment

This contains the API to trigger the Chaos Experiment.

### Tunables 
- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

The above tunables are mandatory to provide. You need to replace it in the given API call.

- Before getting started checkout the [pre-requisite](https://uditgaurav.github.io/hce-api-template//#pre-requisite) section and follow the steps to download `hce-api` binary.

**Non-Intractive Mode:**

- Provide the values in the flags given below to get the API command in a file.
- Please refer [derive tunables](https://uditgaurav.github.io/hce-api-template//#derive-tunables) section to know more about the tunables.

```bash
./hce-api generate --api launch-experiment --hce-endpoint=<HCE-ENDPOINT> --project-id <PROJECT-ID> --workflow-id <WORKFLOW-ID> \
--access-key <ACCESS-KEY> --access-id <ACCESS-ID> --file-name <FILE-NAME>
```

Example:

```bash
./hce-api generate --api launch-experiment --hce-endpoint=http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/ --project-id abceb5f4-4268-4467-9818-ad6e3b6bfd78 --workflow-id f4581780-efaf-4155-956e-6c379f24394b --access-key nEdGNDDrTFHyCnl --access-id adminNCWQu --file-name hce-api.sh
```

The default value for `--file-name` is `hce-api.sh` all other variables are mandatory.

**Intractive Mode:**

- Run the downloaded `hce-api` binary as shown below. It will not trigger the chaos but will prepare the launch command in a file:

- Please refer [derive tunables](https://uditgaurav.github.io/hce-api-template//#derive-tunables) section to know more about the tunables.

```bash
$ ./hce-api


		Provide the index number to create a file with the API command from the given options.
		For Example to create API that launches experiment give 1 as input.
	
		[1] API to Launch Chaos Experiment
		[2] API to Monitor Chaos Experiment
		[3] API to Validate Resilience Score
	
		
Select from the given options: 1
Provide the HCE endpoint: http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/
Provide the Project ID: abceb5f4-4268-4467-9818-ad6e3b6bfd78
Provide the Workflow ID: f4581780-efaf-4155-956e-6c379f24394b
Provide the HCE Access Key: nEdGNDDrTFHyCnl
Provide the HCE Access ID: adminNCWQu
Provide the File Name for api [Default is hce-api.sh]: hce-api.sh
```

Output:

```bash
The file containing API command is created successfully
```

- Check out the file `hce-api.sh` and get the launch command

## API to Monitor Chaos Experiment

This contains the API to monitor the Chaos Experiment that is this API will help us to wait for the workflow completion.

### Tunables 

- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

- Before getting started checkout the [pre-requisite](https://uditgaurav.github.io/hce-api-template//#pre-requisite) section and follow the steps to download `hce-api` binary.

**Non-Intractive Mode:**

- Provide the values in the flags given below to get the API command in a file.
- Please refer [derive tunables](https://uditgaurav.github.io/hce-api-template//#derive-tunables) section to know more about the tunables.

```bash
./hce-api generate --api monitor-experiment --hce-endpoint=<HCE-ENDPOINT> --project-id <PROJECT-ID> --workflow-id <WORKFLOW-ID> \
--access-key <ACCESS-KEY> --access-id <ACCESS-ID> --file-name <FILE-NAME>
```

Example:

```bash
./hce-api generate --api monitor-experiment --hce-endpoint=http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/ --project-id abceb5f4-4268-4467-9818-ad6e3b6bfd78 --workflow-id f4581780-efaf-4155-956e-6c379f24394b --access-key nEdGNDDrTFHyCnl --access-id adminNCWQu --file-name hce-api.sh
```

The default value for `--file-name` is `hce-api.sh` all other variables are mandatory.

**Intractive Mode:**

- Run the downloaded `hce-api` binary as shown below. It will not trigger the command but will prepare the monitor command in a given file:
- Please refer [derive tunables](https://uditgaurav.github.io/hce-api-template//#derive-tunables) section to know more about the tunables.

- To know more about any tunables [check here](https://uditgaurav.github.io/hce-api-template//#derive-tunables).

```bash
$ ./hce-api


		Provide the index number to create a file with the API command from the given options.
		For Example to create API that launches experiment give 1 as input.
	
		[1] API to Launch Chaos Experiment
		[2] API to Monitor Chaos Experiment
		[3] API to Validate Resilience Score
	
		
Select from the given options: 2
Provide the HCE endpoint: http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/
Provide the Project ID: abceb5f4-4268-4467-9818-ad6e3b6bfd78
Provide the Workflow ID: f4581780-efaf-4155-956e-6c379f24394b
Provide the HCE Access Key: nEdGNDDrTFHyCnl
Provide the HCE Access ID: adminNCWQu
Provide the File Name for api [Default is hce-api.sh]: hce-api.sh
```

Output:

```bash
The file containing API command is created successfully
```

- Check out the file `hce-api.sh` and get the launch command

#### A sample shell script to monitor Chaos Experiment

- In this sample script we will wait for the workflow completion with the delay of 2 seconds and 150 retries, you can adjust these values based on total chaos duration.

```bash
#!/bin/sh

$delay=2
$retry=150

cmd=$(cat hce-api.sh)
for i in range {0..$retry}; do
    res=$(echo $cmd)
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

## API to Validate Resilience Score

This contains the API to get the resilience score for a workflow run and validate against the expected probe success percentage.

### Tunables 

- `ACCESS_KEY`
- `ACCESS_ID`
- `PROJECT_ID`
- `WORKFLOW_ID`
- `HCE_ENDPOINT`

- Before getting started checkout the [pre-requisite](https://uditgaurav.github.io/hce-api-template//#pre-requisite) section and follow the steps to download `hce-api` binary.

**Non-Intractive Mode:**

- Provide the values in the flags given below to get the API command in a file.
- Please refer [derive tunables](https://uditgaurav.github.io/hce-api-template//#derive-tunables) section to know more about the tunables.

```bash
./hce-api generate --api validate-resilience-score --hce-endpoint=<HCE-ENDPOINT> --project-id <PROJECT-ID> --workflow-id <WORKFLOW-ID> \
--access-key <ACCESS-KEY> --access-id <ACCESS-ID> --file-name <FILE-NAME>
```

Example:

```bash
./hce-api generate --api validate-resilience-score --hce-endpoint=http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/ --project-id abceb5f4-4268-4467-9818-ad6e3b6bfd78 --workflow-id f4581780-efaf-4155-956e-6c379f24394b --access-key nEdGNDDrTFHyCnl --access-id adminNCWQu --file-name hce-api.sh
```

The default value for `--file-name` is `hce-api.sh` all other variables are mandatory.

**Intractive Mode:**

- Run the downloaded `hce-api` binary as shown below. It will not trigger the command but will prepare the monitor command in a given file:
- Please refer [derive tunables](https://uditgaurav.github.io/hce-api-template//#derive-tunables) section to know more about the tunables.

- To know more about any tunables [check here](https://uditgaurav.github.io/hce-api-template//#derive-tunables).

```bash
$ ./hce-api


		Provide the index number to create a file with the API command from the given options.
		For Example to create API that launches experiment give 1 as input.
	
		[1] API to Launch Chaos Experiment
		[2] API to Monitor Chaos Experiment
		[3] API to Validate Resilience Score
	
		
Select from the given options: 2
Provide the HCE endpoint: http://ae1a8f465611b4c07bbbc2e7d669f533-1139615104.us-east-2.elb.amazonaws.com:9091/
Provide the Project ID: abceb5f4-4268-4467-9818-ad6e3b6bfd78
Provide the Workflow ID: f4581780-efaf-4155-956e-6c379f24394b
Provide the HCE Access Key: nEdGNDDrTFHyCnl
Provide the HCE Access ID: adminNCWQu
Provide the File Name for api [Default is hce-api.sh]: hce-api.sh
```

Output:

```bash
The file containing API command is created successfully
```

- Check out the file `hce-api.sh` and get the launch command
#### A sample shell script to validate resiliency score

- In this sample script you will get the probe success percentage for the last workflow run, you can make use of it to compare with the expected probe success percentage.

```bash
#!/bin/sh

$expectedProbeSuccessPercentage=100

cmd=$(cat validate-api.sh)

res=$(echo $cmd)
if [ "$res" != "$expectedProbeSuccessPercentage" ]; then
    echo "The probe success percentage is: $res, expected probe sucess percentage: $expectedProbeSuccessPercentage"
    exit 1
fi

echo "The probe success percentage is equal to expected probe success percentage"
exit 0
```
