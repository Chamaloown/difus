package formatter

import (
	"fmt"

	models "github.com/chamaloown/difus/Models"
)

func FormatAlmanax(almanax models.Almanax) string {
	return fmt.Sprintf(
		"Salut les Dofusiens !\n\n📅 Almanax du **%s**\n\n🔮 **Méryde** : %s\n📈 **Type de Bonus** : %s\n🎁 **Bonus** : %s\n🎒 **Offrande** : %s x%d\n💰 **Prix estimé** : %d kamas\n",
		almanax.Date.Format("2006/01/02"), almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas)
}

func FormatWeeklyAlmanax(almanaxes []models.Almanax) string {
	var result string

	for _, almanax := range almanaxes {
		result += fmt.Sprintf(
			"📅 Almanax du **%s**\n🔮 **Méryde** : %s\n📈 **Type de Bonus** : %s\n🎁 **Bonus** : %s\n🎒 **Offrande** : %s x%d\n💰 **Prix estimé** : %d kamas\n\n",
			almanax.Date.Format("2006/01/02"), almanax.Merydes, almanax.Type, almanax.Bonus, almanax.Offerings, almanax.QuantityOffered, almanax.Kamas,
		)
	}
	return result
}
