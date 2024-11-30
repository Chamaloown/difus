package user

import (
	"database/sql"
	"strings"

	creader "github.com/chamaloown/difus/Class/Reader"
	database "github.com/chamaloown/difus/Database"
	models "github.com/chamaloown/difus/Models"
	ureader "github.com/chamaloown/difus/User/Reader"
	writer "github.com/chamaloown/difus/User/Writer"
)

func AddUSer(message string) (string, error) {
	db := database.GetDBInstance()
	strArr := strings.Split(message, " ")
	if len(strArr) != 4 {
		return "Le nombre d'argument n'est pas le bon. Consulter l'aide avec un !help", nil
	}
	name := strArr[1]
	userName := strArr[2]
	className := strArr[3]

	class, err := creader.GetClassByName(db, className)
	if err != nil {
		return "La classe n'existe pas", nil
	}

	_, err = ureader.GetUserByUsername(db, userName)
	if err == nil {
		return "L'utilisateur existe déja", nil
	}

	if err == sql.ErrNoRows {
		newUser := models.User{
			Name:     name,
			Username: userName,
			Class:    class,
		}
		_, err := writer.CreateUser(db, newUser)
		if err != nil {
			return "Erreur à la création de l'utilisateur", nil
		}
	}

	return "L'utilisateur a bien été ajouté!", nil

}
