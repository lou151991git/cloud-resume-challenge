package main

import (
	"context"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func buildResponse(count string) string {
	return `{"count":` + count + `}`
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	tableName := os.Getenv("TABLE_NAME")

	cfg, err := config.LoadDefaultConfig(context.TODO())

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	client := dynamodb.NewFromConfig(cfg)

	result, err := client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(tableName),

		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: "counter",
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	count := result.Item["visits"].(*types.AttributeValueMemberN).Value

	return events.APIGatewayProxyResponse{
		StatusCode: 200,

		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},

		Body: buildResponse(count),
	}, nil
}

func main() {
	lambda.Start(handler)
}
