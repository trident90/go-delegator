// Package db is a DB helper interface for AWS DynamoDB
package db

import (
	"os"
	"sync"

	"go-delegator/log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

// DynamoDBHelper is a DB helper structure
type DynamoDBHelper struct {
	Region string
	client *dynamodb.DynamoDB
}

// For singleton
var instance *DynamoDBHelper
var once sync.Once

// GetInstance returns an instance of DbHelper
// region is blank or aws-region such as ap-northeast-2
// In case of blank, use AWS_DEFAULT_REGION as region
func GetInstance(region string) *DynamoDBHelper {
	once.Do(func() {
		instance = New(region)
	})
	return instance
}

// New makes an instance for DbHelper
func New(region string) *DynamoDBHelper {
	// Check if instance is already assigned
	if instance != nil {
		return instance
	}

	// Create AWS session
	if region == "" {
		region = os.Getenv("AWS_DEFAULT_REGION")
	}
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	},
	)
	if err != nil {
		log.Error("db: failed to create AWS session")
		return nil
	}

	// Create DynamoDB client
	svc := dynamodb.New(sess)
	if svc == nil {
		log.Error("db: failed to create DynamoDB client")
		return nil
	}

	// Assign
	instance = &DynamoDBHelper{
		Region: region,
		client: svc,
	}
	return instance
}

// ListTables lists table names
func (d *DynamoDBHelper) ListTables() {
	result, err := d.client.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Info(err)
	}

	log.Info("Tables:")
	for _, n := range result.TableNames {
		log.Info(*n)
	}
}

// GetItem returns queried result following inputs
// For table named "tblName" in DynamoDB,
// Scan target value following propName is propVal
//   -----------------------------
//   |  propName  |  targetName  |
//   -----------------------------
//   |  propVal   |  targetVal   |
//   |   ...      |    ...       |
//   -----------------------------
func (d *DynamoDBHelper) GetItem(tblName, propName, propVal, targetName string) *dynamodb.ScanOutput {
	/*
		result, err := d.client.GetItem(&dynamodb.GetItemInput{
			TableName: aws.String(tblName),
			Key: map[string]*dynamodb.AttributeValue{
				propName: {
					S: aws.String(propVal),
				},
			},
		})
	*/
	filter := expression.Name(propName).Equal(expression.Value(propVal))
	projection := expression.NamesList(expression.Name(propName), expression.Name(targetName))
	expr, err := expression.NewBuilder().WithFilter(filter).WithProjection(projection).Build()
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tblName),
	}

	// Make the DynamoDB Query API call
	result, err := d.client.Scan(params)
	if err != nil {
		log.Info(err.Error())
		return nil
	}
	return result
}

// UnmarshalMap makes output data from DynamoDB output
func (d *DynamoDBHelper) UnmarshalMap(in map[string]*dynamodb.AttributeValue, out interface{}) {
	if in == nil {
		return
	} else if err := dynamodbattribute.UnmarshalMap(in, &out); err != nil {
		log.Error("db: failed to unmarshalMap of dynamoDb output")
	}
}
