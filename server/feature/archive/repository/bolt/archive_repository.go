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

type archiveRepository struct {
	Conn *storm.DB
}

func NewArchiveRepository(conn *storm.DB) domain.ArchiveRepository {
	return &archiveRepository{Conn: conn}
}

func (r archiveRepository) Fetch(ctx context.Context, params domain.ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error) {
	var conditions []common.Condition
	if params.UserID > 0 {
		conditions = append(conditions, common.Condition{Field: domain.ArchiveField.UserID, Op: common.Eq, Val: params.UserID})
	}
	if params.TagID > 0 {
		conditions = append(conditions, common.Condition{Field: domain.ArchiveField.TagIDs, Op: common.In, Val: params.TagID})
	}
	matchers := common.ConditionsToMatchers(conditions)

	if len(params.Keyword) > 0 {
		matchers = append(matchers,
			q.Re(domain.ArchiveField.Title.Name, fmt.Sprintf(`^.*%s.*$`, params.Keyword)),
			q.Re(domain.ArchiveField.Memo.Name, fmt.Sprintf(`^.*%s.*$`, params.Keyword)),
			q.Re(domain.ArchiveField.Summary.Name, fmt.Sprintf(`^.*%s.*$`, params.Keyword)),
		)
	}

	//queries = append(queries,
	//	InnerJoin("web_pages w ON w.id = archives.web_page_id"),
	//	Load(mysql.ArchiveRels.ArchiveTagMappings),
	//	Load(mysql.ArchiveRels.WebPage),
	//)
	//if params.TagID > 0 {
	//	queries = append(queries, InnerJoin("archive_tag_mappings tm ON tm.archive_id = archives.id AND tm.tag_id = ?", params.TagID))
	//}
	queries := r.Conn.Select(matchers...)
	total, err := queries.Count(&domain.Archive{})
	if err != nil {
		return nil, err
	}
	offset, err := common.DecodeCursor(conditions, cursor)
	if err != nil {
		return nil, err
	}

	if offset != nil {
		matchers = append(matchers, q.Gt(domain.ArchiveField.CreatedAt.Name, offset))
		queries = r.Conn.Select(matchers...)
	}
	queries = queries.OrderBy(domain.ArchiveField.CreatedAt.Name).Limit(count)
	var archives []domain.Archive
	err = queries.Find(&archives)
	if err != nil && err != storm.ErrNotFound {
		return nil, err
	}

	data := make([]domain.Archive, len(archives))
	var lastCreatedAt time.Time
	for idx, archive := range archives {
		var webPage domain.WebPage
		err = r.Conn.One(domain.ArchiveField.ID.Name, archive.WebPageID, &webPage)
		if err != nil {
			return nil, err
		}
		archive.WebPage = &webPage
		data[idx] = archive
		lastCreatedAt = time.Time(archive.CreatedAt)
	}
	result := &common.FetchResult{
		Total:      int64(total),
		Data:       data,
		NextCursor: common.EncodeCursor(conditions, lastCreatedAt),
	}
	return result, nil
}

func (r archiveRepository) One(ctx context.Context, conditions []common.Condition) (*domain.Archive, error) {
	matchers := common.ConditionsToMatchers(conditions)

	var archives []domain.Archive
	err := r.Conn.Select(matchers...).Limit(1).Find(&archives)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	if len(archives) > 0 {
		result := archives[0]
		return &result, nil
	} else {
		return nil, errors.New(fmt.Sprintf("archive: not found(condition %v)", conditions))
	}
}

func (r archiveRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	matchers := common.ConditionsToMatchers(conditions)
	count, err := r.Conn.Select(matchers...).Count(&domain.Archive{})
	if err != nil {
		return false, err
	} else {
		return count > 0, nil
	}
}

func (r archiveRepository) GetByID(ctx context.Context, id int64) (*domain.Archive, error) {
	var archive domain.Archive
	err := r.Conn.One(domain.ArchiveField.ID.Name, id, &archive)
	if err == storm.ErrNotFound {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	return &archive, nil
}

func (r archiveRepository) Store(ctx context.Context, archive *domain.Archive) error {
	archive.ID = 0
	err := r.Conn.Save(archive)
	if err != nil {
		return err
	}
	return nil
}

func (r archiveRepository) Update(ctx context.Context, archive *domain.Archive) error {
	exists, err := r.GetByID(ctx, archive.ID)
	if err != nil {
		return err
	}

	exists.WebPageID = archive.WebPageID
	exists.UserID = archive.UserID
	exists.WaybackID = archive.WaybackID
	exists.JobID = archive.JobID
	exists.Status = archive.Status
	exists.Memo = archive.Memo
	exists.Title = archive.Title
	exists.Summary = archive.Summary
	exists.Public = archive.Public
	return r.Conn.Update(exists)
}

func (r archiveRepository) Delete(ctx context.Context, id int64) error {
	return r.Conn.DeleteStruct(&domain.Archive{ID: id})
}
