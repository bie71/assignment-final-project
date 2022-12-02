package mocks

import (
	"context"
	"github.com/stretchr/testify/mock"
	"go-router-sa/domain/entity"
)

type RepoBukuMocks struct {
	mock.Mock
}

func (b *RepoBukuMocks) InsertDataBuku(ctx context.Context, dataBuku *entity.Buku) error {
	args := b.Called(ctx, dataBuku)
	return args.Error(0)
}

func (b *RepoBukuMocks) GetListBuku(ctx context.Context) ([]*entity.Buku, error) {
	args := b.Called(ctx)
	return args.Get(0).([]*entity.Buku), args.Error(1)
}

func (b *RepoBukuMocks) GetBukuByKode(ctx context.Context, kode string) (*entity.Buku, error) {
	args := b.Called(ctx, kode)
	return args.Get(0).(*entity.Buku), args.Error(1)
}

func (b *RepoBukuMocks) UpdateBukuByKode(ctx context.Context, dataBuku *entity.Buku, kode string) error {
	args := b.Called(ctx, dataBuku, kode)
	return args.Error(0)
}

func (b *RepoBukuMocks) DeleteBukuByKode(ctx context.Context, kode string) error {
	args := b.Called(ctx, kode)
	return args.Error(0)
}

func (b *RepoBukuMocks) GetBukuByPengarangId(ctx context.Context, pengarangId int) ([]*entity.Buku, error) {
	args := b.Called(ctx, pengarangId)
	return args.Get(0).([]*entity.Buku), args.Error(1)
}

func (b *RepoBukuMocks) GetBackupBukuByKode(ctx context.Context, kode string) (*entity.Buku, error) {
	args := b.Called(ctx, kode)
	return args.Get(0).(*entity.Buku), args.Error(1)
}

func (b *RepoBukuMocks) InsertDataBackupBuku(ctx context.Context, dataBuku *entity.Buku) error {
	args := b.Called(ctx, dataBuku)
	return args.Error(0)
}
