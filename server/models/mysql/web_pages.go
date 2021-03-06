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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// WebPage is an object representing the database table.
type WebPage struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	SiteID    int64       `boil:"site_id" json:"site_id" toml:"site_id" yaml:"site_id"`
	URL       string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	Title     null.String `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *webPageR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L webPageL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var WebPageColumns = struct {
	ID        string
	SiteID    string
	URL       string
	Title     string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	SiteID:    "site_id",
	URL:       "url",
	Title:     "title",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var WebPageWhere = struct {
	ID        whereHelperint64
	SiteID    whereHelperint64
	URL       whereHelperstring
	Title     whereHelpernull_String
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`web_pages`.`id`"},
	SiteID:    whereHelperint64{field: "`web_pages`.`site_id`"},
	URL:       whereHelperstring{field: "`web_pages`.`url`"},
	Title:     whereHelpernull_String{field: "`web_pages`.`title`"},
	CreatedAt: whereHelpertime_Time{field: "`web_pages`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`web_pages`.`updated_at`"},
}

// WebPageRels is where relationship names are stored.
var WebPageRels = struct {
	Site     string
	Archives string
}{
	Site:     "Site",
	Archives: "Archives",
}

// webPageR is where relationships are stored.
type webPageR struct {
	Site     *Site        `boil:"Site" json:"Site" toml:"Site" yaml:"Site"`
	Archives ArchiveSlice `boil:"Archives" json:"Archives" toml:"Archives" yaml:"Archives"`
}

// NewStruct creates a new relationship struct
func (*webPageR) NewStruct() *webPageR {
	return &webPageR{}
}

// webPageL is where Load methods for each relationship are stored.
type webPageL struct{}

var (
	webPageAllColumns            = []string{"id", "site_id", "url", "title", "created_at", "updated_at"}
	webPageColumnsWithoutDefault = []string{"site_id", "url", "title"}
	webPageColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	webPagePrimaryKeyColumns     = []string{"id"}
)

type (
	// WebPageSlice is an alias for a slice of pointers to WebPage.
	// This should generally be used opposed to []WebPage.
	WebPageSlice []*WebPage

	webPageQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	webPageType                 = reflect.TypeOf(&WebPage{})
	webPageMapping              = queries.MakeStructMapping(webPageType)
	webPagePrimaryKeyMapping, _ = queries.BindMapping(webPageType, webPageMapping, webPagePrimaryKeyColumns)
	webPageInsertCacheMut       sync.RWMutex
	webPageInsertCache          = make(map[string]insertCache)
	webPageUpdateCacheMut       sync.RWMutex
	webPageUpdateCache          = make(map[string]updateCache)
	webPageUpsertCacheMut       sync.RWMutex
	webPageUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single webPage record from the query.
func (q webPageQuery) One(ctx context.Context, exec boil.ContextExecutor) (*WebPage, error) {
	o := &WebPage{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: failed to execute a one query for web_pages")
	}

	return o, nil
}

// All returns all WebPage records from the query.
func (q webPageQuery) All(ctx context.Context, exec boil.ContextExecutor) (WebPageSlice, error) {
	var o []*WebPage

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to assign all query results to WebPage slice")
	}

	return o, nil
}

// Count returns the count of all WebPage records in the query.
func (q webPageQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to count web_pages rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q webPageQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mysql: failed to check if web_pages exists")
	}

	return count > 0, nil
}

// Site pointed to by the foreign key.
func (o *WebPage) Site(mods ...qm.QueryMod) siteQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.SiteID),
	}

	queryMods = append(queryMods, mods...)

	query := Sites(queryMods...)
	queries.SetFrom(query.Query, "`sites`")

	return query
}

// Archives retrieves all the archive's Archives with an executor.
func (o *WebPage) Archives(mods ...qm.QueryMod) archiveQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`archives`.`web_page_id`=?", o.ID),
	)

	query := Archives(queryMods...)
	queries.SetFrom(query.Query, "`archives`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`archives`.*"})
	}

	return query
}

// LoadSite allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (webPageL) LoadSite(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWebPage interface{}, mods queries.Applicator) error {
	var slice []*WebPage
	var object *WebPage

	if singular {
		object = maybeWebPage.(*WebPage)
	} else {
		slice = *maybeWebPage.(*[]*WebPage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &webPageR{}
		}
		args = append(args, object.SiteID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &webPageR{}
			}

			for _, a := range args {
				if a == obj.SiteID {
					continue Outer
				}
			}

			args = append(args, obj.SiteID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`sites`),
		qm.WhereIn(`sites.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Site")
	}

	var resultSlice []*Site
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Site")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for sites")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for sites")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Site = foreign
		if foreign.R == nil {
			foreign.R = &siteR{}
		}
		foreign.R.WebPages = append(foreign.R.WebPages, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.SiteID == foreign.ID {
				local.R.Site = foreign
				if foreign.R == nil {
					foreign.R = &siteR{}
				}
				foreign.R.WebPages = append(foreign.R.WebPages, local)
				break
			}
		}
	}

	return nil
}

