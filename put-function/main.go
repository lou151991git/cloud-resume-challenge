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

func buildPutResponse(message string) string {
	return `{"message":"` + message + `"}`
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

	_, err = client.UpdateItem(context.TODO(), &dynamodb.UpdateItemInput{
		TableName: aws.String(tableName),

		Key: map[string]types.AttributeValue{
			"ID": &types.AttributeValueMemberS{
				Value: "counter",
			},
		},

		UpdateExpression: aws.String("SET visits = visits + :inc"),

		ExpressionAttributeValues: map[string]types.AttributeValue{
			":inc": &types.AttributeValueMemberN{
				Value: "1",
			},
		},
	})

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       err.Error(),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin": "*",
		},
		Body: buildPutResponse("count updated"),
	}, nil
}

func main() {
	lambda.Start(handler)
}
