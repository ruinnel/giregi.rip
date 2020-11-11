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
)

type webPageRepository struct {
	Conn *sql.DB
}

func NewWebPageRepository(conn *sql.DB) domain.WebPageRepository {
	return &webPageRepository{Conn: conn}
}

func (r webPageRepository) newWebPageDomain(webPage *mysql.WebPage) domain.WebPage {
	return domain.WebPage{
		ID:        webPage.ID,
		SiteID:    webPage.SiteID,
		URL:       webPage.URL,
		Title:     &webPage.Title.String,
		CreatedAt: domain.Time(webPage.CreatedAt),
		UpdatedAt: domain.Time(webPage.UpdatedAt),
	}
}

func (r webPageRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []domain.WebPage, nextCursor string, err error) {
	panic("implement me")
}

func (r webPageRepository) One(ctx context.Context, conditions []common.Condition) (*domain.WebPage, error) {
	queries := common.ConditionsToQueries(conditions)

	webPage, err := mysql.WebPages(queries...).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	result := r.newWebPageDomain(webPage)
	return &result, nil
}

func (r webPageRepository) GetByURL(ctx context.Context, url string) (*domain.WebPage, error) {
	webPage, err := mysql.WebPages(Where("url = ?", url)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if webPage == nil {
		return nil, errors.New("webPage: invalid url")
	}
	result := r.newWebPageDomain(webPage)
	return &result, nil
}

func (r webPageRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	queries := common.ConditionsToQueries(conditions)
	return mysql.WebPages(queries...).Exists(ctx, r.Conn)
}

func (r webPageRepository) GetByID(ctx context.Context, id int64) (*domain.WebPage, error) {
	tk, err := mysql.WebPages(Where("id = ?", id)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if tk == nil {
		return nil, nil
	} else {
		webPage := r.newWebPageDomain(tk)
		return &webPage, nil
	}
}

func (r webPageRepository) Store(ctx context.Context, webPage *domain.WebPage) error {
	newWebPage := mysql.WebPage{
		SiteID: webPage.SiteID,
		URL:    webPage.URL,
		Title:  null.StringFromPtr(webPage.Title),
	}

	err := newWebPage.Insert(ctx, r.Conn, boil.Infer())
	if err != nil {
		return err
	}
	webPage.ID = newWebPage.ID
	return nil
}

func (r webPageRepository) Update(ctx context.Context, webPage *domain.WebPage) error {
	exists, err := mysql.WebPages(Where("id = ?", webPage.ID)).One(ctx, r.Conn)
	if err != nil {
		return err
	}
	exists.SiteID = webPage.SiteID
	exists.URL = webPage.URL
	exists.Title = null.StringFromPtr(webPage.Title)
	_, err = exists.Update(ctx, r.Conn, boil.Infer())
	return err
}

func (r webPageRepository) Delete(ctx context.Context, id int64) error {
	_, err := mysql.WebPages(Where("id = ?", id)).DeleteAll(ctx, r.Conn)
	return err
}
