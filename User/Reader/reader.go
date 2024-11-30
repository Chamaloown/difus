package reader

import (
	"database/sql"
	"log"

	models "github.com/chamaloown/difus/Models"
)

func GetUsersByJob(db *sql.DB, job string) ([]models.User, error) {
	var users []models.User
	query := `
	SELECT DISTINCT u.id, u.name, u.username, c.id AS class_id, c.name as class_name
    FROM almanax.users u
    JOIN almanax.users_jobs uj ON u.id = uj.user_id
    JOIN almanax.jobs j ON uj.job_id = j.id
	JOIN almanax.classes c ON u.class_id = c.id
	WHERE unaccent(lower(j.name)) LIKE unaccent(lower($1))
	`
	rows, err := db.Query(query, "%"+job+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		var classId int
		var className string

		if err := rows.Scan(
			&u.Id,
			&u.Name,
			&u.Username,
			&classId,
			&className,
		); err != nil {
			return nil, err
		}

		u.Class = models.Class{
			Id:   classId,
			Name: className,
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByUsername(db *sql.DB, name string) (models.User, error) {
	var j models.User
	err := db.QueryRow("SELECT id, name, username FROM almanax.users WHERE username = $1", name).Scan(&j.Id, &j.Name, &j.Username)
	return j, err
}

func GetAllUsers(db *sql.DB) ([]models.User, error) {
	query := `
    SELECT u.id, u.name, u.username, c.id AS class_id, c.name AS class_name
    FROM almanax.users u
    JOIN almanax.classes c ON u.class_id = c.id
    ORDER BY u.name
    `

	var users []models.User
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var u models.User
		var classId int
		var className string
		if err := rows.Scan(
			&u.Id,
			&u.Name,
			&u.Username,
			&classId,
			&className,
		); err != nil {
			return nil, err
		}

		u.Class = models.Class{
			Id:   classId,
			Name: className,
		}

		users = append(users, u)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func GetUsers(db *sql.DB) ([]models.User, error) {
	rows, err := db.Query("SELECT id, name FROM almanax.users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var r models.User
		if err := rows.Scan(&r.Id, &r.Name); err != nil {
			return nil, err
		}
		users = append(users, r)
	}
	return users, nil
}

func GetUserById(db *sql.DB, id int) (models.User, error) {
	var j models.User
	err := db.QueryRow("SELECT id, name FROM almanax.Users WHERE id = $1", id).Scan(&j.Id, &j.Name)

	if err != nil {
		log.Fatal(err)
	}
	return j, nil
}
