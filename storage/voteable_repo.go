package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/buffup/GolangTechTask/api"
	"github.com/google/uuid"
	"log"
)

const (
	VoteablesTableName = "Voteables"
)

type VoteableRepo struct {
	db        dynamodb.DynamoDB
	tableName *string
}

func NewVoteableRepo(db dynamodb.DynamoDB) (*VoteableRepo, error) {
	vr := &VoteableRepo{
		db,
		aws.String(VoteablesTableName),
	}
	input := vr.createTableInput()
	err := createTable(&db, input, false)
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

func (vr *VoteableRepo) GetVoteable(uuid string) (*api.Voteable, error) {
	result, err := vr.db.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"uuid": {
				S: aws.String(uuid),
			},
		},
		TableName:                vr.tableName,
	})
	if err != nil {
		return nil, err
	}
	v := &api.Voteable{}
	err = dynamodbattribute.UnmarshalMap(result.Item, v)
	if err != nil {
		return nil, err
	}
	return v, nil
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
