// Code generated by SQLBoiler 4.1.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Pilot is an object representing the database table.
type Pilot struct {
	ID        int       `boil:"id" json:"id" toml:"id" yaml:"id"`
	Name      string    `boil:"name" json:"name" toml:"name" yaml:"name"`
	Surname   string    `boil:"surname" json:"surname" toml:"surname" yaml:"surname"`
	Birthday  time.Time `boil:"birthday" json:"birthday" toml:"birthday" yaml:"birthday"`
	CountryID int       `boil:"country_id" json:"country_id" toml:"country_id" yaml:"country_id"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *pilotR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L pilotL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PilotColumns = struct {
	ID        string
	Name      string
	Surname   string
	Birthday  string
	CountryID string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	Name:      "name",
	Surname:   "surname",
	Birthday:  "birthday",
	CountryID: "country_id",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

// Generated where

var PilotWhere = struct {
	ID        whereHelperint
	Name      whereHelperstring
	Surname   whereHelperstring
	Birthday  whereHelpertime_Time
	CountryID whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperint{field: "`pilots`.`id`"},
	Name:      whereHelperstring{field: "`pilots`.`name`"},
	Surname:   whereHelperstring{field: "`pilots`.`surname`"},
	Birthday:  whereHelpertime_Time{field: "`pilots`.`birthday`"},
	CountryID: whereHelperint{field: "`pilots`.`country_id`"},
	CreatedAt: whereHelpertime_Time{field: "`pilots`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`pilots`.`updated_at`"},
}

// PilotRels is where relationship names are stored.
var PilotRels = struct {
	Flights            string
	SecondPilotFlights string
}{
	Flights:            "Flights",
	SecondPilotFlights: "SecondPilotFlights",
}

// pilotR is where relationships are stored.
type pilotR struct {
	Flights            FlightSlice `boil:"Flights" json:"Flights" toml:"Flights" yaml:"Flights"`
	SecondPilotFlights FlightSlice `boil:"SecondPilotFlights" json:"SecondPilotFlights" toml:"SecondPilotFlights" yaml:"SecondPilotFlights"`
}

// NewStruct creates a new relationship struct
func (*pilotR) NewStruct() *pilotR {
	return &pilotR{}
}

// pilotL is where Load methods for each relationship are stored.
type pilotL struct{}

var (
	pilotAllColumns            = []string{"id", "name", "surname", "birthday", "country_id", "created_at", "updated_at"}
	pilotColumnsWithoutDefault = []string{"name", "surname", "birthday", "country_id"}
	pilotColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	pilotPrimaryKeyColumns     = []string{"id"}
)

type (
	// PilotSlice is an alias for a slice of pointers to Pilot.
	// This should generally be used opposed to []Pilot.
	PilotSlice []*Pilot

	pilotQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	pilotType                 = reflect.TypeOf(&Pilot{})
	pilotMapping              = queries.MakeStructMapping(pilotType)
	pilotPrimaryKeyMapping, _ = queries.BindMapping(pilotType, pilotMapping, pilotPrimaryKeyColumns)
	pilotInsertCacheMut       sync.RWMutex
	pilotInsertCache          = make(map[string]insertCache)
	pilotUpdateCacheMut       sync.RWMutex
	pilotUpdateCache          = make(map[string]updateCache)
	pilotUpsertCacheMut       sync.RWMutex
	pilotUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single pilot record from the query using the global executor.
func (q pilotQuery) OneG() (*Pilot, error) {
	return q.One(boil.GetDB())
}

// One returns a single pilot record from the query.
func (q pilotQuery) One(exec boil.Executor) (*Pilot, error) {
	o := &Pilot{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(nil, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for pilots")
	}

	return o, nil
}

// AllG returns all Pilot records from the query using the global executor.
func (q pilotQuery) AllG() (PilotSlice, error) {
	return q.All(boil.GetDB())
}

// All returns all Pilot records from the query.
func (q pilotQuery) All(exec boil.Executor) (PilotSlice, error) {
	var o []*Pilot

	err := q.Bind(nil, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Pilot slice")
	}

	return o, nil
}

// CountG returns the count of all Pilot records in the query, and panics on error.
func (q pilotQuery) CountG() (int64, error) {
	return q.Count(boil.GetDB())
}

// Count returns the count of all Pilot records in the query.
func (q pilotQuery) Count(exec boil.Executor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count pilots rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q pilotQuery) ExistsG() (bool, error) {
	return q.Exists(boil.GetDB())
}

// Exists checks if the row exists in the table.
func (q pilotQuery) Exists(exec boil.Executor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow(exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if pilots exists")
	}

	return count > 0, nil
}

// Flights retrieves all the flight's Flights with an executor.
func (o *Pilot) Flights(mods ...qm.QueryMod) flightQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`flights`.`pilot_id`=?", o.ID),
	)

	query := Flights(queryMods...)
	queries.SetFrom(query.Query, "`flights`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`flights`.*"})
	}

	return query
}

// SecondPilotFlights retrieves all the flight's Flights with an executor via second_pilot_id column.
func (o *Pilot) SecondPilotFlights(mods ...qm.QueryMod) flightQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`flights`.`second_pilot_id`=?", o.ID),
	)

	query := Flights(queryMods...)
	queries.SetFrom(query.Query, "`flights`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`flights`.*"})
	}

	return query
}

// LoadFlights allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotL) LoadFlights(e boil.Executor, singular bool, maybePilot interface{}, mods queries.Applicator) error {
	var slice []*Pilot
	var object *Pilot

	if singular {
		object = maybePilot.(*Pilot)
	} else {
		slice = *maybePilot.(*[]*Pilot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pilotR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotR{}
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
		qm.From(`flights`),
		qm.WhereIn(`flights.pilot_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load flights")
	}

	var resultSlice []*Flight
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice flights")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on flights")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for flights")
	}

	if singular {
		object.R.Flights = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &flightR{}
			}
			foreign.R.Pilot = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.PilotID {
				local.R.Flights = append(local.R.Flights, foreign)
				if foreign.R == nil {
					foreign.R = &flightR{}
				}
				foreign.R.Pilot = local
				break
			}
		}
	}

	return nil
}

