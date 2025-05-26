package annuaire

type Contact struct {
	Nom string
	Tel string
}

var Liste = make(map[string]Contact)

func Ajouter(nom, tel string) string {
	if _, existe := Liste[nom]; existe {
		return "Ce nom existe déjà."
	}
	Liste[nom] = Contact{Nom: nom, Tel: tel}
	return "Contact ajouté."
}

func Rechercher(nom string) (Contact, bool) {
	contact, existe := Liste[nom]
	return contact, existe
}

func Supprimer(nom string) string {
	if _, existe := Liste[nom]; existe {
		delete(Liste, nom)
		return "Contact supprimé."
	}
	return "Contact introuvable."
}

func Modifier(nom, tel string) string {
	if _, existe := Liste[nom]; existe {
		Liste[nom] = Contact{Nom: nom, Tel: tel}
		return "Contact modifié."
	}
	return "Contact introuvable."
}
