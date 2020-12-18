package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
	"reflect"
)

type Token struct {
	ID        int64   `json:"id" mysql:"id" storm:"id,increment"`
	UserID    int64   `json:"userId" mysql:"user_id" storm:"index"`
	Token     string  `json:"token" mysql:"token" storm:"index"`
	TokenID   string  `json:"tokenId,omitempty" mysql:"token_id"` // response only
	UserAgent *string `json:"userAgent,omitempty" mysql:"user_agent"`
	ExpireAt  *Time   `json:"expireAt,omitempty" mysql:"expire_at" storm:"index"`
	CreatedAt Time    `json:"createdAt" mysql:"created_at" storm:"index"`
	UpdatedAt Time    `json:"updatedAt" mysql:"updated_at"`
}

type tokenField struct {
	ID        reflect.StructField
	UserID    reflect.StructField
	Token     reflect.StructField
	TokenID   reflect.StructField
	UserAgent reflect.StructField
	ExpireAt  reflect.StructField
	CreatedAt reflect.StructField
	UpdatedAt reflect.StructField
}

var TokenField = makeFields(&Token{}, &tokenField{}).(*tokenField)

type TokenRepository interface {
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByAccessToken(ctx context.Context, accessToken string) (*Token, error)
	GetByID(ctx context.Context, id int64) (*Token, error)
	Store(ctx context.Context, token *Token) error
	Update(ctx context.Context, token *Token) error
	Delete(ctx context.Context, id int64) error
}
