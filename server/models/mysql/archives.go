// Code generated by SQLBoiler 4.2.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Archive is an object representing the database table.
type Archive struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	WebPageID int64     `boil:"web_page_id" json:"web_page_id" toml:"web_page_id" yaml:"web_page_id"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Memo      string    `boil:"memo" json:"memo" toml:"memo" yaml:"memo"`
	Title     string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	Status    string    `boil:"status" json:"status" toml:"status" yaml:"status"`
	JobID     string    `boil:"job_id" json:"job_id" toml:"job_id" yaml:"job_id"`
	WaybackID string    `boil:"wayback_id" json:"wayback_id" toml:"wayback_id" yaml:"wayback_id"`
	Summary   string    `boil:"summary" json:"summary" toml:"summary" yaml:"summary"`
	Public    bool      `boil:"public" json:"public" toml:"public" yaml:"public"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *archiveR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L archiveL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ArchiveColumns = struct {
	ID        string
	WebPageID string
	UserID    string
	Memo      string
	Title     string
	Status    string
	JobID     string
	WaybackID string
	Summary   string
	Public    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	WebPageID: "web_page_id",
	UserID:    "user_id",
	Memo:      "memo",
	Title:     "title",
	Status:    "status",
	JobID:     "job_id",
	WaybackID: "wayback_id",
	Summary:   "summary",
	Public:    "public",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperbool struct{ field string }

func (w whereHelperbool) EQ(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperbool) NEQ(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperbool) LT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperbool) LTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperbool) GT(x bool) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperbool) GTE(x bool) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var ArchiveWhere = struct {
	ID        whereHelperint64
	WebPageID whereHelperint64
	UserID    whereHelperint64
	Memo      whereHelperstring
	Title     whereHelperstring
	Status    whereHelperstring
	JobID     whereHelperstring
	WaybackID whereHelperstring
	Summary   whereHelperstring
	Public    whereHelperbool
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`archives`.`id`"},
	WebPageID: whereHelperint64{field: "`archives`.`web_page_id`"},
	UserID:    whereHelperint64{field: "`archives`.`user_id`"},
	Memo:      whereHelperstring{field: "`archives`.`memo`"},
	Title:     whereHelperstring{field: "`archives`.`title`"},
	Status:    whereHelperstring{field: "`archives`.`status`"},
	JobID:     whereHelperstring{field: "`archives`.`job_id`"},
	WaybackID: whereHelperstring{field: "`archives`.`wayback_id`"},
	Summary:   whereHelperstring{field: "`archives`.`summary`"},
	Public:    whereHelperbool{field: "`archives`.`public`"},
	CreatedAt: whereHelpertime_Time{field: "`archives`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`archives`.`updated_at`"},
}

// ArchiveRels is where relationship names are stored.
var ArchiveRels = struct {
	User               string
	WebPage            string
	ArchiveTagMappings string
}{
	User:               "User",
	WebPage:            "WebPage",
	ArchiveTagMappings: "ArchiveTagMappings",
}

// archiveR is where relationships are stored.
type archiveR struct {
	User               *User                  `boil:"User" json:"User" toml:"User" yaml:"User"`
	WebPage            *WebPage               `boil:"WebPage" json:"WebPage" toml:"WebPage" yaml:"WebPage"`
	ArchiveTagMappings ArchiveTagMappingSlice `boil:"ArchiveTagMappings" json:"ArchiveTagMappings" toml:"ArchiveTagMappings" yaml:"ArchiveTagMappings"`
}

// NewStruct creates a new relationship struct
func (*archiveR) NewStruct() *archiveR {
	return &archiveR{}
}

// archiveL is where Load methods for each relationship are stored.
type archiveL struct{}

var (
	archiveAllColumns            = []string{"id", "web_page_id", "user_id", "memo", "title", "status", "job_id", "wayback_id", "summary", "public", "created_at", "updated_at"}
	archiveColumnsWithoutDefault = []string{"web_page_id", "user_id", "memo", "title", "status", "job_id", "wayback_id", "summary"}
	archiveColumnsWithDefault    = []string{"id", "public", "created_at", "updated_at"}
	archivePrimaryKeyColumns     = []string{"id"}
)

