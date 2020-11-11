package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
)

type Token struct {
	ID        int64   `json:"id"`
	UserID    int64   `json:"userId"`
	Token     string  `json:"token"`
	TokenID   string  `json:"tokenId,omitempty"` // response only
	UserAgent *string `json:"userAgent,omitempty"`
	ExpireAt  *Time   `json:"expireAt,omitempty"`
	CreatedAt Time    `json:"createdAt"`
	UpdatedAt Time    `json:"updatedAt"`
}

type TokenRepository interface {
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByAccessToken(ctx context.Context, accessToken string) (*Token, error)
	GetByID(ctx context.Context, id int64) (*Token, error)
	Store(ctx context.Context, token *Token) error
	Update(ctx context.Context, token *Token) error
	Delete(ctx context.Context, id int64) error
}
