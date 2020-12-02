package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"math"
	"time"
)

type userService struct {
	userRepository    domain.UserRepository
	tokenRepository   domain.TokenRepository
	tagRepository     domain.TagRepository
	archiveRepository domain.ArchiveRepository
}

func NewUserService(
	userRepository domain.UserRepository,
	tokenRepository domain.TokenRepository,
	tagRepository domain.TagRepository,
	archiveRepository domain.ArchiveRepository,
) domain.UserService {
	return &userService{
		userRepository:    userRepository,
		tokenRepository:   tokenRepository,
		tagRepository:     tagRepository,
		archiveRepository: archiveRepository,
	}
}

func (u *userService) Login(ctx context.Context, email string, idToken string, tokenId int64, userAgent string) (*domain.Token, error) {
	authClient, err := common.NewAuthClient()
	if err != nil {
		return nil, common.NewFirebaseError("firebase init fail", err)
	}

	token, err := authClient.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, common.NewInvalidParamError("invalid param(idToken verify)", err)
	}

	user, err := u.getByUID(ctx, token.UID)
	if err != nil {
		return nil, common.NewUnknownError("get user fail", err)
	}

	userToken, err := u.getTokenByID(ctx, tokenId)
	if err != nil {
		return nil, common.NewUnknownError("get token fail", err)
	}

	if user == nil {
		newUser := domain.User{
			UID:   token.UID,
			Email: email,
		}
		err := u.userRepository.Store(ctx, &newUser)
		user = &newUser
		if err != nil {
			return nil, common.NewDatabaseError("create user fail", err)
		}
	}

	accessToken := common.GenerateAccessToken(token.UID, userAgent)
	if userToken == nil {
		newUserToken := domain.Token{
			UserID:    user.ID,
			Token:     accessToken,
			UserAgent: &userAgent,
		}

		err := u.tokenRepository.Store(ctx, &newUserToken)
		if err != nil {
			return nil, common.NewDatabaseError("create user token fail", err)
		}
		userToken = &newUserToken
	} else {
		expireAt := domain.Time(time.Now().Add(time.Duration(common.GetConfig().AccessTokenTtl) * time.Second))
		userToken.Token = accessToken
		userToken.UserAgent = &userAgent
		userToken.ExpireAt = &expireAt
		err = u.tokenRepository.Update(ctx, userToken)
		if err != nil {
			return nil, common.NewDatabaseError("update user token fail", err)
		}
	}

	userToken.TokenID = common.EncodeHashId(userToken.ID)

	return userToken, nil
}

func (u *userService) Logout(ctx context.Context, accessToken string) error {
	token, err := u.tokenRepository.GetByAccessToken(ctx, accessToken)
	if err != nil {
		return err
	}
	expireAt := domain.Time(time.Now())
	token.ExpireAt = &expireAt

	return u.tokenRepository.Update(ctx, token)
}

func (u *userService) GetByAccessToken(ctx context.Context, accessToken string) (user *domain.User, err error) {
	token, err := u.tokenRepository.GetByAccessToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}
	return u.userRepository.GetByID(ctx, token.UserID)
}

func (u *userService) GetByID(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepository.GetByID(ctx, id)
}

func (u *userService) getByUID(ctx context.Context, uid string) (user *domain.User, err error) {
	conditions := []common.Condition{
		{
			Field: "uid",
			Op:    common.Eq,
			Val:   uid,
		},
	}
	exists, err := u.userRepository.Exists(ctx, conditions)
	if exists {
		user, err := u.userRepository.One(ctx, conditions)
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.New(fmt.Sprintf("user: not found(condition %v)", conditions))
		}
		return user, nil
	} else {
		return nil, err
	}
}

func (u *userService) getTokenByID(ctx context.Context, id int64) (token *domain.Token, err error) {
	conditions := []common.Condition{
		{
			Field: "id",
			Op:    common.Eq,
			Val:   id,
		},
	}
	exists, err := u.tokenRepository.Exists(ctx, conditions)
	if exists {
		return u.tokenRepository.GetByID(ctx, id)
	} else {
		return nil, err
	}
}

func (u *userService) Tags(ctx context.Context) ([]domain.Tag, error) {
	result, err := u.tagRepository.Fetch(ctx, nil, "", math.MaxInt32)
	if err != nil {
		common.GetLogger().Printf("user: tags - %v", err)
		return nil, err
	}
	return result.Data.([]domain.Tag), nil
}

func (u *userService) GetArchives(ctx context.Context, params domain.ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error) {
	return u.archiveRepository.Fetch(ctx, params, cursor, count)
}
