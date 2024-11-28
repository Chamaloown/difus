package models

import "time"

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