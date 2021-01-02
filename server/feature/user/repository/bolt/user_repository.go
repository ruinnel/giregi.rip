package mysql

import (
	"context"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
)

type userRepository struct {
	Conn *storm.DB
}

func NewUserRepository(conn *storm.DB) domain.UserRepository {
	return &userRepository{Conn: conn}
}

func (r userRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []domain.User, nextCursor string, err error) {
	panic("implement me")
}

func (r userRepository) One(ctx context.Context, conditions []common.Condition) (*domain.User, error) {
	matchers := common.ConditionsToMatchers(conditions)

	var users []domain.User
	err := r.Conn.Select(matchers...).Limit(1).Find(&users)
	if err != nil {
		return nil, err
	}
	if len(users) > 0 {
		result := users[0]
		return &result, nil
	} else {
		return nil, errors.New(fmt.Sprintf("user: not found(condition %v)", conditions))
	}
}

func (r userRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.User{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	var user domain.User
	err := r.Conn.One(domain.UserField.ID.Name, id, &user)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r userRepository) Store(ctx context.Context, user *domain.User) error {
	user.ID = 0
	err := r.Conn.Save(user)
	if err != nil {
		return err
	}
	return nil
}

func (r userRepository) Update(ctx context.Context, user *domain.User) error {
	exists, err := r.GetByID(ctx, user.ID)
	if err != nil {
		return err
	}
	exists.UID = user.UID
	exists.IsAdmin = user.IsAdmin
	exists.Email = user.Email
	return r.Conn.Update(exists)
}

func (r userRepository) Delete(ctx context.Context, id int64) error {
	return r.Conn.DeleteStruct(&domain.User{ID: id})
}
