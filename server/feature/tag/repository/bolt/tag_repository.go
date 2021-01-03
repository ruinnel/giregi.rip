package bolt

import (
	"context"
	"errors"
	"fmt"
	"github.com/asdine/storm/v3"
	"github.com/asdine/storm/v3/q"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"time"
)

type tagRepository struct {
	Conn *storm.DB
}

func NewTagRepository(conn *storm.DB) domain.TagRepository {
	return &tagRepository{Conn: conn}
}

func (r tagRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (*common.FetchResult, error) {
	matchers := common.ConditionsToMatchers(conditions)
	queries := r.Conn.Select(matchers...)
	total, err := queries.Count(&domain.Tag{})
	if err != nil {
		return nil, err
	}
	offset, err := common.DecodeCursor(conditions, cursor)
	if err != nil {
		return nil, err
	}

	if offset != nil {
		matchers = append(matchers, q.Gt(domain.TagField.CreatedAt.Name, offset))
		queries = r.Conn.Select(matchers...)
	}
	queries = queries.OrderBy(domain.TagField.CreatedAt.Name).Reverse().Limit(count)

	var tags []domain.Tag
	err = queries.Find(&tags)
	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	var lastCreatedAt time.Time
	for _, t := range tags {
		lastCreatedAt = time.Time(t.CreatedAt)
	}
	result := &common.FetchResult{
		Total:      int64(total),
		Data:       tags,
		NextCursor: common.EncodeCursor(conditions, lastCreatedAt),
	}
	return result, nil
}

func (r tagRepository) One(ctx context.Context, conditions []common.Condition) (*domain.Tag, error) {
	matchers := common.ConditionsToMatchers(conditions)

	var tags []domain.Tag
	err := r.Conn.Select(matchers...).Limit(1).Find(&tags)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	if len(tags) > 0 {
		result := tags[0]
		return &result, nil
	} else {
		return nil, errors.New(fmt.Sprintf("tag: not found(condition %v)", conditions))
	}
}

func (r tagRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.Tag{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r tagRepository) GetByID(ctx context.Context, id int64) (*domain.Tag, error) {
	var tag domain.Tag
	err := r.Conn.One(domain.TagField.ID.Name, id, &tag)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r tagRepository) Store(ctx context.Context, tag *domain.Tag) error {
	tag.ID = 0
	tag.CreatedAt = domain.Time(time.Now())
	tag.UpdatedAt = domain.Time(time.Now())
	_, err := r.store(ctx, tag)
	return err
}

func (r tagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	exists, err := r.GetByID(ctx, tag.ID)
	if err != nil {
		return err
	}
	exists.Name = tag.Name
	exists.UserID = tag.UserID
	exists.UpdatedAt = domain.Time(time.Now())
	return r.Conn.Update(exists)
}

func (r tagRepository) Delete(ctx context.Context, id int64) error {
	return r.Conn.DeleteStruct(&domain.Tag{ID: id})
}

func (r tagRepository) ExistsMapping(ctx context.Context, archiveId int64, tagId int64) (bool, error) {
	conditions := []common.Condition{
		{
			Field: domain.ArchiveTagMappingField.ArchiveID,
			Op:    common.Eq,
			Val:   archiveId,
		},
		{
			Field: domain.ArchiveTagMappingField.TagID,
			Op:    common.Eq,
			Val:   tagId,
		},
	}
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.ArchiveTagMapping{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r tagRepository) AddMapping(ctx context.Context, archiveId int64, tagId int64) error {
	mapping := &domain.ArchiveTagMapping{
		TagID:     tagId,
		ArchiveID: archiveId,
	}
	return r.Conn.Save(mapping)
}

func (r tagRepository) store(ctx context.Context, tag *domain.Tag) (*domain.Tag, error) {
	err := r.Conn.Save(tag)
	if err != nil {
		return nil, err
	}
	return tag, nil
}