// LoadSecondPilotFlights allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (pilotL) LoadSecondPilotFlights(e boil.Executor, singular bool, maybePilot interface{}, mods queries.Applicator) error {
	var slice []*Pilot
	var object *Pilot

	if singular {
		object = maybePilot.(*Pilot)
	} else {
		slice = *maybePilot.(*[]*Pilot)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &pilotR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &pilotR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ID) {
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
		qm.From(`flights`),
		qm.WhereIn(`flights.second_pilot_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.Query(e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load flights")
	}

	var resultSlice []*Flight
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice flights")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on flights")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for flights")
	}

	if singular {
		object.R.SecondPilotFlights = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &flightR{}
			}
			foreign.R.SecondPilot = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.ID, foreign.SecondPilotID) {
				local.R.SecondPilotFlights = append(local.R.SecondPilotFlights, foreign)
				if foreign.R == nil {
					foreign.R = &flightR{}
				}
				foreign.R.SecondPilot = local
				break
			}
		}
	}

	return nil
}

// AddFlightsG adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.Flights.
// Sets related.R.Pilot appropriately.
// Uses the global database handle.
func (o *Pilot) AddFlightsG(insert bool, related ...*Flight) error {
	return o.AddFlights(boil.GetDB(), insert, related...)
}

// AddFlights adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.Flights.
// Sets related.R.Pilot appropriately.
func (o *Pilot) AddFlights(exec boil.Executor, insert bool, related ...*Flight) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.PilotID = o.ID
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `flights` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"pilot_id"}),
				strmangle.WhereClause("`", "`", 0, flightPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.PilotID = o.ID
		}
	}

	if o.R == nil {
		o.R = &pilotR{
			Flights: related,
		}
	} else {
		o.R.Flights = append(o.R.Flights, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &flightR{
				Pilot: o,
			}
		} else {
			rel.R.Pilot = o
		}
	}
	return nil
}

// AddSecondPilotFlightsG adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.SecondPilotFlights.
// Sets related.R.SecondPilot appropriately.
// Uses the global database handle.
func (o *Pilot) AddSecondPilotFlightsG(insert bool, related ...*Flight) error {
	return o.AddSecondPilotFlights(boil.GetDB(), insert, related...)
}

// AddSecondPilotFlights adds the given related objects to the existing relationships
// of the pilot, optionally inserting them as new records.
// Appends related to o.R.SecondPilotFlights.
// Sets related.R.SecondPilot appropriately.
func (o *Pilot) AddSecondPilotFlights(exec boil.Executor, insert bool, related ...*Flight) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.SecondPilotID, o.ID)
			if err = rel.Insert(exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `flights` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"second_pilot_id"}),
				strmangle.WhereClause("`", "`", 0, flightPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}
			if _, err = exec.Exec(updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.SecondPilotID, o.ID)
		}
	}

	if o.R == nil {
		o.R = &pilotR{
			SecondPilotFlights: related,
		}
	} else {
		o.R.SecondPilotFlights = append(o.R.SecondPilotFlights, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &flightR{
				SecondPilot: o,
			}
		} else {
			rel.R.SecondPilot = o
		}
	}
	return nil
}

// SetSecondPilotFlightsG removes all previously related items of the
// pilot replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.SecondPilot's SecondPilotFlights accordingly.
// Replaces o.R.SecondPilotFlights with related.
// Sets related.R.SecondPilot's SecondPilotFlights accordingly.
// Uses the global database handle.
func (o *Pilot) SetSecondPilotFlightsG(insert bool, related ...*Flight) error {
	return o.SetSecondPilotFlights(boil.GetDB(), insert, related...)
}

// SetSecondPilotFlights removes all previously related items of the
// pilot replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.SecondPilot's SecondPilotFlights accordingly.
// Replaces o.R.SecondPilotFlights with related.
// Sets related.R.SecondPilot's SecondPilotFlights accordingly.
func (o *Pilot) SetSecondPilotFlights(exec boil.Executor, insert bool, related ...*Flight) error {
	query := "update `flights` set `second_pilot_id` = null where `second_pilot_id` = ?"
	values := []interface{}{o.ID}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}
	_, err := exec.Exec(query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.SecondPilotFlights {
			queries.SetScanner(&rel.SecondPilotID, nil)
			if rel.R == nil {
				continue
			}

			rel.R.SecondPilot = nil
		}

		o.R.SecondPilotFlights = nil
	}
	return o.AddSecondPilotFlights(exec, insert, related...)
}

// RemoveSecondPilotFlightsG relationships from objects passed in.
// Removes related items from R.SecondPilotFlights (uses pointer comparison, removal does not keep order)
// Sets related.R.SecondPilot.
// Uses the global database handle.
func (o *Pilot) RemoveSecondPilotFlightsG(related ...*Flight) error {
	return o.RemoveSecondPilotFlights(boil.GetDB(), related...)
}

// RemoveSecondPilotFlights relationships from objects passed in.
// Removes related items from R.SecondPilotFlights (uses pointer comparison, removal does not keep order)
// Sets related.R.SecondPilot.
func (o *Pilot) RemoveSecondPilotFlights(exec boil.Executor, related ...*Flight) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.SecondPilotID, nil)
		if rel.R != nil {
			rel.R.SecondPilot = nil
		}
		if _, err = rel.Update(exec, boil.Whitelist("second_pilot_id")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.SecondPilotFlights {
			if rel != ri {
				continue
			}

			ln := len(o.R.SecondPilotFlights)
			if ln > 1 && i < ln-1 {
				o.R.SecondPilotFlights[i] = o.R.SecondPilotFlights[ln-1]
			}
			o.R.SecondPilotFlights = o.R.SecondPilotFlights[:ln-1]
			break
		}
	}

	return nil
}

// Pilots retrieves all the records using an executor.
func Pilots(mods ...qm.QueryMod) pilotQuery {
	mods = append(mods, qm.From("`pilots`"))
	return pilotQuery{NewQuery(mods...)}
}

// FindPilotG retrieves a single record by ID.
func FindPilotG(iD int, selectCols ...string) (*Pilot, error) {
	return FindPilot(boil.GetDB(), iD, selectCols...)
}

// FindPilot retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPilot(exec boil.Executor, iD int, selectCols ...string) (*Pilot, error) {
	pilotObj := &Pilot{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `pilots` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(nil, exec, pilotObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from pilots")
	}

	return pilotObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Pilot) InsertG(columns boil.Columns) error {
	return o.Insert(boil.GetDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Pilot) Insert(exec boil.Executor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pilots provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	if o.UpdatedAt.IsZero() {
		o.UpdatedAt = currTime
	}

	nzDefaults := queries.NonZeroDefaultSet(pilotColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	pilotInsertCacheMut.RLock()
	cache, cached := pilotInsertCache[key]
	pilotInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			pilotAllColumns,
			pilotColumnsWithDefault,
			pilotColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(pilotType, pilotMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `pilots` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `pilots` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `pilots` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, pilotPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into pilots")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == pilotMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}
	err = exec.QueryRow(cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for pilots")
	}

CacheNoHooks:
	if !cached {
		pilotInsertCacheMut.Lock()
		pilotInsertCache[key] = cache
		pilotInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Pilot record using the global executor.
// See Update for more documentation.
func (o *Pilot) UpdateG(columns boil.Columns) (int64, error) {
	return o.Update(boil.GetDB(), columns)
}

// Update uses an executor to update the Pilot.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Pilot) Update(exec boil.Executor, columns boil.Columns) (int64, error) {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt = currTime

	var err error
	key := makeCacheKey(columns, nil)
	pilotUpdateCacheMut.RLock()
	cache, cached := pilotUpdateCache[key]
	pilotUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			pilotAllColumns,
			pilotPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update pilots, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `pilots` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, pilotPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, append(wl, pilotPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update pilots row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for pilots")
	}

	if !cached {
		pilotUpdateCacheMut.Lock()
		pilotUpdateCache[key] = cache
		pilotUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q pilotQuery) UpdateAllG(cols M) (int64, error) {
	return q.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q pilotQuery) UpdateAll(exec boil.Executor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for pilots")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o PilotSlice) UpdateAllG(cols M) (int64, error) {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PilotSlice) UpdateAll(exec boil.Executor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `pilots` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pilotPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in pilot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all pilot")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Pilot) UpsertG(updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(boil.GetDB(), updateColumns, insertColumns)
}

var mySQLPilotUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Pilot) Upsert(exec boil.Executor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no pilots provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.IsZero() {
		o.CreatedAt = currTime
	}
	o.UpdatedAt = currTime

	nzDefaults := queries.NonZeroDefaultSet(pilotColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPilotUniqueColumns, o)

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

	pilotUpsertCacheMut.RLock()
	cache, cached := pilotUpsertCache[key]
	pilotUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			pilotAllColumns,
			pilotColumnsWithDefault,
			pilotColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			pilotAllColumns,
			pilotPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert pilots, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "pilots", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `pilots` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(pilotType, pilotMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(pilotType, pilotMapping, ret)
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
	result, err := exec.Exec(cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for pilots")
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

	o.ID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == pilotMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(pilotType, pilotMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for pilots")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}
	err = exec.QueryRow(cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for pilots")
	}

CacheNoHooks:
	if !cached {
		pilotUpsertCacheMut.Lock()
		pilotUpsertCache[key] = cache
		pilotUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Pilot record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Pilot) DeleteG() (int64, error) {
	return o.Delete(boil.GetDB())
}

// Delete deletes a single Pilot record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Pilot) Delete(exec boil.Executor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Pilot provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), pilotPrimaryKeyMapping)
	sql := "DELETE FROM `pilots` WHERE `id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for pilots")
	}

	return rowsAff, nil
}

func (q pilotQuery) DeleteAllG() (int64, error) {
	return q.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all matching rows.
func (q pilotQuery) DeleteAll(exec boil.Executor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no pilotQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.Exec(exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pilots")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pilots")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o PilotSlice) DeleteAllG() (int64, error) {
	return o.DeleteAll(boil.GetDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PilotSlice) DeleteAll(exec boil.Executor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `pilots` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pilotPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}
	result, err := exec.Exec(sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from pilot slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for pilots")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Pilot) ReloadG() error {
	if o == nil {
		return errors.New("models: no Pilot provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Pilot) Reload(exec boil.Executor) error {
	ret, err := FindPilot(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PilotSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty PilotSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PilotSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PilotSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), pilotPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `pilots`.* FROM `pilots` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, pilotPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(nil, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PilotSlice")
	}

	*o = slice

	return nil
}

// PilotExistsG checks if the Pilot row exists.
func PilotExistsG(iD int) (bool, error) {
	return PilotExists(boil.GetDB(), iD)
}

// PilotExists checks if the Pilot row exists.
func PilotExists(exec boil.Executor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `pilots` where `id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, iD)
	}
	row := exec.QueryRow(sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if pilots exists")
	}

	return exists, nil
}
