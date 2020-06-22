package main
import (
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb"

    "fmt"
    "os"
)

func main(){
  sess := session.Must(session.NewSessionWithOptions(session.Options{
    SharedConfigState: session.SharedConfigEnable,
  }))

  // Create DynamoDB client
  sess, err := session.NewSession(&aws.Config{
    Region:   aws.String("us-west-2"),
    Endpoint: aws.String("http://localhost:8000")})
    if err != nil {
      fmt.Println(err)
    }
  svc := dynamodb.New(sess)
  tableName := "FileMetadata"

  input := &dynamodb.CreateTableInput{
    AttributeDefinitions: []*dynamodb.AttributeDefinition{
      {
        AttributeName: aws.String("Key"),
        AttributeType: aws.String("N"),
      },
          },
    KeySchema: []*dynamodb.KeySchemaElement{
      {
        AttributeName: aws.String("Key"),
        KeyType:       aws.String("HASH"),
      },
    },
    ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
      ReadCapacityUnits:  aws.Int64(10),
      WriteCapacityUnits: aws.Int64(10),
    },
    TableName: aws.String(tableName),
  }
  _, errinput := svc.CreateTable(input)
  if errinput != nil {
    fmt.Println("Got error calling CreateTable:")
    fmt.Println(err.Error())
    os.Exit(1)
  }

  fmt.Println("Created the table", tableName)
}
