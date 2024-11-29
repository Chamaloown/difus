package formatter

import (
	"fmt"
	"strings"

	models "github.com/chamaloown/difus/Models"
)

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
