package reader

import (
	"database/sql"

	models "github.com/chamaloown/difus/Models"
)

func GetJobWithAffiliatedUser(db *sql.DB) (map[string][]models.User, error) {
	query := `
		SELECT 
			j.name AS job_name, 
			COALESCE(u.id, 0) AS user_id, 
			COALESCE(u.name, 'Nobody') AS user_name, 
			COALESCE(u.username, '') AS user_username, 
			COALESCE(c.id, 0) AS class_id, 
			COALESCE(c.name, 'N/A') AS class_name
		FROM almanax.jobs j
		LEFT JOIN 
			almanax.users_jobs uj ON j.id = uj.job_id
		LEFT JOIN 
			almanax.users u ON uj.user_id = u.id
		LEFT JOIN 
			almanax.classes c ON u.class_id = c.id
		ORDER BY 
			j.name, u.name
		`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobUsers := make(map[string][]models.User)

	for rows.Next() {
		var jobName string
		var userId sql.NullInt64
		var userName, userUsername sql.NullString
		var classId sql.NullInt64
		var className sql.NullString

		err := rows.Scan(
			&jobName,
			&userId,
			&userName,
			&userUsername,
			&classId,
			&className,
		)
		if err != nil {
			return nil, err
		}

		if userId.Valid {
			user := models.User{
				Id:       int(userId.Int64),
				Name:     userName.String,
				Username: userUsername.String,
			}

			if classId.Valid {
				user.Class = models.Class{
					Id:   int(classId.Int64),
					Name: className.String,
				}
			}

			jobUsers[jobName] = append(jobUsers[jobName], user)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return jobUsers, nil
}

func GetJobByName(db *sql.DB, name string) (models.Job, error) {
	var j models.Job
	err := db.QueryRow("SELECT id, name, type FROM almanax.jobs WHERE unaccent(upper(name)) LIKE unaccent(upper($1))", "%"+name+"%").Scan(&j.Id, &j.Name, &j.Type)

	if err != nil {
		return models.Job{}, err
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
		return models.Job{}, err
	}
	return j, nil
}
