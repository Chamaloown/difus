package reader

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	models "github.com/chamaloown/difus/Models"
)

const daysInAYear = 365

func weekStart(year, week int) time.Time {
    t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

    if wd := t.Weekday(); wd == time.Sunday {
        t = t.AddDate(0, 0, -6)
    } else {
        t = t.AddDate(0, 0, -int(wd)+1)
    }

    _, w := t.ISOWeek()
    t = t.AddDate(0, 0, (week-w)*7)

    return t
}

func weekRange(year, week int) (start, end time.Time) {
    start = weekStart(year, week)
    end = start.AddDate(0, 0, 6)
    return
}


func IsAlmanaxComplet(db *sql.DB) bool {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM almanax.almanaxes`).Scan(&count)
	if err != nil {
		log.Printf("Error querying almanax.almanaxes: %v\n", err)
		count = 0
	}
	return count == daysInAYear + 54 // + 54 because we start from 08/11/2024
}

func ReadAlmanaxes(db *sql.DB) ([]models.Almanax, error) {
	rows, err := db.Query("SELECT id, date, merydes, type, bonus, offerings, quantity_offered, kamas FROM almanax.almanaxes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var almanaxes []models.Almanax
	for rows.Next() {
		var r models.Almanax
		if err := rows.Scan(&r.Id, &r.Date, &r.Merydes, &r.Type, &r.Bonus, &r.Offerings, &r.QuantityOffered, &r.Kamas); err != nil {
			return nil, err
		}
		almanaxes = append(almanaxes, r)
	}
	return almanaxes, nil
}


func GetAlmanax(db *sql.DB, date time.Time) (models.Almanax, error) {
	fmt.Println("Get Almanax")
	var a models.Almanax
	err := db.QueryRow("SELECT id, date, merydes, type, bonus, offerings, quantity_offered, kamas FROM almanax.almanaxes WHERE date = $1", date).Scan(&a.Id, &a.Date, &a.Merydes, &a.Type, &a.Bonus, &a.Offerings, &a.QuantityOffered, &a.Kamas)

	if err != nil {
		log.Fatal(err)
	}
	return a, nil
}

func GetAlmanaxesInRange(db *sql.DB, start time.Time, end time.Time) ([]models.Almanax, error) {
	rows, err := db.Query("SELECT id, date, merydes, type, bonus, offerings, quantity_offered, kamas FROM almanax.almanaxes WHERE date >= $1 AND date <= $2", start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var almanaxes []models.Almanax
	for rows.Next() {
		var r models.Almanax
		if err := rows.Scan(&r.Id, &r.Date, &r.Merydes, &r.Type, &r.Bonus, &r.Offerings, &r.QuantityOffered, &r.Kamas); err != nil {
			return nil, err
		}
		almanaxes = append(almanaxes, r)
	}
	return almanaxes, nil
}

func GetWeeklyAlmanax(db *sql.DB) ([]models.Almanax, error) {
	year, month := time.Now().ISOWeek()
	start, end := weekRange(year, month)
	alamanax, err := GetAlmanaxesInRange(db, start, end)
	if err != nil {
		log.Fatal(err)
	}
	return alamanax, nil
}