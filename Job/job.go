package job

import (
	"strings"

	database "github.com/chamaloown/difus/Database"
	formatter "github.com/chamaloown/difus/Job/Formatter"
	jreader "github.com/chamaloown/difus/Job/Reader"
	ureader "github.com/chamaloown/difus/User/Reader"
)

func GetUsersByJob(message string) (string, error) {
	db := database.GetDBInstance()
	strArr := strings.Split(message, " ")

	if len(strArr) == 2 {
		name := strArr[1]
		users, err := ureader.GetUsersByJob(db, name)

		if err != nil {
			return "", err
		}
		return formatter.ListUsersByJob(users), nil
	} else if len(strArr) == 1 {
		jobMap, err := jreader.GetJobWithAffiliatedUser(db)
		if err != nil {
			return "", err
		}
		return formatter.ListJobsWithAssociatedUser(jobMap), nil
	} else {
		return "Le nombre d'argument n'est pas le bon. Consulter l'aide avec un !help", nil
	}
}
