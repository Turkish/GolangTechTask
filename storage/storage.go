package storage

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/buffup/GolangTechTask/api"
	"log"
)

const (
	VoteablesTableName = "Voteables"
)

type Repository interface {
	GetVoteables() ([]api.Voteable, error)
	CreateVoteable(*api.CreateVoteableRequest) (api.CreateVoteableResponse, error)
	CastVote()
}

type VoteableRepo struct {
	db *dynamodb.DynamoDB
	tableName *string
}

func New(config *aws.Config) (*VoteableRepo, error) {
	sess, err := session.NewSession(config)
	if err != nil {
		panic(err)
	}
	vr :=  &VoteableRepo{
		dynamodb.New(sess, config),
		aws.String(VoteablesTableName),
	}
	vr.createTable()
	return vr, nil
}

func (vr *VoteableRepo) CreateVoteable(v *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	av, err := dynamodbattribute.MarshalMap(v)
	if err != nil {
		return nil, err
	}
	output, err := vr.db.PutItem(&dynamodb.PutItemInput{Item: av, TableName: vr.tableName})
	if err != nil {
		return nil, err
	}
	fmt.Printf("%+v", output)
	return nil, nil
}

func (vr *VoteableRepo) createTable() {
	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Uuid"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Question"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Answers"),
				AttributeType: aws.String("SS"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Uuid"),
				KeyType:       aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: vr.tableName,
	}

	_, err := vr.db.CreateTable(input)
	if err != nil {
		log.Fatalf("Got error calling CreateTable: %s", err)
	}

	fmt.Println("Created the table", vr.tableName)
}
//
//func (vr *VoteableRepo) GetVoteables() ([]api.Voteable, error) {
//	params := &dynamodb.ScanInput{
//		TableName:                 aws.String(VoteablesTableName),
//	}
//	result, err := vr.db.Scan(params)
//}
//
//func (vr *VoteableRepo) CastVote() () {
//
//}
//
//func initDynamodb() {
//
//}
