package parser

import (
	"encoding/csv"
	"log"
	"os"
)

func Run() (records [][]string, err error) {
	f, err := os.Open("./almanax.csv")

	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	data, err := r.ReadAll()

	if err != nil {
		log.Fatal(err)
	}

	return data, err
}
