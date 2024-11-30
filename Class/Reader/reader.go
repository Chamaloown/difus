package reader

import (
	"database/sql"

	models "github.com/chamaloown/difus/Models"
)

func GetClassByName(db *sql.DB, name string) (models.Class, error) {
	var class models.Class
	query := `
		SELECT id, name FROM almanax.classes WHERE name = $1
	`
	err := db.QueryRow(query, name).Scan(&class.Id, &class.Name)
	if err != nil {
		return models.Class{}, err
	}
	return class, nil
}
