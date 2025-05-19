package models

import "context"

type IModels interface {
	Save(ctx context.Context) error
	GetKey(ctx context.Context) string
	GetData(ctx context.Context) error
	Delete(ctx context.Context) error
	String() string
}
