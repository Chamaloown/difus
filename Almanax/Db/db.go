package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	parser "github.com/chamaloown/difus/Almanax/Parser"
	_ "github.com/lib/pq"
)

var (
	dbInstance *sql.DB
	dbOnce     sync.Once
)

const daysInAYear = 365

type Almanax struct {
	Id              int
	Date            time.Time
	Merydes         string
	Type            string
	Bonus           string
	Offerings       string
	QuantityOffered int
	Kamas           int
}

func GetDBInstance() *sql.DB {
	dbOnce.Do(func() {
		fmt.Println("Initializing database connection...")

		connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=%s",
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_HOST"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_SSL"))

		var err error
		dbInstance, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Fatal("Error opening database connection:", err)
		}

		dbInstance.SetMaxOpenConns(25)
		dbInstance.SetMaxIdleConns(25)
		dbInstance.SetConnMaxLifetime(5 * time.Minute)

		if err = dbInstance.Ping(); err != nil {
			log.Fatal("Error connecting to the database:", err)
		}
	})
	return dbInstance
}

func isAlreadyCharged(db *sql.DB) bool {
	var count int
	err := db.QueryRow(`SELECT COUNT(*) FROM almanax.almanaxes`).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count == daysInAYear + 54 // + 54 because we start from 08/11/2024
}

func Setup() {
	db := GetDBInstance()

	if isAlreadyCharged(db) {
		fmt.Println("Database is already set to use!")
		return
	}

	_, err := db.Exec(`
        CREATE SCHEMA IF NOT EXISTS almanax;

        CREATE TABLE IF NOT EXISTS almanax.almanaxes (
            id serial primary key not null,
            date date not null,
            merydes varchar not null,
            type varchar not null,
            bonus varchar not null,
            offerings varchar not null,
            quantity_offered int not null,
            kamas int not null
        )
    `)

	if err != nil {
		log.Fatal(err)
	}

	records, err := parser.Run()
	if err != nil {
		log.Fatal(err)
	}

	for _, val := range records[1:] {
		date, err := time.Parse("02/01/2006", val[0])
		if err != nil {
			log.Fatal(err)
		}

		qty, err := strconv.Atoi(val[5])
		if err != nil {
			log.Fatal(err)
		}

		kamas, err := strconv.Atoi(val[6])
		if err != nil {
			log.Fatal(err)
		}

		newEntry := Almanax{
			Date:            date,
			Merydes:         val[1],
			Type:            val[2],
			Bonus:           val[3],
			Offerings:       val[4],
			QuantityOffered: qty,
			Kamas:           kamas,
		}

		_, err = createAlmanax(db, newEntry)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Successfully charged the database!")
}

func createAlmanax(db *sql.DB, almanax Almanax) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO almanax.almanaxes(date, merydes, type, bonus, offerings, quantity_offered, kamas) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id",
		almanax.Date, almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func ReadAlmanaxes(db *sql.DB) ([]Almanax, error) {
	rows, err := db.Query("SELECT id, date, merydes, type, bonus, offerings, quantity_offered, kamas FROM almanax.almanaxes")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var almanaxes []Almanax
	for rows.Next() {
		var r Almanax
		if err := rows.Scan(&r.Id, &r.Date, &r.Merydes, &r.Type, &r.Bonus, &r.Offerings, &r.QuantityOffered, &r.Kamas); err != nil {
			return nil, err
		}
		almanaxes = append(almanaxes, r)
	}
	return almanaxes, nil
}


func GetAlmanax(db *sql.DB, date time.Time) (Almanax, error) {
	var a Almanax
	err := db.QueryRow("SELECT id, date, merydes, type, bonus, offerings, quantity_offered, kamas FROM almanax.almanaxes WHERE date = $1", date).Scan(&a.Id, &a.Date, &a.Merydes, &a.Type, &a.Bonus, &a.Offerings, &a.QuantityOffered, &a.Kamas)

	if err != nil {
		log.Fatal(err)
	}
	return a, nil
}

func GetAlmanaxesInRange(db *sql.DB, start time.Time, end time.Time) ([]Almanax, error) {
	rows, err := db.Query("SELECT id, date, merydes, type, bonus, offerings, quantity_offered, kamas FROM almanax.almanaxes WHERE date >= $1 AND date <= $2", start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var almanaxes []Almanax
	for rows.Next() {
		var r Almanax
		if err := rows.Scan(&r.Id, &r.Date, &r.Merydes, &r.Type, &r.Bonus, &r.Offerings, &r.QuantityOffered, &r.Kamas); err != nil {
			return nil, err
		}
		almanaxes = append(almanaxes, r)
	}
	return almanaxes, nil
}


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

func GetWeeklyAlmanax(db *sql.DB) ([]Almanax, error) {
	year, month := time.Now().ISOWeek()
	start, end := weekRange(year, month)
	alamanax, err := GetAlmanaxesInRange(db, start, end)
	if err != nil {
		log.Fatal(err)
	}
	return alamanax, nil
}