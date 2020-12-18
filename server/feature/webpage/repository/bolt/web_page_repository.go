package bolt

import (
	"context"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
)

type webPageRepository struct {
	Conn *storm.DB
}

func NewWebPageRepository(conn *storm.DB) domain.WebPageRepository {
	return &webPageRepository{Conn: conn}
}

func (r webPageRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []domain.WebPage, nextCursor string, err error) {
	panic("implement me")
}

func (r webPageRepository) One(ctx context.Context, conditions []common.Condition) (*domain.WebPage, error) {
	matchers := common.ConditionsToMatchers(conditions)

	var webPages []domain.WebPage
	err := r.Conn.Select(matchers...).Limit(1).Find(&webPages)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	if len(webPages) > 0 {
		result := webPages[0]
		return &result, nil
	} else {
		return nil, errors.New(fmt.Sprintf("webPage: not found(condition %v)", conditions))
	}
}

func (r webPageRepository) GetByURL(ctx context.Context, url string) (*domain.WebPage, error) {
	var webPage domain.WebPage
	err := r.Conn.One(domain.WebPageField.URL.Name, url, &webPage)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &webPage, nil
}

func (r webPageRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.WebPage{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r webPageRepository) GetByID(ctx context.Context, id int64) (*domain.WebPage, error) {
	var webPage domain.WebPage
	err := r.Conn.One(domain.WebPageField.ID.Name, id, &webPage)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &webPage, nil
}

func (r webPageRepository) Store(ctx context.Context, webPage *domain.WebPage) error {
	err := r.Conn.Save(webPage)
	if err != nil {
		return err
	}
	return nil
}

func (r webPageRepository) Update(ctx context.Context, webPage *domain.WebPage) error {
	exists, err := r.GetByID(ctx, webPage.ID)
	if err != nil {
		return err
	}
	exists.SiteID = webPage.SiteID
	exists.URL = webPage.URL
	exists.Title = webPage.Title
	return r.Conn.Update(exists)
}

func (r webPageRepository) Delete(ctx context.Context, id int64) error {
	return r.Conn.DeleteStruct(&domain.WebPage{ID: id})
}
