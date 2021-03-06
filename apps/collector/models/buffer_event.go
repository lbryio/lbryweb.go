// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// BufferEvent is an object representing the database table.
type BufferEvent struct {
	ID             int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	Time           time.Time   `boil:"time" json:"time" toml:"time" yaml:"time"`
	URL            string      `boil:"url" json:"url" toml:"url" yaml:"url"`
	Client         string      `boil:"client" json:"client" toml:"client" yaml:"client"`
	Device         string      `boil:"device" json:"device" toml:"device" yaml:"device"`
	Position       int         `boil:"position" json:"position" toml:"position" yaml:"position"`
	Duration       null.Int    `boil:"duration" json:"duration,omitempty" toml:"duration" yaml:"duration,omitempty"`
	StreamDuration null.Int    `boil:"stream_duration" json:"stream_duration,omitempty" toml:"stream_duration" yaml:"stream_duration,omitempty"`
	StreamBitrate  null.Int    `boil:"stream_bitrate" json:"stream_bitrate,omitempty" toml:"stream_bitrate" yaml:"stream_bitrate,omitempty"`
	Player         null.String `boil:"player" json:"player,omitempty" toml:"player" yaml:"player,omitempty"`
	ReadyState     int16       `boil:"ready_state" json:"ready_state" toml:"ready_state" yaml:"ready_state"`

	R *bufferEventR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L bufferEventL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var BufferEventColumns = struct {
	ID             string
	Time           string
	URL            string
	Client         string
	Device         string
	Position       string
	Duration       string
	StreamDuration string
	StreamBitrate  string
	Player         string
	ReadyState     string
}{
	ID:             "id",
	Time:           "time",
	URL:            "url",
	Client:         "client",
	Device:         "device",
	Position:       "position",
	Duration:       "duration",
	StreamDuration: "stream_duration",
	StreamBitrate:  "stream_bitrate",
	Player:         "player",
	ReadyState:     "ready_state",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

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

type whereHelperint16 struct{ field string }

func (w whereHelperint16) EQ(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint16) NEQ(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint16) LT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint16) LTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint16) GT(x int16) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint16) GTE(x int16) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }

var BufferEventWhere = struct {
	ID             whereHelperint
	Time           whereHelpertime_Time
	URL            whereHelperstring
	Client         whereHelperstring
	Device         whereHelperstring
	Position       whereHelperint
	Duration       whereHelpernull_Int
	StreamDuration whereHelpernull_Int
	StreamBitrate  whereHelpernull_Int
	Player         whereHelpernull_String
	ReadyState     whereHelperint16
}{
	ID:             whereHelperint{field: "\"buffer_event\".\"id\""},
	Time:           whereHelpertime_Time{field: "\"buffer_event\".\"time\""},
	URL:            whereHelperstring{field: "\"buffer_event\".\"url\""},
	Client:         whereHelperstring{field: "\"buffer_event\".\"client\""},
	Device:         whereHelperstring{field: "\"buffer_event\".\"device\""},
	Position:       whereHelperint{field: "\"buffer_event\".\"position\""},
	Duration:       whereHelpernull_Int{field: "\"buffer_event\".\"duration\""},
	StreamDuration: whereHelpernull_Int{field: "\"buffer_event\".\"stream_duration\""},
	StreamBitrate:  whereHelpernull_Int{field: "\"buffer_event\".\"stream_bitrate\""},
	Player:         whereHelpernull_String{field: "\"buffer_event\".\"player\""},
	ReadyState:     whereHelperint16{field: "\"buffer_event\".\"ready_state\""},
}

// BufferEventRels is where relationship names are stored.
var BufferEventRels = struct {
}{}

// bufferEventR is where relationships are stored.
type bufferEventR struct {
}

// NewStruct creates a new relationship struct
func (*bufferEventR) NewStruct() *bufferEventR {
	return &bufferEventR{}
}

// bufferEventL is where Load methods for each relationship are stored.
type bufferEventL struct{}

var (
	bufferEventAllColumns            = []string{"id", "time", "url", "client", "device", "position", "duration", "stream_duration", "stream_bitrate", "player", "ready_state"}
	bufferEventColumnsWithoutDefault = []string{"url", "client", "device", "position", "duration", "stream_duration", "stream_bitrate", "player", "ready_state"}
	bufferEventColumnsWithDefault    = []string{"id", "time"}
	bufferEventPrimaryKeyColumns     = []string{"id", "time"}
)

