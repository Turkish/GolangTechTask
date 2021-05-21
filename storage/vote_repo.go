package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/buffup/GolangTechTask/domain"
	"github.com/google/uuid"
)

const (
	VotesTableName = "Votes"
)

type VoteRepo struct {
	db        dynamodb.DynamoDB
	tableName *string
}

func NewVoteRepo(db dynamodb.DynamoDB) (*VoteRepo, error) {
	vr := &VoteRepo{
		db,
		aws.String(VotesTableName),
	}
	input := vr.createTableInput()
	err := createTable(&db, input, false)
	if err != nil {
		return nil, err
	}
	return vr, nil
}

func (vr *VoteRepo) CastVote(v domain.Vote) (string, error) {
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


func (vr *VoteRepo) createTableInput() *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("vote_uuid"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("vote_uuid"),
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

