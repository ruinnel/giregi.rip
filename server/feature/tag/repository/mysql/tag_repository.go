package mysql

import (
	"context"
	"database/sql"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/models/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type tagRepository struct {
	Conn *sql.DB
}

func NewTagRepository(conn *sql.DB) domain.TagRepository {
	return &tagRepository{Conn: conn}
}

func (r tagRepository) newTagDomain(tag *mysql.Tag) domain.Tag {
	return domain.Tag{
		ID:        tag.ID,
		Name:      tag.Name,
		UserID:    tag.UserID,
		CreatedAt: domain.Time(tag.CreatedAt),
		UpdatedAt: domain.Time(tag.UpdatedAt),
	}
}

func (r tagRepository) Fetch(ctx context.Context, conditions []common.Condition, cursor string, count int) (*common.FetchResult, error) {
	queries := common.ConditionsToQueries(conditions)
	total, err := mysql.Tags(queries...).Count(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	offset, err := common.DecodeCursor(conditions, cursor)
	if err != nil {
		return nil, err
	}
	queries = append(queries, OrderBy("created_at desc"))
	queries = append(queries, Limit(count))
	if offset != nil {
		queries = append(queries, Where("created_at = ?", offset))
	}
	tags, err := mysql.Tags(queries...).All(ctx, r.Conn)
	if err != nil {
		return nil, err
	}

	data := make([]domain.Tag, len(tags))
	var lastCreatedAt time.Time
	for idx, t := range tags {
		data[idx] = r.newTagDomain(t)
		lastCreatedAt = t.CreatedAt
	}
	result := &common.FetchResult{
		Total:      total,
		Data:       data,
		NextCursor: common.EncodeCursor(conditions, lastCreatedAt),
	}
	return result, nil
}

func (r tagRepository) One(ctx context.Context, conditions []common.Condition) (*domain.Tag, error) {
	queries := common.ConditionsToQueries(conditions)

	tag, err := mysql.Tags(queries...).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	result := r.newTagDomain(tag)
	return &result, nil
}

func (r tagRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	queries := common.ConditionsToQueries(conditions)
	return mysql.Tags(queries...).Exists(ctx, r.Conn)
}

func (r tagRepository) GetByID(ctx context.Context, id int64) (*domain.Tag, error) {
	tag, err := r.getByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, nil
	} else {
		t := r.newTagDomain(tag)
		return &t, nil
	}
}

func (r tagRepository) Store(ctx context.Context, tag *domain.Tag) error {
	_, err := r.store(ctx, tag)
	return err
}

func (r tagRepository) Update(ctx context.Context, tag *domain.Tag) error {
	exists, err := mysql.Tags(Where("id = ?", tag.ID)).One(ctx, r.Conn)
	if err != nil {
		return err
	}
	exists.Name = tag.Name
	exists.UserID = tag.UserID
	_, err = exists.Update(ctx, r.Conn, boil.Infer())
	return err
}

func (r tagRepository) Delete(ctx context.Context, id int64) error {
	_, err := mysql.Tags(Where("id = ?", id)).DeleteAll(ctx, r.Conn)
	return err
}

func (r tagRepository) ExistsMapping(ctx context.Context, archiveId int64, tagId int64) (bool, error) {
	return mysql.ArchiveTagMappings(
		Where("archive_id = ?", archiveId),
		Where("tag_id = ?", tagId),
	).Exists(ctx, r.Conn)
}

func (r tagRepository) AddMapping(ctx context.Context, archiveId int64, tagId int64) error {
	mapping := &mysql.ArchiveTagMapping{
		TagID:     tagId,
		ArchiveID: archiveId,
	}
	return mapping.Insert(ctx, r.Conn, boil.Infer())
}

func (r tagRepository) store(ctx context.Context, tag *domain.Tag) (*mysql.Tag, error) {
	newTag := &mysql.Tag{
		Name:   tag.Name,
		UserID: tag.UserID,
	}

	err := newTag.Insert(ctx, r.Conn, boil.Infer())
	if err != nil {
		return nil, err
	}
	tag.ID = newTag.ID
	return newTag, nil
}

func (r tagRepository) getByID(ctx context.Context, id int64) (*mysql.Tag, error) {
	tag, err := mysql.Tags(Where("id = ?", id)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if tag == nil {
		return nil, nil
	} else {
		return tag, nil
	}
}
