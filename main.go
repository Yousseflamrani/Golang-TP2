package main

import (
	"flag"
	"fmt"

	"github.com/Yousseflamrani/tp2_annuaire/annuaire"
	"github.com/Yousseflamrani/tp2_annuaire/web"
)

func main() {
	action := flag.String("action", "", "Action à effectuer (ajouter, rechercher, supprimer, modifier, lister)")
	nom := flag.String("nom", "", "Nom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone")
	webFlag := flag.Bool("web", false, "Lancer le serveur web")
	flag.Parse()

	if *webFlag {
		web.Start()
		return
	}

	switch *action {
	case "ajouter":
		fmt.Println(annuaire.Ajouter(*nom, *tel))
	case "rechercher":
		if contact, ok := annuaire.Rechercher(*nom); ok {
			fmt.Printf("Nom: %s, Téléphone: %s\n", contact.Nom, contact.Tel)
		} else {
			fmt.Println("Contact introuvable.")
		}
	case "supprimer":
		fmt.Println(annuaire.Supprimer(*nom))
	case "modifier":
		fmt.Println(annuaire.Modifier(*nom, *tel))
	case "lister":
		for _, c := range annuaire.Liste {
			fmt.Printf("Nom: %s, Téléphone: %s\n", c.Nom, c.Tel)
		}
	default:
		fmt.Println("Action inconnue ou manquante.")
	}
}