type (
	// BufferEventSlice is an alias for a slice of pointers to BufferEvent.
	// This should generally be used opposed to []BufferEvent.
	BufferEventSlice []*BufferEvent
	// BufferEventHook is the signature for custom BufferEvent hook methods
	BufferEventHook func(boil.Executor, *BufferEvent) error

	bufferEventQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	bufferEventType                 = reflect.TypeOf(&BufferEvent{})
	bufferEventMapping              = queries.MakeStructMapping(bufferEventType)
	bufferEventPrimaryKeyMapping, _ = queries.BindMapping(bufferEventType, bufferEventMapping, bufferEventPrimaryKeyColumns)
	bufferEventInsertCacheMut       sync.RWMutex
	bufferEventInsertCache          = make(map[string]insertCache)
	bufferEventUpdateCacheMut       sync.RWMutex
	bufferEventUpdateCache          = make(map[string]updateCache)
	bufferEventUpsertCacheMut       sync.RWMutex
	bufferEventUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var bufferEventBeforeInsertHooks []BufferEventHook
var bufferEventBeforeUpdateHooks []BufferEventHook
var bufferEventBeforeDeleteHooks []BufferEventHook
var bufferEventBeforeUpsertHooks []BufferEventHook

var bufferEventAfterInsertHooks []BufferEventHook
var bufferEventAfterSelectHooks []BufferEventHook
var bufferEventAfterUpdateHooks []BufferEventHook
var bufferEventAfterDeleteHooks []BufferEventHook
var bufferEventAfterUpsertHooks []BufferEventHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *BufferEvent) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *BufferEvent) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *BufferEvent) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *BufferEvent) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *BufferEvent) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *BufferEvent) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *BufferEvent) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *BufferEvent) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *BufferEvent) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range bufferEventAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddBufferEventHook registers your hook function for all future operations.
func AddBufferEventHook(hookPoint boil.HookPoint, bufferEventHook BufferEventHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		bufferEventBeforeInsertHooks = append(bufferEventBeforeInsertHooks, bufferEventHook)
	case boil.BeforeUpdateHook:
		bufferEventBeforeUpdateHooks = append(bufferEventBeforeUpdateHooks, bufferEventHook)
	case boil.BeforeDeleteHook:
		bufferEventBeforeDeleteHooks = append(bufferEventBeforeDeleteHooks, bufferEventHook)
	case boil.BeforeUpsertHook:
		bufferEventBeforeUpsertHooks = append(bufferEventBeforeUpsertHooks, bufferEventHook)
	case boil.AfterInsertHook:
		bufferEventAfterInsertHooks = append(bufferEventAfterInsertHooks, bufferEventHook)
	case boil.AfterSelectHook:
		bufferEventAfterSelectHooks = append(bufferEventAfterSelectHooks, bufferEventHook)
	case boil.AfterUpdateHook:
		bufferEventAfterUpdateHooks = append(bufferEventAfterUpdateHooks, bufferEventHook)
	case boil.AfterDeleteHook:
		bufferEventAfterDeleteHooks = append(bufferEventAfterDeleteHooks, bufferEventHook)
	case boil.AfterUpsertHook:
		bufferEventAfterUpsertHooks = append(bufferEventAfterUpsertHooks, bufferEventHook)
	}
}

// OneG returns a single bufferEvent record from the query using the global executor.
func (q bufferEventQuery) OneG() (*BufferEvent, error) {
	return q.One(boil.GetDB())
}

// One returns a single bufferEvent record from the query.
func (q bufferEventQuery) One(exec boil.Executor) (*BufferEvent, error) {
	o := &BufferEvent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for buffer_event")
	}

	if err := o.doAfterSelectHooks(exec); err != nil {
		return o, err
	}

	return o, nil
}

