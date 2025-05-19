package storage

import (
	"context"
	"tkh/models"
)

func init() {
	st = &storage{}
}

type IStorage interface {
	Save(ctx context.Context, models models.IModels) error
	Get(ctx context.Context, models models.IModels) error
	Delete(ctx context.Context, models models.IModels) error
}

type storageCtxKey string

var st *storage

type storage struct {
}

func GetStorage(ctx context.Context) IStorage {
	s, ok := ctx.Value(storageCtxKey("storage")).(IStorage)
	if ok {
		return s
	}
	return st
}
func SetStorage(ctx context.Context, s IStorage) context.Context {
	return context.WithValue(ctx, storageCtxKey("storage"), s)
}

func (s *storage) Save(ctx context.Context, m models.IModels) error {
	return m.Save(ctx)
}

func (s *storage) Get(ctx context.Context, m models.IModels) error {
	return m.GetData(ctx)

}

func (s *storage) Delete(ctx context.Context, m models.IModels) error {
	return m.Delete(ctx)
}
