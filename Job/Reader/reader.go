package reader

import (
	"database/sql"
	"fmt"
	"log"

	models "github.com/chamaloown/difus/Models"
)

func GetJobByName(db *sql.DB, name string) (models.Job, error) {
	var j models.Job
	fmt.Println(name)
	err := db.QueryRow("SELECT id, name, type FROM almanax.jobs WHERE UPPER(name) LIKE UPPER($1)", "%"+name+"%").Scan(&j.Id, &j.Name, &j.Type)

	if err != nil {
		log.Fatal(err)
	}
	return j, nil
}

func GetJobs(db *sql.DB) ([]models.Job, error) {
	rows, err := db.Query("SELECT id, name, type FROM almanax.jobs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobs []models.Job
	for rows.Next() {
		var r models.Job
		if err := rows.Scan(&r.Id, &r.Name, &r.Type); err != nil {
			return nil, err
		}
		jobs = append(jobs, r)
	}
	return jobs, nil
}

func GetJobById(db *sql.DB, id int) (models.Job, error) {
	var j models.Job
	err := db.QueryRow("SELECT id, name, type FROM almanax.Jobs WHERE id = $1", id).Scan(&j.Id, &j.Name, &j.Type)

	if err != nil {
		log.Fatal(err)
	}
	return j, nil
}
