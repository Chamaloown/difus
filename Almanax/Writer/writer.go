package writer

import (
	"database/sql"

	models "github.com/chamaloown/difus/Models"
)

func CreateAlmanax(db *sql.DB, almanax models.Almanax) (int, error) {
	var id int
	err := db.QueryRow("INSERT INTO almanax.almanaxes(date, merydes, type, bonus, offerings, quantity_offered, kamas) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id",
		almanax.Date, almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
