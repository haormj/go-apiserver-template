package dao

import (
	"context"

	"github.com/haormj/go-apiserver-template/internal/model"
)

type User interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int64) (*model.User, error)
	UpdateByID(ctx context.Context, id int64, user *model.User) error
	GetByName(name string) (*model.User, error)
	List(pageNum, pageSize int) ([]*model.User, int64, error)
	Count() (int64, error)
}
