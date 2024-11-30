package writer

import (
	"database/sql"

	models "github.com/chamaloown/difus/Models"
)

func CreateUser(db *sql.DB, user models.User) (int, error) {
	var id int
	query := `
		INSERT INTO almanax.users (name, username, class_id)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	err := db.QueryRow(query, user.Name, user.Username, user.Class.Id).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil

}
