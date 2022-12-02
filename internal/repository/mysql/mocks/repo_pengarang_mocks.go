package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-router-sa/domain/entity"
)

type RepoPengarangMocks struct {
	mock.Mock
}

func (p *RepoPengarangMocks) GetListPengarang(ctx context.Context) ([]*entity.Pengarang, error) {
	args := p.Called(ctx)
	return args.Get(0).([]*entity.Pengarang), args.Error(1)
}

func (p *RepoPengarangMocks) GetPengarangById(ctx context.Context, id int) (*entity.Pengarang, error) {
	args := p.Called(ctx, id)
	return args.Get(0).(*entity.Pengarang), args.Error(1)
}
