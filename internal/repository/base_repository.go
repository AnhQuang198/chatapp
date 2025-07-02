package repository

import "context"

type BaseRepository[T any, C any] interface {
	Create(ctx context.Context, data C) error
	Delete(ctx context.Context, id int64) error
	GetById(ctx context.Context, id int64) (*T, error)
	GetByIds(ctx context.Context, ids []int64) ([]T, error)
}
