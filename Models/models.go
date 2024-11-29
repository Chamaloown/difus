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

type Job struct {
	Id   int
	Name string
	Type string
}

type Class struct {
	Id   int
	Name string
}

type User struct {
	Id       int
	Name     string
	Username string
	Class    Class
	Jobs     []Job
}
