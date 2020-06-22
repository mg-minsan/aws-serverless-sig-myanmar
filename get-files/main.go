package main

import (
	"encoding/json"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
  "github.com/aws/aws-sdk-go/aws"
  "github.com/aws/aws-sdk-go/aws/session"
  "github.com/aws/aws-sdk-go/service/dynamodb"
  "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
type Item struct{
  Key string
  Timestamp string
}
func main(){
  lambda.Start(handler)
}


func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))
  var items[]Item
  svc := dynamodb.New(sess)
  result, err := svc.Scan(&dynamodb.ScanInput{
    TableName: aws.String("FileMetadata"),
  })
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }
  err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)
  if err != nil {
    return events.APIGatewayProxyResponse{}, err
  }
  body, parsedErr := json.Marshal(items)
  if parsedErr != nil{
      return events.APIGatewayProxyResponse{}, err
    }

  return events.APIGatewayProxyResponse{
    Body: string(body),
    StatusCode: 200,
  }, nil
}
