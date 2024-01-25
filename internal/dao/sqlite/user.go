package sqlite

import (
	"context"

	"github.com/haormj/go-apiserver-template/internal/dao"
	"github.com/haormj/go-apiserver-template/internal/model"
)

type user struct{}

func NewUser() (dao.User, error) {
	return nil, nil
}

func(u *user) Create(ctx context.Context, user *model.User) error {
	return nil
}

func(u *user) GetByID(ctx context.Context, id int64) (*model.User, error) {
	return nil, nil
}

func(u *user) UpdateByID(ctx context.Context, id int64, user *model.User) error {
	return nil
}

func(u *user) GetByName(name string) (*model.User, error) {
	return nil,nil
}

func(u *user) List(pageNum, pageSize int) ([]*model.User, int64, error) {
	return nil, 0,nil
}

func(u *user) Count() (int64, error) {
	return 0, nil
}