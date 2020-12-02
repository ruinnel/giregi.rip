package mysql

import (
	"context"
	"database/sql"
	"errors"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/models/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type siteRepository struct {
	Conn *sql.DB
}

func NewSiteRepository(conn *sql.DB) domain.SiteRepository {
	return &siteRepository{Conn: conn}
}

func (r siteRepository) newSiteDomain(site *mysql.Site) domain.Site {
	return domain.Site{
		ID:        site.ID,
		Name:      site.Name,
		Domain:    site.Domain,
		CreatedAt: domain.Time(site.CreatedAt),
		UpdatedAt: domain.Time(site.UpdatedAt),
	}
}

func (r siteRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []domain.Site, nextCursor string, err error) {
	panic("implement me")
}

func (r siteRepository) One(ctx context.Context, conditions []common.Condition) (*domain.Site, error) {
	queries := common.ConditionsToQueries(conditions)

	site, err := mysql.Sites(queries...).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	result := r.newSiteDomain(site)
	return &result, nil
}

func (r siteRepository) GetByDomain(ctx context.Context, domain string) (*domain.Site, error) {
	site, err := mysql.Sites(Where("domain = ?", domain)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if site == nil {
		return nil, errors.New("site: invalid domain")
	}
	result := r.newSiteDomain(site)
	return &result, nil
}

func (r siteRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	queries := common.ConditionsToQueries(conditions)
	return mysql.Sites(queries...).Exists(ctx, r.Conn)
}

func (r siteRepository) GetByID(ctx context.Context, id int64) (*domain.Site, error) {
	tk, err := mysql.Sites(Where("id = ?", id)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if tk == nil {
		return nil, nil
	} else {
		site := r.newSiteDomain(tk)
		return &site, nil
	}
}

func (r siteRepository) Store(ctx context.Context, site *domain.Site) error {
	newSite := mysql.Site{
		Name:   site.Name,
		Domain: site.Domain,
	}

	err := newSite.Insert(ctx, r.Conn, boil.Infer())
	if err != nil {
		return err
	}
	site.ID = newSite.ID
	return nil
}

func (r siteRepository) Update(ctx context.Context, site *domain.Site) error {
	exists, err := mysql.Sites(Where("id = ?", site.ID)).One(ctx, r.Conn)
	if err != nil {
		return err
	}
	exists.Name = site.Name
	exists.Domain = site.Domain
	_, err = exists.Update(ctx, r.Conn, boil.Infer())
	return err
}

func (r siteRepository) Delete(ctx context.Context, id int64) error {
	_, err := mysql.Sites(Where("id = ?", id)).DeleteAll(ctx, r.Conn)
	return err
}
