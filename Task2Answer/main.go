package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ceeyahya/spotlas-backend-test/database"
	"github.com/ceeyahya/spotlas-backend-test/models"
)

func spotsInArea(db database.DBInstance, lat float64, long float64, radius int, shape string) {
	spots := []models.Spot{}
	cRadius := float64(radius) / 100000

	if shape == "circle" {
		db.Db.Table("MY_TABLE").Model(&models.Spot{}).Where(fmt.Sprintf("ST_Transform('SRID=21781;POINT(%f %f)', 21781) && ST_expand(ST_SetSRID(ST_GeomFromText(CONCAT('SRID=21781;', ST_AsText(coordinates))),21781),%f)", lat, long, cRadius)).Order("rating DESC").Find(&spots)
	} else {
		db.Db.Table("MY_TABLE").Model(&models.Spot{}).Where(fmt.Sprintf("ST_Transform(ST_GeomFromText('SRID=21781;POINT(%f %f)'),21781) && ST_MakeEnvelope(ST_X(ST_GeomFromEWKT(ST_AsText(coordinates))) - %f, ST_Y(ST_GeomFromEWKT(ST_AsText(coordinates))) - %f, ST_X(ST_GeomFromEWKT(ST_AsText(coordinates))) + %f, ST_Y(ST_GeomFromEWKT(ST_AsText(coordinates))) + %f, 21781)", lat, long, cRadius, cRadius, cRadius, cRadius)).Order("rating DESC").Find(&spots)
	}

	sJSON, err := json.MarshalIndent(spots, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Println(string(sJSON))
}

func main() {
	database.Connect()
	db := database.Instance.Db

	// Create the `postgis` extension if it doesn't exist
	db.Exec("CREATE EXTENSION IF NOT EXISTS postgis")

	// An Endpoint for a circle shaped area
	spotsInArea(database.Instance, 13.4068255, 52.5247749, 900, "circle")

	// An Endpoint for a square shaped area
	spotsInArea(database.Instance, 13.4068255, 52.5247749, 900, "square")
}
