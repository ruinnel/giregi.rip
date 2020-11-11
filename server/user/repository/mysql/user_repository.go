package mysql

import (
	"context"
	"database/sql"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/models/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepository struct {
	Conn *sql.DB
}

func (r userRepository) newUserDomain(user *mysql.User) domain.User {
	return domain.User{
		ID:        user.ID,
		UID:       user.UID,
		IsAdmin:   user.IsAdmin,
		Email:     user.Email,
		CreatedAt: domain.Time(user.CreatedAt),
		UpdatedAt: domain.Time(user.UpdatedAt),
	}
}

func (r userRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []domain.User, nextCursor string, err error) {
	panic("implement me")
}

func (r userRepository) One(ctx context.Context, conditions []common.Condition) (*domain.User, error) {
	queries := common.ConditionsToQueries(conditions)

	user, err := mysql.Users(queries...).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	result := r.newUserDomain(user)
	return &result, nil
}

func (r userRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	queries := common.ConditionsToQueries(conditions)
	return mysql.Users(queries...).Exists(ctx, r.Conn)
}

func (r userRepository) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	usr, err := mysql.Users(Where("id = ?", id)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if usr == nil {
		return nil, nil
	} else {
		user := r.newUserDomain(usr)
		return &user, nil
	}
}

func (r userRepository) Store(ctx context.Context, user *domain.User) error {
	newUser := mysql.User{
		UID:     user.UID,
		IsAdmin: user.IsAdmin,
		Email:   user.Email,
	}

	err := newUser.Insert(ctx, r.Conn, boil.Infer())
	if err != nil {
		return err
	}
	user.ID = newUser.ID
	return nil
}

func (r userRepository) Update(ctx context.Context, user *domain.User) error {
	exists, err := mysql.Users(Where("id = ?", user.ID)).One(ctx, r.Conn)
	if err != nil {
		return err
	}
	exists.UID = user.UID
	exists.IsAdmin = user.IsAdmin
	exists.Email = user.Email
	_, err = exists.Update(ctx, r.Conn, boil.Infer())
	return err
}

func (r userRepository) Delete(ctx context.Context, id int64) error {
	_, err := mysql.Users(Where("id = ?", id)).One(ctx, r.Conn)
	return err
}

func NewUserRepository(conn *sql.DB) domain.UserRepository {
	return &userRepository{Conn: conn}
}
