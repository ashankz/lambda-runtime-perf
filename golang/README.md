# CRUD api : Basic api for write operations on a DynamoDB Table using Go language
#

This project hosts basic REST API for Write operations on a simple DynamoDb table. It is based on [AWS Serverless Application Model (SAM)](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/what-is-sam.html) 

# Pre-requisites to running the app
1. Download and install the listed services, if not already installed:

[Docker](https://docs.docker.com/get-docker/)

[AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/install-cliv2-mac.html)

[AWS SAM](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install.html)

Run the following commands to check if the above services are installed:

```
aws --version
docker version
sam --version

```

2. This step is required only if planning to run [SAM cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-deploying.html) command to deploy code to AWS. Create an IAM user in your AWS Account with appropriate permissions to be able to deploy a Cloudformation stack. Configure with 'Programmatic Access' so that you get Access KeyID/Secret Access credentials. If this step has not been done earlier, run the following command. Provide the credentials when prompted.

```
aws configure
```


# Run the app
Run the following command to build and start the api locally. 

```
sam build
sam local start-api
```

## Testing the API locally
Only 'echo' operation can be tested, as Dynamodb local is not setup. 

```
curl --location --request POST 'http://127.0.0.1:3000/items' \
--header 'Content-Type: application/json' \
--data-raw '{
    "operation": "echo",
    "payload": {
        "Item": {
            "id": "1",
            "year": "2022"
        }
    }
}'

```

### Deploying code to AWS 

Use [SAM cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-deploying.html) to deploy code to AWS. In general, this may not be required, as it is done as part CI/CD. Also, it requires appropriate credentials present in .aws/credentials file. 

The following commands:
- Builds the Go binary
- Uses AWS SAM cli to build and package the template and files.
- Uses AWS SAM cli to deploy the package to AWS. 

```
GOOS=linux GOARCH=amd64 go build -o dynamo-handler/dynamo-handler ./dynamo-handler

sam package --template-file template.yaml --resolve-s3 --output-template-file packaged.yaml --region us-west-2

sam deploy --template-file packaged.yaml --capabilities CAPABILITY_IAM --stack-name perfTestPocApiOnGo  --region us-west-2

```

### Cleaning up resources

Run the following command to remove all resources from AWS:

```
sam delete perfTestPocApiOnGo --region us-west-2

```