package domain

import (
	"context"
	"github.com/ruinnel/giregi.rip-server/common"
)

type User struct {
	ID        int64  `json:"id"`
	UID       string `json:"uid"`
	IsAdmin   bool   `json:"isAdmin"`
	Email     string `json:"email"`
	CreatedAt Time   `json:"createdAt"`
	UpdatedAt Time   `json:"updatedAt"`
}

type UserService interface {
	Login(ctx context.Context, Email string, IdToken string, TokenId int64, userAgent string) (token *Token, err error)
	GetByID(ctx context.Context, id int64) (*User, error)
	GetByAccessToken(ctx context.Context, accessToken string) (user *User, err error)
	Logout(ctx context.Context, accessToken string) error
	Tags(ctx context.Context) ([]Tag, error)
	GetArchives(ctx context.Context, archive ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error)
}

type UserRepository interface {
	Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []User, nextCursor string, err error)
	One(ctx context.Context, conditions []common.Condition) (*User, error)
	Exists(ctx context.Context, conditions []common.Condition) (bool, error)
	GetByID(ctx context.Context, id int64) (*User, error)
	Store(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id int64) error
}
