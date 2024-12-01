package user

import (
	"database/sql"
	"fmt"
	"strings"

	creader "github.com/chamaloown/difus/Class/Reader"
	database "github.com/chamaloown/difus/Database"
	jreader "github.com/chamaloown/difus/Job/Reader"
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

func AddUserJob(message string) (string, error) {
	db := database.GetDBInstance()
	strArr := strings.Split(message, " ")
	if len(strArr) != 3 {
		return "Le nombre d'argument n'est pas le bon. Consulter l'aide avec un !help", nil
	}
	userName := strArr[1]
	jobName := strArr[2]

	job, err := jreader.GetJobByName(db, jobName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Métier introuvable, veuillez consulter l'aide", nil
		} else {
			return "", err
		}
	}

	user, err := ureader.GetUserByUsername(db, userName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "Utilisateur introuvable, veuillez consulter l'aide", nil
		} else {
			return "", err
		}
	}

	_, err = writer.LinkUserToJob(db, user.Id, job.Id)
	if err != nil {
		return "", err
	}
	return "L'utilisateur " + user.Username + " a correctement été lié au métier " + job.Name, nil
}

func GetUsers() (string, error) {
	db := database.GetDBInstance()
	users, err := ureader.GetAllUsers(db)
	if err != nil {
		return "", err
	}
	var msg string
	msg = "Liste des Utilisateurs : "
	for _, user := range users {
		msg += fmt.Sprintf("\n - %s -> %s", user.Username, user.Class.Name)
	}
	return msg, nil
}

func DeleteUser(message string) (string, error) {
	db := database.GetDBInstance()
	strArr := strings.Split(message, " ")
	if len(strArr) > 2 {
		return "Le nombre d'argument n'est pas le bon. Consulter l'aide avec un !help", nil
	}
	username := strArr[1]
	user, err := ureader.GetUserByUsername(db, username)
	if err != nil {
		return "", err
	}
	_, err = writer.DeleteUser(db, user)
	if err != nil {
		return "", err
	}
	return "L'utilisateur " + user.Username + " a bien été supprimé", nil
}
