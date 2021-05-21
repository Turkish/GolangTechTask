package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/buffup/GolangTechTask/api"
	"github.com/google/uuid"
	"log"
)

const (
	VoteablesTableName = "Voteables"
)

type Repository interface {
	CreateVoteable(*api.Voteable) (string, error)
	GetVoteables() ([]*api.Voteable, error)
	CastVote()
}

type VoteableRepo struct {
	db        *dynamodb.DynamoDB
	tableName *string
}

func New(config *aws.Config) (*VoteableRepo, error) {
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	vr := &VoteableRepo{
		dynamodb.New(sess, config),
		aws.String(VoteablesTableName),
	}
	input := vr.createTableInput()
	err = createTable(vr.db, input, false)
	if err != nil {
		return nil, err
	}
	return vr, nil
}

func (vr *VoteableRepo) CreateVoteable(v *api.Voteable) (string, error) {
	v.Uuid = uuid.New().String()
	av, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		return "", err
	}
	_, err = vr.db.PutItem(&dynamodb.PutItemInput{Item: av, TableName: vr.tableName})
	if err != nil {
		return "", err
	}
	return v.Uuid, nil
}

func (vr *VoteableRepo) createTableInput() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("uuid"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("uuid"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: vr.tableName,
	}
}

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


func (vr *VoteableRepo) GetVoteables() ([]*api.Voteable, error) {
	params := &dynamodb.ScanInput{
		TableName:                 vr.tableName,
	}
	result, err := vr.db.Scan(params)
	if err != nil {
		return nil, err
	}
	var res []*api.Voteable
	for _, item := range result.Items {
		v := &api.Voteable{}
		err = dynamodbattribute.UnmarshalMap(item, v)
		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}
		res = append(res, v)
	}
	return res, nil
}

func (vr *VoteableRepo) CastVote() () {

}

func initDynamodb() {

}