// AllG returns all BufferEvent records from the query using the global executor.
func (q bufferEventQuery) AllG() (BufferEventSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all BufferEvent records from the query.
func (q bufferEventQuery) All(exec boil.Executor) (BufferEventSlice, error) {
	var o []*BufferEvent

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to BufferEvent slice")
	}

	if len(bufferEventAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountG returns the count of all BufferEvent records in the query, and panics on error.
func (q bufferEventQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all BufferEvent records in the query.
func (q bufferEventQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count buffer_event rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q bufferEventQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q bufferEventQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if buffer_event exists")
	}

	return count > 0, nil
}

// BufferEvents retrieves all the records using an executor.
func BufferEvents(mods ...qm.QueryMod) bufferEventQuery {
	mods = append(mods, qm.From("\"buffer_event\""))
	return bufferEventQuery{NewQuery(mods...)}
}

// FindBufferEventG retrieves a single record by ID.
func FindBufferEventG(iD int, time time.Time, selectCols ...string) (*BufferEvent, error) {
	return FindBufferEvent(boil.GetDB(), iD, time, selectCols...)
}

// FindBufferEvent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindBufferEvent(exec boil.Executor, iD int, time time.Time, selectCols ...string) (*BufferEvent, error) {
	bufferEventObj := &BufferEvent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"buffer_event\" where \"id\"=$1 AND \"time\"=$2", sel,
	)

	q := queries.Raw(query, iD, time)

	err := q.Bind(nil, exec, bufferEventObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from buffer_event")
	}

	return bufferEventObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *BufferEvent) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *BufferEvent) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buffer_event provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bufferEventColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	bufferEventInsertCacheMut.RLock()
	cache, cached := bufferEventInsertCache[key]
	bufferEventInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			bufferEventAllColumns,
			bufferEventColumnsWithDefault,
			bufferEventColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(bufferEventType, bufferEventMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(bufferEventType, bufferEventMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"buffer_event\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"buffer_event\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into buffer_event")
	}

	if !cached {
		bufferEventInsertCacheMut.Lock()
		bufferEventInsertCache[key] = cache
		bufferEventInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single BufferEvent record using the global executor.
// See Update for more documentation.
func (o *BufferEvent) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the BufferEvent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *BufferEvent) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	bufferEventUpdateCacheMut.RLock()
	cache, cached := bufferEventUpdateCache[key]
	bufferEventUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			bufferEventAllColumns,
			bufferEventPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update buffer_event, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"buffer_event\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, bufferEventPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(bufferEventType, bufferEventMapping, append(wl, bufferEventPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.Exec(cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update buffer_event row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for buffer_event")
	}

	if !cached {
		bufferEventUpdateCacheMut.Lock()
		bufferEventUpdateCache[key] = cache
		bufferEventUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(exec)
}

// UpdateAllG updates all rows with the specified column values.
func (q bufferEventQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q bufferEventQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for buffer_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for buffer_event")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o BufferEventSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o BufferEventSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bufferEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"buffer_event\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, bufferEventPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in bufferEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all bufferEvent")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *BufferEvent) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *BufferEvent) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no buffer_event provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(bufferEventColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	bufferEventUpsertCacheMut.RLock()
	cache, cached := bufferEventUpsertCache[key]
	bufferEventUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			bufferEventAllColumns,
			bufferEventColumnsWithDefault,
			bufferEventColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			bufferEventAllColumns,
			bufferEventPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert buffer_event, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(bufferEventPrimaryKeyColumns))
			copy(conflict, bufferEventPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"buffer_event\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(bufferEventType, bufferEventMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(bufferEventType, bufferEventMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert buffer_event")
	}

	if !cached {
		bufferEventUpsertCacheMut.Lock()
		bufferEventUpsertCache[key] = cache
		bufferEventUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteG deletes a single BufferEvent record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *BufferEvent) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single BufferEvent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *BufferEvent) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no BufferEvent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), bufferEventPrimaryKeyMapping)
	sql := "DELETE FROM \"buffer_event\" WHERE \"id\"=$1 AND \"time\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from buffer_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for buffer_event")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q bufferEventQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no bufferEventQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from buffer_event")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buffer_event")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o BufferEventSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o BufferEventSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(bufferEventBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bufferEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"buffer_event\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bufferEventPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from bufferEvent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for buffer_event")
	}

	if len(bufferEventAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *BufferEvent) ReloadG() error {
	if o == nil {
		return errors.New("models: no BufferEvent provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *BufferEvent) Reload(exec boil.Executor) error {
	ret, err := FindBufferEvent(exec, o.ID, o.Time)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BufferEventSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty BufferEventSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *BufferEventSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := BufferEventSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), bufferEventPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"buffer_event\".* FROM \"buffer_event\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, bufferEventPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in BufferEventSlice")
	}

	*o = slice

	return nil
}

// BufferEventExistsG checks if the BufferEvent row exists.
func BufferEventExistsG(iD int, time time.Time) (bool, error) {
	return BufferEventExists(boil.GetDB(), iD, time)
}

// BufferEventExists checks if the BufferEvent row exists.
func BufferEventExists(exec boil.Executor, iD int, time time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"buffer_event\" where \"id\"=$1 AND \"time\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD, time)
	}

	row := exec.QueryRow(sql, iD, time)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if buffer_event exists")
	}

	return exists, nil
}
