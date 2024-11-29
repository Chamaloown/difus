package job

import (
	"fmt"

	database "github.com/chamaloown/difus/Database"
	formatter "github.com/chamaloown/difus/Job/Formatter"
	parser "github.com/chamaloown/difus/Job/Parser"
	uReader "github.com/chamaloown/difus/User/Reader"
)

func GetUsersByJob(message string) (string, error) {
	args, err := parser.Parse(message)
	if err != nil {
		return "", err
	}
	name := args[1]
	db := database.GetDBInstance()
	users, err := uReader.GetUsersByJob(db, name)

	if err != nil {
		fmt.Println(err)
	}

	return formatter.ListUsersByJob(users), nil
}