// LoadArchives allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (webPageL) LoadArchives(ctx context.Context, e boil.ContextExecutor, singular bool, maybeWebPage interface{}, mods queries.Applicator) error {
	var slice []*WebPage
	var object *WebPage

	if singular {
		object = maybeWebPage.(*WebPage)
	} else {
		slice = *maybeWebPage.(*[]*WebPage)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &webPageR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &webPageR{}
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
		qm.From(`archives`),
		qm.WhereIn(`archives.web_page_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load archives")
	}

	var resultSlice []*Archive
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice archives")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on archives")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for archives")
	}

	if singular {
		object.R.Archives = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &archiveR{}
			}
			foreign.R.WebPage = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.WebPageID {
				local.R.Archives = append(local.R.Archives, foreign)
				if foreign.R == nil {
					foreign.R = &archiveR{}
				}
				foreign.R.WebPage = local
				break
			}
		}
	}

	return nil
}

// SetSite of the webPage to the related item.
// Sets o.R.Site to related.
// Adds o to related.R.WebPages.
func (o *WebPage) SetSite(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Site) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `web_pages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"site_id"}),
		strmangle.WhereClause("`", "`", 0, webPagePrimaryKeyColumns),
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

	o.SiteID = related.ID
	if o.R == nil {
		o.R = &webPageR{
			Site: related,
		}
	} else {
		o.R.Site = related
	}

	if related.R == nil {
		related.R = &siteR{
			WebPages: WebPageSlice{o},
		}
	} else {
		related.R.WebPages = append(related.R.WebPages, o)
	}

	return nil
}

// AddArchives adds the given related objects to the existing relationships
// of the web_page, optionally inserting them as new records.
// Appends related to o.R.Archives.
// Sets related.R.WebPage appropriately.
func (o *WebPage) AddArchives(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Archive) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.WebPageID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `archives` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"web_page_id"}),
				strmangle.WhereClause("`", "`", 0, archivePrimaryKeyColumns),
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

			rel.WebPageID = o.ID
		}
	}

	if o.R == nil {
		o.R = &webPageR{
			Archives: related,
		}
	} else {
		o.R.Archives = append(o.R.Archives, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &archiveR{
				WebPage: o,
			}
		} else {
			rel.R.WebPage = o
		}
	}
	return nil
}

// WebPages retrieves all the records using an executor.
func WebPages(mods ...qm.QueryMod) webPageQuery {
	mods = append(mods, qm.From("`web_pages`"))
	return webPageQuery{NewQuery(mods...)}
}

// FindWebPage retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindWebPage(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*WebPage, error) {
	webPageObj := &WebPage{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `web_pages` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, webPageObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: unable to select from web_pages")
	}

	return webPageObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *WebPage) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no web_pages provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(webPageColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	webPageInsertCacheMut.RLock()
	cache, cached := webPageInsertCache[key]
	webPageInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			webPageAllColumns,
			webPageColumnsWithDefault,
			webPageColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(webPageType, webPageMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(webPageType, webPageMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `web_pages` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `web_pages` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `web_pages` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, webPagePrimaryKeyColumns))
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
		return errors.Wrap(err, "mysql: unable to insert into web_pages")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == webPageMapping["id"] {
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
		return errors.Wrap(err, "mysql: unable to populate default values for web_pages")
	}

CacheNoHooks:
	if !cached {
		webPageInsertCacheMut.Lock()
		webPageInsertCache[key] = cache
		webPageInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the WebPage.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *WebPage) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	webPageUpdateCacheMut.RLock()
	cache, cached := webPageUpdateCache[key]
	webPageUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			webPageAllColumns,
			webPagePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("mysql: unable to update web_pages, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `web_pages` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, webPagePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(webPageType, webPageMapping, append(wl, webPagePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "mysql: unable to update web_pages row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by update for web_pages")
	}

	if !cached {
		webPageUpdateCacheMut.Lock()
		webPageUpdateCache[key] = cache
		webPageUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q webPageQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all for web_pages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected for web_pages")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o WebPageSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webPagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `web_pages` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webPagePrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all in webPage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected all in update all webPage")
	}
	return rowsAff, nil
}

var mySQLWebPageUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *WebPage) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no web_pages provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(webPageColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLWebPageUniqueColumns, o)

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

	webPageUpsertCacheMut.RLock()
	cache, cached := webPageUpsertCache[key]
	webPageUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			webPageAllColumns,
			webPageColumnsWithDefault,
			webPageColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			webPageAllColumns,
			webPagePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("mysql: unable to upsert web_pages, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "web_pages", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `web_pages` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(webPageType, webPageMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(webPageType, webPageMapping, ret)
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
		return errors.Wrap(err, "mysql: unable to upsert for web_pages")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == webPageMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(webPageType, webPageMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to retrieve unique values for web_pages")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for web_pages")
	}

CacheNoHooks:
	if !cached {
		webPageUpsertCacheMut.Lock()
		webPageUpsertCache[key] = cache
		webPageUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single WebPage record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *WebPage) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mysql: no WebPage provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), webPagePrimaryKeyMapping)
	sql := "DELETE FROM `web_pages` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete from web_pages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by delete for web_pages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q webPageQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mysql: no webPageQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from web_pages")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for web_pages")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o WebPageSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webPagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `web_pages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webPagePrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from webPage slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for web_pages")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *WebPage) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindWebPage(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *WebPageSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := WebPageSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), webPagePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `web_pages`.* FROM `web_pages` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, webPagePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to reload all in WebPageSlice")
	}

	*o = slice

	return nil
}

// WebPageExists checks if the WebPage row exists.
func WebPageExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `web_pages` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mysql: unable to check if web_pages exists")
	}

	return exists, nil
}
