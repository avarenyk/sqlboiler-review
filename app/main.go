package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"sqlboiler-review/app/models"
	"time"
)


// used to hack Set() and Add() functions
type EmptyExec struct {
}

func (e EmptyExec) Exec(query string, args ...interface{}) (sql.Result, error) {
	return nil, nil
}
func (e EmptyExec) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return nil, nil
}
func (e EmptyExec) QueryRow(query string, args ...interface{}) *sql.Row {
	return nil
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp/test?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	boil.SetDB(db)
	boil.DebugMode = true

	if _, err := models.Countries().DeleteAllG(); err != nil {
		log.Fatal(err)
	}

	UA := &models.Country{
		Code: "UA",
		Name: "Ukraine",
	}

	GB := &models.Country{
		Code: "GB",
		Name: "Great Britain",
	}

	if err := UA.InsertG(boil.Infer()); err != nil {
		log.Fatal(err)
	}
	if err := GB.InsertG(boil.Infer()); err != nil {
		log.Fatal(err)
	}

	log.Printf("UA country id = %d\n", UA.ID)
	log.Printf("GB country id = %d\n", GB.ID)

	boryspil := &models.Airport{
		Code: "KBP",
	}

	if err := UA.AddAirportsG(true, boryspil); err != nil {
		log.Fatal(err)
	}

	heathrow := &models.Airport{
		Code: "LHR",
	}

	if err := GB.AddAirportsG(true, heathrow); err != nil {
		log.Fatal(err)
	}

	boeing := &models.Jet{
		Type:  "Boeing-737",
		Color: "Red",
	}

	if err := boryspil.AddHomeAirportJetsG(true, boeing); err != nil {
		log.Fatal(err)
	}

	pilotAndrey := &models.Pilot{
		Name:      "Andrey",
		Surname:   "Invanov",
		Birthday:  time.Date(1990, 1, 10, 0, 0, 0, 0, time.UTC),
		CountryID: UA.ID,
	}

	pilotKyle := &models.Pilot{
		Name:      "Kyle",
		Surname:   "Jonson",
		Birthday:  time.Date(1987, 12, 2, 0, 0, 0, 0, time.UTC),
		CountryID: GB.ID,
	}

	if err := pilotAndrey.InsertG(boil.Infer()); err != nil {
		log.Fatal(err)
	}
	if err := pilotKyle.InsertG(boil.Infer()); err != nil {
		log.Fatal(err)
	}

	kbp_lhr := &models.Flight{}
	// with default executor it would perform an update query on each of this methods
	// the second parameter is about what to do:
	// true - insert related entity
	// false - update current entity
	// so we want to set false, coz it builds less complex query, and do not execute update query immedieate
	kbp_lhr.SetOrigin(EmptyExec{}, false, boryspil)
	kbp_lhr.SetDestination(EmptyExec{}, false, heathrow)
	kbp_lhr.SetJet(EmptyExec{}, false, boeing)
	kbp_lhr.SetPilot(EmptyExec{}, false, pilotAndrey)
	kbp_lhr.SetSecondPilot(EmptyExec{}, false, pilotKyle)

	lhr_kpb := &models.Flight{}
	lhr_kpb.SetOrigin(EmptyExec{}, false, heathrow)
	lhr_kpb.SetDestination(EmptyExec{}, false, boryspil)
	lhr_kpb.SetJet(EmptyExec{}, false, boeing)
	lhr_kpb.SetPilot(EmptyExec{}, false, pilotAndrey)

	if err := kbp_lhr.InsertG(boil.Infer()); err != nil {
		log.Fatal(err)
	}
	if err := lhr_kpb.InsertG(boil.Infer()); err != nil {
		log.Fatal(err)
	}

	_, err = models.Countries(
		models.CountryWhere.Code.EQ("UA"),
		Load(Rels(models.CountryRels.Airports, models.AirportRels.OriginFlights)),
		Load(Rels(models.CountryRels.Airports, models.AirportRels.DestinationFlights)),
	).OneG()

	if err != nil {
		log.Fatal(err)
	}

	flights, err := models.Flights(
		InnerJoin("airports a on a.id = flights.origin_id or a.id = flights.destination_id"),
		InnerJoin("countries c on c.id = a.country_id"),
		Where("c.code = 'UA'"),
		Load(Rels(models.FlightRels.Origin), Where("airports.code = ?", "test")),
		Load(Rels(models.FlightRels.Origin, models.AirportRels.Country)),
		Load(Rels(models.FlightRels.Destination, models.AirportRels.Country)),
		).AllG()

	for _, f := range flights {
		fmt.Printf("flight: %+v\n", f.R.Destination)
	}

	pilotAndrey.Surname = "Other Surname"
	pilotAndrey.UpdateG(boil.Whitelist("surname"))
}
