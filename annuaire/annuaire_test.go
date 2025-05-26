package annuaire

import "testing"

func TestAjouter(t *testing.T) {
	Liste = make(map[string]Contact)
	msg := Ajouter("Youssef", "0102030405")
	if msg != "Contact ajouté." {
		t.Errorf("Échec ajout: %s", msg)
	}

	msg2 := Ajouter("Youssef", "0102030405")
	if msg2 != "Ce nom existe déjà." {
		t.Errorf("Échec détection doublon: %s", msg2)
	}
}

func TestRechercher(t *testing.T) {
	Liste = make(map[string]Contact)
	Ajouter("Bob", "0607080910")

	contact, ok := Rechercher("Bob")
	if !ok || contact.Tel != "0607080910" {
		t.Error("Echec recherche")
	}
}
