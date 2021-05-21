package main

import (
	"context"
	"errors"
	"github.com/buffup/GolangTechTask/api"
	"github.com/buffup/GolangTechTask/storage"
	"google.golang.org/grpc"
)

type Service struct {
	repo storage.Repository
}

func NewService(repo storage.Repository) Service {
	return Service{repo}
}

func (s Service) CreateVoteable(ctx context.Context, cvr *api.CreateVoteableRequest) (*api.CreateVoteableResponse, error) {
	//uuid, err := s.repo.CreateVoteable(cvr)
	return nil, errors.New("not implemented")
}

func (s Service) ListVoteables(ctx context.Context, in *api.ListVoteableRequest, opts ...grpc.CallOption) (*api.ListVoteableResponse, error) {
	return nil, errors.New("not implemented")
}

func (s Service) CastVote(context.Context, *api.CastVoteRequest) (*api.CastVoteResponse, error) {
	return nil, errors.New("not implemented")
}


