package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/models/mysql"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type tokenRepository struct {
	Conn *sql.DB
}

func NewTokenRepository(conn *sql.DB) domain.TokenRepository {
	return &tokenRepository{Conn: conn}
}

func (t tokenRepository) newTokenDomain(token *mysql.Token) domain.Token {
	var userAgent *string
	var expireAt *time.Time
	if token.UserAgent.Valid {
		userAgent = &token.UserAgent.String
	} else {
		userAgent = nil
	}
	if token.ExpireAt.Valid {
		expireAt = &token.ExpireAt.Time
	} else {
		expireAt = nil
	}
	return domain.Token{
		ID:        token.ID,
		UserID:    token.UserID,
		Token:     token.Token,
		UserAgent: userAgent,
		ExpireAt:  (*domain.Time)(expireAt),
		CreatedAt: domain.Time(token.CreatedAt),
		UpdatedAt: domain.Time(token.UpdatedAt),
	}
}

func (t tokenRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	queries := common.ConditionsToQueries(conditions)
	return mysql.Tokens(queries...).Exists(ctx, t.Conn)
}

func (t tokenRepository) GetByAccessToken(ctx context.Context, accessToken string) (*domain.Token, error) {
	token, err := mysql.Tokens(Where("token = ?", accessToken)).One(ctx, t.Conn)
	if err != nil {
		return nil, err
	}
	if token == nil {
		return nil, errors.New("token: invalid token")
	}
	user := t.newTokenDomain(token)
	return &user, nil
}

func (t tokenRepository) GetByID(ctx context.Context, id int64) (*domain.Token, error) {
	tk, err := mysql.Tokens(Where("id = ?", id)).One(ctx, t.Conn)
	if err != nil {
		return nil, err
	}
	if tk == nil {
		return nil, nil
	} else {
		token := t.newTokenDomain(tk)
		return &token, nil
	}
}

func (t tokenRepository) Store(ctx context.Context, token *domain.Token) error {
	newToken := mysql.Token{
		UserID:    token.UserID,
		Token:     token.Token,
		UserAgent: null.StringFromPtr(token.UserAgent),
		ExpireAt:  null.TimeFromPtr((*time.Time)(token.ExpireAt)),
	}

	err := newToken.Insert(ctx, t.Conn, boil.Infer())
	if err != nil {
		return err
	}
	token.ID = newToken.ID
	return nil
}

func (t tokenRepository) Update(ctx context.Context, token *domain.Token) error {
	exists, err := mysql.Tokens(Where("id = ?", token.ID)).One(ctx, t.Conn)
	if err != nil {
		return err
	}
	exists.UserID = token.UserID
	exists.Token = token.Token
	exists.UserAgent = null.StringFromPtr(token.UserAgent)
	exists.ExpireAt = null.TimeFromPtr((*time.Time)(token.ExpireAt))
	_, err = exists.Update(ctx, t.Conn, boil.Infer())
	return err
}

func (t tokenRepository) Delete(ctx context.Context, id int64) error {
	panic("implement me")
}