type (
	// ArchiveSlice is an alias for a slice of pointers to Archive.
	// This should generally be used opposed to []Archive.
	ArchiveSlice []*Archive

	archiveQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	archiveType                 = reflect.TypeOf(&Archive{})
	archiveMapping              = queries.MakeStructMapping(archiveType)
	archivePrimaryKeyMapping, _ = queries.BindMapping(archiveType, archiveMapping, archivePrimaryKeyColumns)
	archiveInsertCacheMut       sync.RWMutex
	archiveInsertCache          = make(map[string]insertCache)
	archiveUpdateCacheMut       sync.RWMutex
	archiveUpdateCache          = make(map[string]updateCache)
	archiveUpsertCacheMut       sync.RWMutex
	archiveUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single archive record from the query.
func (q archiveQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Archive, error) {
	o := &Archive{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: failed to execute a one query for archives")
	}

	return o, nil
}

// All returns all Archive records from the query.
func (q archiveQuery) All(ctx context.Context, exec boil.ContextExecutor) (ArchiveSlice, error) {
	var o []*Archive

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to assign all query results to Archive slice")
	}

	return o, nil
}

// Count returns the count of all Archive records in the query.
func (q archiveQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to count archives rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q archiveQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mysql: failed to check if archives exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Archive) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// WebPage pointed to by the foreign key.
func (o *Archive) WebPage(mods ...qm.QueryMod) webPageQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.WebPageID),
	}

	queryMods = append(queryMods, mods...)

	query := WebPages(queryMods...)
	queries.SetFrom(query.Query, "`web_pages`")

	return query
}

// ArchiveTagMappings retrieves all the archive_tag_mapping's ArchiveTagMappings with an executor.
func (o *Archive) ArchiveTagMappings(mods ...qm.QueryMod) archiveTagMappingQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`archive_tag_mappings`.`archive_id`=?", o.ID),
	)

	query := ArchiveTagMappings(queryMods...)
	queries.SetFrom(query.Query, "`archive_tag_mappings`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`archive_tag_mappings`.*"})
	}

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (archiveL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArchive interface{}, mods queries.Applicator) error {
	var slice []*Archive
	var object *Archive

	if singular {
		object = maybeArchive.(*Archive)
	} else {
		slice = *maybeArchive.(*[]*Archive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &archiveR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &archiveR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Archives = append(foreign.R.Archives, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Archives = append(foreign.R.Archives, local)
				break
			}
		}
	}

	return nil
}

// LoadWebPage allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (archiveL) LoadWebPage(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArchive interface{}, mods queries.Applicator) error {
	var slice []*Archive
	var object *Archive

	if singular {
		object = maybeArchive.(*Archive)
	} else {
		slice = *maybeArchive.(*[]*Archive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &archiveR{}
		}
		args = append(args, object.WebPageID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &archiveR{}
			}

			for _, a := range args {
				if a == obj.WebPageID {
					continue Outer
				}
			}

			args = append(args, obj.WebPageID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`web_pages`),
		qm.WhereIn(`web_pages.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load WebPage")
	}

	var resultSlice []*WebPage
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice WebPage")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for web_pages")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for web_pages")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.WebPage = foreign
		if foreign.R == nil {
			foreign.R = &webPageR{}
		}
		foreign.R.Archives = append(foreign.R.Archives, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.WebPageID == foreign.ID {
				local.R.WebPage = foreign
				if foreign.R == nil {
					foreign.R = &webPageR{}
				}
				foreign.R.Archives = append(foreign.R.Archives, local)
				break
			}
		}
	}

	return nil
}

// LoadArchiveTagMappings allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (archiveL) LoadArchiveTagMappings(ctx context.Context, e boil.ContextExecutor, singular bool, maybeArchive interface{}, mods queries.Applicator) error {
	var slice []*Archive
	var object *Archive

	if singular {
		object = maybeArchive.(*Archive)
	} else {
		slice = *maybeArchive.(*[]*Archive)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &archiveR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &archiveR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`archive_tag_mappings`),
		qm.WhereIn(`archive_tag_mappings.archive_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load archive_tag_mappings")
	}

	var resultSlice []*ArchiveTagMapping
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice archive_tag_mappings")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on archive_tag_mappings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for archive_tag_mappings")
	}

	if singular {
		object.R.ArchiveTagMappings = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &archiveTagMappingR{}
			}
			foreign.R.Archive = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.ArchiveID {
				local.R.ArchiveTagMappings = append(local.R.ArchiveTagMappings, foreign)
				if foreign.R == nil {
					foreign.R = &archiveTagMappingR{}
				}
				foreign.R.Archive = local
				break
			}
		}
	}

	return nil
}

