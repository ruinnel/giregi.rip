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

// Token is an object representing the database table.
type Token struct {
	ID        int64       `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int64       `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Token     string      `boil:"token" json:"token" toml:"token" yaml:"token"`
	UserAgent null.String `boil:"user_agent" json:"user_agent,omitempty" toml:"user_agent" yaml:"user_agent,omitempty"`
	ExpireAt  null.Time   `boil:"expire_at" json:"expire_at,omitempty" toml:"expire_at" yaml:"expire_at,omitempty"`
	CreatedAt time.Time   `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time   `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *tokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TokenColumns = struct {
	ID        string
	UserID    string
	Token     string
	UserAgent string
	ExpireAt  string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	UserID:    "user_id",
	Token:     "token",
	UserAgent: "user_agent",
	ExpireAt:  "expire_at",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TokenWhere = struct {
	ID        whereHelperint64
	UserID    whereHelperint64
	Token     whereHelperstring
	UserAgent whereHelpernull_String
	ExpireAt  whereHelpernull_Time
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "`tokens`.`id`"},
	UserID:    whereHelperint64{field: "`tokens`.`user_id`"},
	Token:     whereHelperstring{field: "`tokens`.`token`"},
	UserAgent: whereHelpernull_String{field: "`tokens`.`user_agent`"},
	ExpireAt:  whereHelpernull_Time{field: "`tokens`.`expire_at`"},
	CreatedAt: whereHelpertime_Time{field: "`tokens`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`tokens`.`updated_at`"},
}

// TokenRels is where relationship names are stored.
var TokenRels = struct {
	User string
}{
	User: "User",
}

// tokenR is where relationships are stored.
type tokenR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*tokenR) NewStruct() *tokenR {
	return &tokenR{}
}

// tokenL is where Load methods for each relationship are stored.
type tokenL struct{}

var (
	tokenAllColumns            = []string{"id", "user_id", "token", "user_agent", "expire_at", "created_at", "updated_at"}
	tokenColumnsWithoutDefault = []string{"user_id", "token", "user_agent", "expire_at"}
	tokenColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	tokenPrimaryKeyColumns     = []string{"id"}
)

type (
	// TokenSlice is an alias for a slice of pointers to Token.
	// This should generally be used opposed to []Token.
	TokenSlice []*Token

	tokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tokenType                 = reflect.TypeOf(&Token{})
	tokenMapping              = queries.MakeStructMapping(tokenType)
	tokenPrimaryKeyMapping, _ = queries.BindMapping(tokenType, tokenMapping, tokenPrimaryKeyColumns)
	tokenInsertCacheMut       sync.RWMutex
	tokenInsertCache          = make(map[string]insertCache)
	tokenUpdateCacheMut       sync.RWMutex
	tokenUpdateCache          = make(map[string]updateCache)
	tokenUpsertCacheMut       sync.RWMutex
	tokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single token record from the query.
func (q tokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Token, error) {
	o := &Token{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: failed to execute a one query for tokens")
	}

	return o, nil
}

// All returns all Token records from the query.
func (q tokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (TokenSlice, error) {
	var o []*Token

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "mysql: failed to assign all query results to Token slice")
	}

	return o, nil
}

// Count returns the count of all Token records in the query.
func (q tokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to count tokens rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "mysql: failed to check if tokens exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *Token) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "`users`")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeToken interface{}, mods queries.Applicator) error {
	var slice []*Token
	var object *Token

	if singular {
		object = maybeToken.(*Token)
	} else {
		slice = *maybeToken.(*[]*Token)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tokenR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tokenR{}
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
		foreign.R.Tokens = append(foreign.R.Tokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Tokens = append(foreign.R.Tokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the token to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Tokens.
func (o *Token) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"user_id"}),
		strmangle.WhereClause("`", "`", 0, tokenPrimaryKeyColumns),
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
		o.R = &tokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Tokens: TokenSlice{o},
		}
	} else {
		related.R.Tokens = append(related.R.Tokens, o)
	}

	return nil
}

// Tokens retrieves all the records using an executor.
func Tokens(mods ...qm.QueryMod) tokenQuery {
	mods = append(mods, qm.From("`tokens`"))
	return tokenQuery{NewQuery(mods...)}
}

// FindToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindToken(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*Token, error) {
	tokenObj := &Token{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `tokens` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, tokenObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "mysql: unable to select from tokens")
	}

	return tokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Token) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no tokens provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(tokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tokenInsertCacheMut.RLock()
	cache, cached := tokenInsertCache[key]
	tokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tokenAllColumns,
			tokenColumnsWithDefault,
			tokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tokenType, tokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tokenType, tokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `tokens` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `tokens` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `tokens` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tokenPrimaryKeyColumns))
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
		return errors.Wrap(err, "mysql: unable to insert into tokens")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tokenMapping["id"] {
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
		return errors.Wrap(err, "mysql: unable to populate default values for tokens")
	}

CacheNoHooks:
	if !cached {
		tokenInsertCacheMut.Lock()
		tokenInsertCache[key] = cache
		tokenInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the Token.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Token) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	key := makeCacheKey(columns, nil)
	tokenUpdateCacheMut.RLock()
	cache, cached := tokenUpdateCache[key]
	tokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tokenAllColumns,
			tokenPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("mysql: unable to update tokens, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `tokens` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tokenType, tokenMapping, append(wl, tokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "mysql: unable to update tokens row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by update for tokens")
	}

	if !cached {
		tokenUpdateCacheMut.Lock()
		tokenUpdateCache[key] = cache
		tokenUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q tokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all for tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected for tokens")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `tokens` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to update all in token slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to retrieve rows affected all in update all token")
	}
	return rowsAff, nil
}

var mySQLTokenUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Token) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("mysql: no tokens provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(tokenColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTokenUniqueColumns, o)

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

	tokenUpsertCacheMut.RLock()
	cache, cached := tokenUpsertCache[key]
	tokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tokenAllColumns,
			tokenColumnsWithDefault,
			tokenColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tokenAllColumns,
			tokenPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("mysql: unable to upsert tokens, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "tokens", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `tokens` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tokenType, tokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tokenType, tokenMapping, ret)
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
		return errors.Wrap(err, "mysql: unable to upsert for tokens")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == tokenMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tokenType, tokenMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to retrieve unique values for tokens")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to populate default values for tokens")
	}

CacheNoHooks:
	if !cached {
		tokenUpsertCacheMut.Lock()
		tokenUpsertCache[key] = cache
		tokenUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single Token record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Token) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("mysql: no Token provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tokenPrimaryKeyMapping)
	sql := "DELETE FROM `tokens` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete from tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by delete for tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("mysql: no tokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from tokens")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for tokens")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "mysql: unable to delete all from token slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "mysql: failed to get rows affected by deleteall for tokens")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Token) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindToken(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `tokens`.* FROM `tokens` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "mysql: unable to reload all in TokenSlice")
	}

	*o = slice

	return nil
}

// TokenExists checks if the Token row exists.
func TokenExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `tokens` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "mysql: unable to check if tokens exists")
	}

	return exists, nil
}
