package main

import (
	"context"
	"encoding/json"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type ApiResponseMessage struct {
	Message string `json:"message"`
}

type Payload struct {
	Item map[string]interface{}
}
type RequestBody struct {
	Operation string
	Payload   Payload
}

var tableName string
var svc *dynamodb.Client
var cfg aws.Config

func init() {
	var tableVarExists bool

	tableName, tableVarExists = os.LookupEnv("TABLE_NAME")
	if !tableVarExists {
		tableName = "poc-items-go"
	}

	var err error

	cfg, err = config.LoadDefaultConfig(context.TODO(), func(o *config.LoadOptions) error {
		o.Region = os.Getenv("AWS_REGION")
		return nil
	})
	if err != nil {
		panic(err)
	}

	svc = dynamodb.NewFromConfig(cfg)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	//requestJson, _ := json.Marshal(request)
	//log.Printf("EVENT: %s", requestJson)

	var body RequestBody
	json.Unmarshal([]byte(request.Body), &body)

	item := body.Payload.Item
	responseJson, _ := json.Marshal(&ApiResponseMessage{Message: "ok"})
	statusCode := 200

	if body.Operation == "create" {

		itemJson, marshallErr := attributevalue.MarshalMap(item)
		if marshallErr != nil {
			panic(marshallErr)
		}

		//fmt.Println("marshelled item:", itemJson)

		_, err := svc.PutItem(context.TODO(), &dynamodb.PutItemInput{
			TableName: aws.String(tableName),
			Item:      itemJson,
		})
		if err != nil {
			panic(err)
		}

	} else if body.Operation == "echo" {
		responseJson, _ = json.Marshal(item)
	} else {
		responseJson, _ = json.Marshal(&ApiResponseMessage{Message: "not supported"})
		statusCode = 400
	}

	resp := events.APIGatewayProxyResponse{StatusCode: statusCode, Body: string(responseJson)}
	resp.Headers = make(map[string]string)
	resp.Headers["Content-Type"] = "application/json"

	return resp, nil

}

func main() {
	lambda.Start(handler)
}
