package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func tableExists(db *dynamodb.DynamoDB, tableName *string) (bool, error) {
	_, err := db.DescribeTable(&dynamodb.DescribeTableInput{TableName: tableName})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeResourceNotFoundException:
				return false, nil
			default:
				return false, err
			}
		}
	}
	return true, nil
}

// TODO correctly handle schema migration instead of dropping and recreating table
func createTable(db *dynamodb.DynamoDB, input *dynamodb.CreateTableInput, refreshSchema bool) error {
	yes, err := tableExists(db, input.TableName)
	if err != nil {
		return err
	}
	if yes {
		if !refreshSchema {
			return nil
		}
		_, err := db.DeleteTable(&dynamodb.DeleteTableInput{TableName: input.TableName})
		if err != nil {
			return err
		}
		fmt.Printf("Table %s deleted\n", *input.TableName)
	}
	_, err = db.CreateTable(input)
	if err != nil {
		return err
	}
	fmt.Printf("Table %s recreated\n", *input.TableName)

	return nil
}