package api

import (
	"context"
	"errors"
	"github.com/buffup/GolangTechTask/domain"
	"github.com/buffup/GolangTechTask/storage"
)

type VotingService struct {
	voteableRepo storage.VoteableRepo
	voteRepo storage.VoteRepo
}

func NewVotingService(voteableRepo storage.VoteableRepo, voteRepo storage.VoteRepo) VotingService {
	return VotingService{voteableRepo, voteRepo}
}

func (s VotingService) mustEmbedUnimplementedVotingServiceServer() {
	return
}

func (s VotingService) CreateVoteable(ctx context.Context, cvr *CreateVoteableRequest) (*CreateVoteableResponse, error) {
	vUuid, err := s.voteableRepo.CreateVoteable(&Voteable{
		Question: cvr.GetQuestion(),
		Answers: cvr.GetAnswers(),
	})
	if err != nil {
		return nil, err
	}
	return &CreateVoteableResponse{Uuid: vUuid}, nil
}

func (s VotingService) ListVoteables(ctx context.Context, request *ListVoteableRequest) (*ListVoteableResponse, error) {
	vs, err := s.voteableRepo.GetVoteables()
	if err != nil {
		return nil, err
	}
	return &ListVoteableResponse{Votables: vs}, nil
}

func (s VotingService) CastVote(ctx context.Context, cvr *CastVoteRequest) (*CastVoteResponse, error) {
	vote := domain.Vote{
		VoteableUuid: cvr.GetUuid(),
		AnswerIndex: cvr.GetAnswerIndex(),
	}
	err := s.validateVote(vote)
	if err != nil {
		return nil, err
	}
	_, err = s.voteRepo.CastVote(vote)
	if err != nil {
		return nil, err
	}

	return &CastVoteResponse{}, errors.New("not implemented")
}

func (s VotingService) validateVote(vote domain.Vote) error {
	voteable, err := s.voteableRepo.GetVoteable(vote.VoteableUuid)
	if err != nil {
		return err
	}
	if len(voteable.Answers) - 1 < int(vote.AnswerIndex)  {
		return errors.New("vote not valid, answer index out of bounds")
	}
	return nil
}