// SetUser of the archive to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Archives.
func (o *Archive) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `archives` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &archiveR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Archives: ArchiveSlice{o},
		}
	} else {
		related.R.Archives = append(related.R.Archives, o)
	}

	return nil
}

// SetWebPage of the archive to the related item.
// Sets o.R.WebPage to related.
// Adds o to related.R.Archives.
func (o *Archive) SetWebPage(ctx context.Context, exec boil.ContextExecutor, insert bool, related *WebPage) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `archives` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"web_page_id"}),
		strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.WebPageID = related.ID
	if o.R == nil {
		o.R = &archiveR{
			WebPage: related,
		}
	} else {
		o.R.WebPage = related
	}

	if related.R == nil {
		related.R = &webPageR{
			Archives: ArchiveSlice{o},
		}
	} else {
		related.R.Archives = append(related.R.Archives, o)
	}

	return nil
}

// AddArchiveTagMappings adds the given related objects to the existing relationships
// of the archive, optionally inserting them as new records.
// Appends related to o.R.ArchiveTagMappings.
// Sets related.R.Archive appropriately.
func (o *Archive) AddArchiveTagMappings(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*ArchiveTagMapping) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ArchiveID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `archive_tag_mappings` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"archive_id"}),
				strmangle.WhereClause("`", "`", 0, archiveTagMappingPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ArchiveID = o.ID
		}
	}

	if o.R == nil {
		o.R = &archiveR{
			ArchiveTagMappings: related,
		}
	} else {
		o.R.ArchiveTagMappings = append(o.R.ArchiveTagMappings, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &archiveTagMappingR{
				Archive: o,
			}
		} else {
			rel.R.Archive = o
		}
	}
	return nil
}

// Archives retrieves all the records using an executor.
func Archives(mods ...qm.QueryMod) archiveQuery {
	mods = append(mods, qm.From("`archives`"))
	return archiveQuery{NewQuery(mods...)}
}

// FindArchive retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindArchive(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Archive, error) {
	archiveObj := &Archive{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `archives` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, archiveObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: unable to select from archives")
	}

	return archiveObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Archive) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no archives provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	archiveInsertCacheMut.RLock()
	cache, cached := archiveInsertCache[key]
	archiveInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			archiveAllColumns,
			archiveColumnsWithDefault,
			archiveColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(archiveType, archiveMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `archives` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `archives` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `archives` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "mysql: unable to insert into archives")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == archiveMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for archives")
	}

CacheNoHooks:
	if !cached {
		archiveInsertCacheMut.Lock()
		archiveInsertCache[key] = cache
		archiveInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Archive.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Archive) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	archiveUpdateCacheMut.RLock()
	cache, cached := archiveUpdateCache[key]
	archiveUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			archiveAllColumns,
			archivePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("mysql: unable to update archives, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `archives` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, append(wl, archivePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update archives row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by update for archives")
	}

	if !cached {
		archiveUpdateCacheMut.Lock()
		archiveUpdateCache[key] = cache
		archiveUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q archiveQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all for archives")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected for archives")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ArchiveSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("mysql: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `archives` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all in archive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected all in update all archive")
	}
	return rowsAff, nil
}

var mySQLArchiveUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Archive) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no archives provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(archiveColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLArchiveUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	archiveUpsertCacheMut.RLock()
	cache, cached := archiveUpsertCache[key]
	archiveUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			archiveAllColumns,
			archiveColumnsWithDefault,
			archiveColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			archiveAllColumns,
			archivePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("mysql: unable to upsert archives, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "archives", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `archives` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(archiveType, archiveMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(archiveType, archiveMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "mysql: unable to upsert for archives")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = int64(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == archiveMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(archiveType, archiveMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to retrieve unique values for archives")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for archives")
	}

CacheNoHooks:
	if !cached {
		archiveUpsertCacheMut.Lock()
		archiveUpsertCache[key] = cache
		archiveUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Archive record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Archive) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mysql: no Archive provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), archivePrimaryKeyMapping)
	sql := "DELETE FROM `archives` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete from archives")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by delete for archives")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q archiveQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mysql: no archiveQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from archives")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for archives")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ArchiveSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `archives` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from archive slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for archives")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Archive) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindArchive(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ArchiveSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ArchiveSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), archivePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `archives`.* FROM `archives` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, archivePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to reload all in ArchiveSlice")
	}

	*o = slice

	return nil
}

// ArchiveExists checks if the Archive row exists.
func ArchiveExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `archives` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mysql: unable to check if archives exists")
	}

	return exists, nil
}
