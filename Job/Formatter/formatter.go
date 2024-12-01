package formatter

import (
	"fmt"
	"sort"
	"strings"

	models "github.com/chamaloown/difus/Models"
)

func ListJobsWithAssociatedUser(jobsMap map[string][]models.User) string {
	var jobList []string

	jobNames := make([]string, 0, len(jobsMap))
	for jobName := range jobsMap {
		jobNames = append(jobNames, jobName)
	}
	sort.Strings(jobNames)

	for _, jobName := range jobNames {
		users := jobsMap[jobName]

		jobEntry := fmt.Sprintf("**%s** :", jobName)

		if len(users) == 0 || (len(users) == 1 && users[0].Name == "Nobody") {
			jobEntry += " N/A"
		} else {
			userNames := make([]string, 0, len(users))
			for _, user := range users {
				if user.Name != "Nobody" {
					userNames = append(userNames, fmt.Sprintf("%s (%s)", user.Name, user.Class.Name))
				}
			}
			jobEntry += " " + strings.Join(userNames, ", ")
		}

		jobList = append(jobList, jobEntry)
	}

	return strings.Join(jobList, "\n")
}

func ListUsersByJob(users []models.User) string {
	var list []string

	for _, user := range users {
		userEntry := fmt.Sprintf("%s (Classe: %s)", user.Name, user.Class.Name)
		list = append(list, userEntry)
	}

	if len(list) == 0 {
		return "Aucun utilisateur trouvé pour ce métier."
	}

	message := fmt.Sprintf("Utilisateurs :\n- %s", strings.Join(list, "\n- "))
	return message
}
