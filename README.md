# aws lambda language runtime performance comparison
 Performance comparison of Lambda runtimes - NodeJS, Python and Go


This repo hosts AWS SAM based serverless projects to run performance tests on Lambda/API Gateway with different language runtimes: Python, Node and Go. 
The tests are conducted in as identical manner as possible. Each stack has a simple api endpoint that takes in a json item and writes it to a DynamoDB table.
Lambda memory and DynamodB configurations are kept identical.

## Deploying stacks to AWS

### Pre-requisites
1. Download and install the listed services, if not already installed:

Run the following commands to check if the below mention services are already installed:

```
aws --version
sam --version
go version

```

If not installed, follow these links on installation steps:

[AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-mac.html)

[AWS SAM CLI](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

[Golang](https://go.dev/doc/install)


2. In order to deploy stacks to AWS, either refer to or create an IAM user in your AWS Account, with appropriate permissions to be able to deploy a Cloudformation stack. Configure with 'Programmatic Access' so that you get Access KeyID/Secret Access credentials. If this step has not been done earlier, run the following command. Provide the credentials when prompted.

```

aws configure

```

To deploy AWS stacks for each language, follow the scripts mentioned below. Deployment is done using AWS SAM CLI. It requires appropriate credentials present in .aws/credentials file. 

At the end of deploy, note the ***'APIDomain'*** Key's value - these will need to be passed to JMeter script, as mentioned below.


### Python

Run these commands at the terminal, with directory being the current folder:

```
cd python

sam build

sam deploy --template-file template.yaml --stack-name perfTestPocApiOnPython --region us-west-2  --resolve-s3 --no-fail-on-empty-changeset --capabilities CAPABILITY_IAM 

cd ..

```
### Node

Run these commands at the terminal, with directory being the current folder:

```
cd node

sam build

sam deploy --template-file template.yaml --stack-name perfTestPocApiOnNode --region us-west-2  --resolve-s3 --no-fail-on-empty-changeset --capabilities CAPABILITY_IAM 

cd ..

```


### Golang

Run these commands at the terminal, with directory being the current folder:

```
cd golang

GOOS=linux GOARCH=amd64 go build -o dynamo-handler/dynamo-handler ./dynamo-handler

sam package --template-file template.yaml --resolve-s3 --output-template-file packaged.yaml --region us-west-2

sam deploy --template-file packaged.yaml --capabilities CAPABILITY_IAM --stack-name perfTestPocApiOnGo  --region us-west-2

cd ..

```



## Running performance test

### Steps to run jmeter load tests locally: 

_If jmeter is not installed:_
1. Download jmeter from [this location](https://jmeter.apache.org/download_jmeter.cgi)
2. Unzip the download package to a particular location. 
3. Add the jmeter bin folder to your PATH:
   On  MAC, run this in terminal (Double check the location and jemter version): export PATH=$PATH:~/Downloads/apache-jmeter-5.4.3/bin
   On Windows, follow this [link](https://stackoverflow.com/a/44272417) to add to PATH.


_If jmeter is installed:_

### Test Script

Run the following command after replacing the domain name with appropriate API gateway domain.
This test script executes 20,000 requests against the api over a period of about 15 minutes. 

Check the correct location for the jmeter binary.

__Make sure to update the *Jdomain* parameter with the value of 'APIDomain' Key in output of SAM deploy commands__ 

```

~/Downloads/apache-jmeter-5.4.3/bin/jmeter -n -t tests/jmeter/setup/createapi-test.jmx  -Jdomain=xxxxx.execute-api.us-west-2.amazonaws.com -Jthreads=200 -Jrampup=0 -Jiterations=100 -JiterationDelay=10000 -e -l tests/jmeter/runs/logs_$(date '+%Y%m%d%H%M%S').jtl -o tests/jmeter/runs/run_$(date '+%Y%m%d%H%M%S')

```

Results are in the 'tests/jmeter/runs' folder. To view the results, go to the specific 'run_yyyymmdd' folder and open the index.html page.

## Postman collection
These can be used to run some basic tests on the API.
Import the collection and set the variables to appropriate API gateway domains

domain-node
domain-go
domain-python


## Cleaning up resources

Run the following command to remove all resources from AWS:

```
sam delete perfTestPocApiOnPython --region us-west-2

sam delete perfTestPocApiOnNode --region us-west-2

sam delete perfTestPocApiOnGo --region us-west-2

```

When prompted to provide stack name, input either: perfTestPocApiOnPython, perfTestPocApiOnNode, perfTestPocApiOnGo
Respond 'y' to next questions.
 

## Notes on test setup
 
 - Each of these runtimes is deployed to AWS using [AWS Serverless Application Model (SAM)](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html). A deployment consists of a Lambda, an API gateway and a DynamoDB table. These are completely separate from each other – no sharing of tables or API gateway between runtimes.
 - Code snippets were taken from AWS documentation, then modified to make them identical. http://docs.aws.amazon.com/en_us/lambda/latest/dg/services-apigateway-tutorial.html
 - For all runtimes, operation is identical: API takes in a json record in POST request, that is passed on to Lambda, which in turn writes it to a DynamoDB table.
 - Each runtime is configured on their latest version supported by Lambda (as of April 2022): Node: 14.x, Python: 3.9, Go: 1.x
 - Performance tests are done with identical load structure and payload - using JMeter.

## Observations

 - Go runtime seems to give a little better performance than others, in terms of response times and throughput
 - AWS documentation on Go is significantly less than Python/Node.
 - Package size of Go runtime is large: ~ 6 MB, compared to a few hundred bytes for Python/Node. 
 
 
Here are the stats. See '90th Perc', ‘Average’ and ‘Transactions/sec’ columns:

![Performance run statistics](https://github.com/ashankz/lambda-runtime-perf/blob/main/tests/jmeter/images/testrun-20220314.png?raw=true)

 
Performance tests were done with identical load structure and payload - using JMeter:
 

![JMeter run](https://github.com/ashankz/lambda-runtime-perf/blob/main/tests/jmeter/images/jmeter-run-20220314.png?raw=true)


Lambda Package size:

![Lambda packages](https://github.com/ashankz/lambda-runtime-perf/blob/main/tests/jmeter/images/lambda-size-20220314.png?raw=true)



