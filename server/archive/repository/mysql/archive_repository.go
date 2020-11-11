package mysql

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/ruinnel/giregi.rip-server/common"
	"github.com/ruinnel/giregi.rip-server/domain"
	"github.com/ruinnel/giregi.rip-server/models/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"time"
)

type archiveRepository struct {
	Conn *sql.DB
}

func (r archiveRepository) newArchiveDomain(archive *mysql.Archive) domain.Archive {
	summary := new([]map[string]interface{})
	_ = json.Unmarshal([]byte(archive.Summary), summary)

	var tagIds []int64
	if archive.R != nil && archive.R.ArchiveTagMappings != nil && len(archive.R.ArchiveTagMappings) > 0 {
		for _, m := range archive.R.ArchiveTagMappings {
			tagIds = append(tagIds, m.TagID)
		}
	}

	var webPage *domain.WebPage
	if archive.R != nil && archive.R.WebPage != nil {
		webPage = &domain.WebPage{
			ID:        archive.R.WebPage.ID,
			SiteID:    archive.R.WebPage.SiteID,
			URL:       archive.R.WebPage.URL,
			Title:     &archive.R.WebPage.Title.String,
			CreatedAt: domain.Time(archive.R.WebPage.CreatedAt),
			UpdatedAt: domain.Time(archive.R.WebPage.UpdatedAt),
		}
	}
	return domain.Archive{
		ID:        archive.ID,
		WebPageID: archive.WebPageID,
		UserID:    archive.UserID,
		WaybackID: archive.WaybackID,
		Memo:      archive.Memo,
		Status:    archive.Status,
		JobID:     archive.JobID,
		Summary:   *summary,
		Public:    archive.Public,
		CreatedAt: domain.Time(archive.CreatedAt),
		UpdatedAt: domain.Time(archive.UpdatedAt),
		TagIDs:    tagIds,
		WebPage:   webPage,
	}
}

func (r archiveRepository) Fetch(ctx context.Context, params domain.ArchiveFetchParams, cursor string, count int) (*common.FetchResult, error) {
	var conditions []common.Condition
	if params.UserID > 0 {
		conditions = append(conditions, common.Condition{Field: "user_id", Op: common.Eq, Val: params.UserID})
	}
	if params.TagID > 0 {
		conditions = append(conditions, common.Condition{Field: "tm.tag_id", Op: common.Eq, Val: params.TagID})
	}
	queries := common.ConditionsToQueries(conditions)

	if len(params.Keyword) > 0 {
		queries = append(queries, Expr(
			Or("w.url LIKE ?", fmt.Sprintf("%%%v%%", params.Keyword)),
			Or("w.title LIKE ?", fmt.Sprintf("%%%v%%", params.Keyword)),
			Or("memo LIKE ?", fmt.Sprintf("%%%v%%", params.Keyword)),
			Or("summary LIKE ?", fmt.Sprintf("%%%v%%", params.Keyword)),
		))
	}

	queries = append(queries,
		InnerJoin("web_pages w ON w.id = archives.web_page_id"),
		Load(mysql.ArchiveRels.ArchiveTagMappings),
		Load(mysql.ArchiveRels.WebPage),
	)
	if params.TagID > 0 {
		queries = append(queries, InnerJoin("archive_tag_mappings tm ON tm.archive_id = archives.id AND tm.tag_id = ?", params.TagID))
	}
	total, err := mysql.Archives(queries...).Count(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	offset, err := common.DecodeCursor(conditions, cursor)
	if err != nil {
		return nil, err
	}
	queries = append(queries, OrderBy("archives.created_at desc"))
	queries = append(queries, Limit(count))
	if offset != nil {
		queries = append(queries, Where("archives.created_at > ?", offset))
	}
	archives, err := mysql.Archives(queries...).All(ctx, r.Conn)
	if err != nil {
		return nil, err
	}

	data := make([]domain.Archive, len(archives))
	var lastCreatedAt time.Time
	for idx, archive := range archives {
		data[idx] = r.newArchiveDomain(archive)
		lastCreatedAt = archive.CreatedAt
	}
	result := &common.FetchResult{
		Total:      total,
		Data:       data,
		NextCursor: common.EncodeCursor(conditions, lastCreatedAt),
	}
	return result, nil
}

func (r archiveRepository) One(ctx context.Context, conditions []common.Condition) (*domain.Archive, error) {
	queries := common.ConditionsToQueries(conditions)

	archive, err := mysql.Archives(queries...).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	result := r.newArchiveDomain(archive)
	return &result, nil
}

func (r archiveRepository) Exists(ctx context.Context, conditions []common.Condition) (bool, error) {
	queries := common.ConditionsToQueries(conditions)
	return mysql.Archives(queries...).Exists(ctx, r.Conn)
}

func (r archiveRepository) GetByID(ctx context.Context, id int64) (*domain.Archive, error) {
	user, err := mysql.Archives(Where("id = ?", id)).One(ctx, r.Conn)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	} else {
		archive := r.newArchiveDomain(user)
		return &archive, nil
	}
}

func (r archiveRepository) Store(ctx context.Context, archive *domain.Archive) error {
	summary, err := json.Marshal(archive.Summary)
	if err != nil {
		return err
	}
	newArchive := mysql.Archive{
		WebPageID: archive.WebPageID,
		UserID:    archive.UserID,
		Memo:      archive.Memo,
		Status:    archive.Status,
		JobID:     archive.JobID,
		WaybackID: archive.WaybackID,
		Summary:   string(summary),
		Public:    archive.Public,
	}

	err = newArchive.Insert(ctx, r.Conn, boil.Infer())
	if err != nil {
		return err
	}
	archive.ID = newArchive.ID
	return nil
}

func (r archiveRepository) Update(ctx context.Context, archive *domain.Archive) error {
	exists, err := mysql.Archives(Where("id = ?", archive.ID)).One(ctx, r.Conn)
	if err != nil {
		return err
	}

	summary, err := json.Marshal(archive.Summary)
	if err != nil {
		return err
	}

	exists.WebPageID = archive.WebPageID
	exists.UserID = archive.UserID
	exists.WaybackID = archive.WaybackID
	exists.JobID = archive.JobID
	exists.Status = archive.Status
	exists.Memo = archive.Memo
	exists.Summary = string(summary)
	exists.Public = archive.Public
	_, err = exists.Update(ctx, r.Conn, boil.Infer())
	return err
}

func (r archiveRepository) Delete(ctx context.Context, id int64) error {
	_, err := mysql.Archives(Where("id = ?", id)).DeleteAll(ctx, r.Conn)
	return err
}

func NewArchiveRepository(conn *sql.DB) domain.ArchiveRepository {
	return &archiveRepository{Conn: conn}
}
