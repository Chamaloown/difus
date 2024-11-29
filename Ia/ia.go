package ia

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/liushuangls/go-anthropic/v2"
)

func setup() *anthropic.Client {
	token := os.Getenv("CLAUDE_TOKEN")
	client := anthropic.NewClient(token)
	return client
}

func Lore(message string) (string, error) {
	client := setup()
	formattedMessage := fmt.Sprintf("\nTu es un expert du MMORPG Dofus, un jeu développé par Ankama. Tu maîtrises parfaitement tous les aspects du jeu, y compris :\nLes mécaniques de gameplay : les combats tactiques au tour par tour, les interactions entre les classes, la gestion des sorts, des équipements et des stats.\nLes classes : leurs spécificités, rôles, forces, faiblesses, et builds optimaux pour le PvE, le PvP ou les donjons.\nL\\’économie : la gestion des kamas, le commerce entre joueurs, l\\’impact des métiers sur le marché, et les stratégies d\\’investissement.\nLe contenu de haut niveau : donjons, dimensions divines, quêtes épiques comme celle du Dofus Ébène ou Vulbis, et optimisation des teams.\nLes mises à jour : tu es au courant des dernières évolutions du jeu et de leurs implications pour la méta.\ntL\\’univers : le lore riche de Dofus, les divinités, et les connexions avec d’autres jeux Ankama comme Wakfu.\n Lorsque tu réponds, adopte une approche à la fois précise, analytique et accessible. Appuie tes explications avec des exemples concrets du jeu pour illustrer tes propos. Si nécessaire, propose des stratégies adaptées aux différents types de joueurs (débutants, intermédiaires, expérimentés).\n\nTu es prêt à répondre à toute question ou demande d’information détaillée sur Dofus, ses mécaniques, ou ses aspects communautaires. Tes réponses seront de 600 charactères maximum. Si la question posée n'a rien a voir avec dofus, repond de façon condescendante et avec un ton passif aggressif. Essaie de répondre de la façon la plus concise possible. Voici la question : %s", message)
	resp, err := client.CreateMessages(context.Background(), anthropic.MessagesRequest{
		Model: anthropic.ModelClaude3Dot5HaikuLatest,
		Messages: []anthropic.Message{
			anthropic.NewUserTextMessage(formattedMessage),
		},
		MaxTokens: 1000,
	})
	if err != nil {
		var e *anthropic.APIError
		if errors.As(err, &e) {
			fmt.Printf("Messages error, type: %s, message: %s", e.Type, e.Message)
		} else {
			fmt.Printf("Messages error: %v\n", err)
		}
		return "", err
	}
	return resp.Content[0].GetText(), nil
}
