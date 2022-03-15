# aws-lambda-runtime-perfcomp
 Performance comparison of Lambda runtimes - NodeJS, Python and Go


This repo hosts AWS SAM based serverless projects to run performance tests on Lambda/API Gateway with different language runtimes: Python, Node and Go. 

To run and deploy the serverless stacks, navigate to individual language runtime folder and follow the steps in README.

## Notes on test setup
 
 - Each of these runtimes is deployed to AWS using [AWS Serverless Application Model (SAM)](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html). A deployment consists of a Lambda, an API gateway and a DynamoDB table. These are completely separate from each other – no sharing of tables or API gateway between runtimes.
 - Code snippets were taken from AWS documentation, then modified to make them identical. http://docs.aws.amazon.com/en_us/lambda/latest/dg/services-apigateway-tutorial.html
 - For all runtimes, operation is identical: API takes in a json record in POST request, that is passed on to Lambda, which in turn writes it to a DynamoDB table.
 - Each runtime is configured on their latest version supported by Lambda (as of March 2022): Node: 14.x, Python: 3.9, Go: 1.x
 - Performance tests were done with identical load structure and payload - using JMeter.

## Observations

 - Go runtime seems to give a little better performance than others, in terms of response times and throughput
 - Node runtime seems to perform better than Python.
 - AWS documentation on Go is significantly less than Python/Node.
 - Package size of Go runtime is large: ~ 6 MB, compared to a few hundred bytes for Python/Node. 
 
 
Here are the stats. See '90th Perc', ‘Average’ and ‘Transactions/sec’ columns:

![Performance run statistics](https://github.com/ashankz/lambda-runtime-perf/blob/main/tests/jmeter/images/testrun-20220314.png?raw=true)

 
Performance tests were done with identical load structure and payload - using JMeter:
 

![JMeter run](https://github.com/ashankz/lambda-runtime-perf/blob/main/tests/jmeter/images/jmeter-run-20220314.png?raw=true)


Lambda Package size:

![Lambda packages](https://github.com/ashankz/lambda-runtime-perf/blob/main/tests/jmeter/images/lambda-size-20220314.png?raw=true)



