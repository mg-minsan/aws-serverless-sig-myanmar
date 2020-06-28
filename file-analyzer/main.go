package main

import (
	"context"
	"errors"
	"log"
	"time"

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

func handler(ctx context.Context, event events.S3Event) error{
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))
  svc := dynamodb.New(sess)
  for _, record := range event.Records {
    s3 := record.S3
    now := time.Now().String()
    item := Item{
      Key: s3.Object.Key,
      Timestamp: now,
    }
    av, err := dynamodbattribute.MarshalMap(item)
    if err != nil {
      return errors.New("Got error marshalling new movie item:")
    }
    input := &dynamodb.PutItemInput{
      Item:      av,
      TableName: aws.String("FileMetadata"),
    }
    _, err = svc.PutItem(input)
    if err != nil{
      return err
    }
  }
    log.Print("Uploaded")
    return nil
}

