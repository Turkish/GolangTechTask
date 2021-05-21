package api

import (
	"context"
	"errors"
	"github.com/buffup/GolangTechTask/storage"
)

type VotingService struct {
	repo storage.VoteableRepo
}

func NewVotingService(repo storage.VoteableRepo) VotingService {
	return VotingService{repo}
}

func (s VotingService) mustEmbedUnimplementedVotingServiceServer() {
	return
}

func (s VotingService) CreateVoteable(ctx context.Context, cvr *CreateVoteableRequest) (*CreateVoteableResponse, error) {
	uuid, err := s.repo.CreateVoteable(&Voteable{
		Question: cvr.GetQuestion(),
		Answers: cvr.GetAnswers(),
	})
	if err != nil {
		return nil, err
	}
	return &CreateVoteableResponse{Uuid: uuid}, nil
}

func (s VotingService) ListVoteables(ctx context.Context, request *ListVoteableRequest) (*ListVoteableResponse, error) {
	vs, err := s.repo.GetVoteables()
	if err != nil {
		return nil, err
	}
	return &ListVoteableResponse{Votables: vs}, nil
}

func (s VotingService) CastVote(context.Context, *CastVoteRequest) (*CastVoteResponse, error) {

	return nil, errors.New("not implemented")
}


