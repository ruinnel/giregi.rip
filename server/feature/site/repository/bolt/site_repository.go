package bolt

import (
	"context"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
)

type siteRepository struct {
	Conn *storm.DB
}

func NewSiteRepository(conn *storm.DB) domain.SiteRepository {
	return &siteRepository{Conn: conn}
}

func (r siteRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (data []domain.Site, nextCursor string, err error) {
	panic("implement me")
}

func (r siteRepository) One(ctx context.Context, conditions []common.Condition) (*domain.Site, error) {
	matchers := common.ConditionsToMatchers(conditions)

	var sites []domain.Site
	err := r.Conn.Select(matchers...).Limit(1).Find(&sites)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	if len(sites) > 0 {
		result := sites[0]
		return &result, nil
	} else {
		return nil, errors.New(fmt.Sprintf("site: not found(condition %v)", conditions))
	}
}

func (r siteRepository) GetByDomain(ctx context.Context, domainText string) (*domain.Site, error) {
	var site domain.Site
	err := r.Conn.One(domain.SiteField.Domain.Name, domainText, &site)
	if err != nil {
		return nil, err
	}
	return &site, nil
}

func (r siteRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.Archive{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r siteRepository) GetByID(ctx context.Context, id int64) (*domain.Site, error) {
	var site domain.Site
	err := r.Conn.One(domain.SiteField.ID.Name, id, &site)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &site, nil
}

func (r siteRepository) Store(ctx context.Context, site *domain.Site) error {
	err := r.Conn.Save(site)
	if err != nil {
		return err
	}
	return nil
}

func (r siteRepository) Update(ctx context.Context, site *domain.Site) error {
	exists, err := r.GetByID(ctx, site.ID)
	if err != nil {
		return err
	}
	exists.Name = site.Name
	exists.Domain = site.Domain
	return r.Conn.Update(exists)
}

func (r siteRepository) Delete(ctx context.Context, id int64) error {
	return r.Conn.DeleteStruct(&domain.Site{ID: id})
}
