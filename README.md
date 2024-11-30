# Difus

Difus is a Go Application for dealing with Dofus quest and help in your journey.

## Installation

You have to have access to `almanax.csv`, ask for it!

You can run directly [go](https://go.dev/) the application by typing :



```bash
go run main.go
```

Otherwise, you can build the binary by typing :

```bash
go build main.go
```

And then running it :

```bash
./difus
```

## Usage

	📜 **!author** - Affiche le nom de l'auteur.
	❓ **!help** - Affiche ce message d'aide.
	📅 **!alma [today | week | JJ/MM/AAAA]** - Récupère l'Almanax pour un jour spécifique :
	      •  **today** : Affiche l'Almanax d'aujourd'hui.
	      •  **week** : Affiche l'Almanax pour toute la semaine.
	      •  **JJ/MM/AAAA** : Affiche l'Almanax pour une date spécifique (ex. 08/11/2024).
	🗣️ **!ask [question]** - Pose une question technique sur dofus (Attention l'IA a comme pour dernière connaissance la mise a jour 2.62).
	🛠️ **!metier [metier] ?[lvl]** - Récupère tous les utilisateurs farmant ce métier, filtrer par niveau si celui-ci est renseigner.
	
	
	Veuillez utiliser le bon format de date ou les mots-clés spécifiés pour chaque option.

## Contributing

Pull requests are welcome. For major changes, please open an issue first
to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License

[MIT](https://choosealicense.com/licenses/mit/)