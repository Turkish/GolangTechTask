package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
	"github.com/buffup/GolangTechTask/api"
	"github.com/buffup/GolangTechTask/domain"
	"log"
)

const (
	VoteablesTableName = "Voteables"
)

type Repository interface {
	GetVoteables() ([]api.Voteable, error)
	CreateVoteable(voteable *api.CreateVoteableRequest) api.CreateVoteableResponse
	CastVote()
}

type VoteableRepo struct {
	db *dynamodb.DynamoDB
	tableName string
}

func New(config aws.Config) (*VoteableRepo, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials("DUMMYIDEXAMPLE", "DUMMYEXAMPLEKEY", ""),
	})
	if err != nil {
		panic(err)
	}
	return &VoteableRepo{
		dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")}),
		VoteablesTableName,
	}, nil
}

func (vr *VoteableRepo) CreateVoteable(voteable *domain.Voteable) (string, error) {
	vr.db.
}

func (vr *VoteableRepo) GetVoteables() ([]api.Voteable, error) {
	params := &dynamodb.ScanInput{
		TableName:                 aws.String(VoteablesTableName),
	}
	result, err := vr.db.Scan(params)
}

func (vr *VoteableRepo) CastVote() () {

}
