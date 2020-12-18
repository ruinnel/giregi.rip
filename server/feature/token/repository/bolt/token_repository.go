package bolt

import (
	"context"
	"github.com/asdine/storm/v3"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
)

type tokenRepository struct {
	Conn *storm.DB
}

func NewTokenRepository(conn *storm.DB) domain.TokenRepository {
	return &tokenRepository{Conn: conn}
}

func (r tokenRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.Token{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r tokenRepository) GetByAccessToken(ctx context.Context, accessToken string) (*domain.Token, error) {
	var token domain.Token
	err := r.Conn.One(domain.TokenField.Token.Name, accessToken, &token)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r tokenRepository) GetByID(ctx context.Context, id int64) (*domain.Token, error) {
	var token domain.Token
	err := r.Conn.One(domain.TokenField.ID.Name, id, &token)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r tokenRepository) Store(ctx context.Context, token *domain.Token) error {
	err := r.Conn.Save(token)
	if err != nil {
		return err
	}
	return nil
}

func (r tokenRepository) Update(ctx context.Context, token *domain.Token) error {
	exists, err := r.GetByID(ctx, token.ID)
	if err != nil {
		return err
	}
	exists.UserID = token.UserID
	exists.Token = token.Token
	exists.UserAgent = token.UserAgent
	exists.ExpireAt = token.ExpireAt
	return r.Conn.Update(exists)
}

func (r tokenRepository) Delete(ctx context.Context, id int64) error {
	return r.Conn.DeleteStruct(&domain.Token{ID: id})
}
