# TP2 - Annuaire (en Go)

Ce projet est un **annuaire en ligne de commande** écrit en Go, permettant d’ajouter, modifier, supprimer, rechercher et lister des contacts. Il propose également un affichage via un **serveur web minimaliste en HTML**.

---

##  Fonctionnalités

-  Ajouter un contact (`nom`, `téléphone`)
-  Rechercher un contact par nom
-  Supprimer un contact
- ️ Modifier un contact
-  Lister tous les contacts
-  Afficher l'annuaire dans un navigateur Web (via `net/http`)
-  Tests unitaires inclus

---

##  Lancer le projet

### 1.  Prérequis

- Go 1.18 ou supérieur (ce projet utilise Go 1.24.3)

### 2.  Cloner le repo

```bash
git clone https://github.com/Yousseflamrani/Golang-TP2.git
cd tp2_annuaire
```

### 3.  Commandes disponibles

#### Ajouter un contact

```bash
go run main.go --action ajouter --nom "Youssef" --tel "0600000000"
```

#### Rechercher un contact

```bash
go run main.go --action rechercher --nom "Youssef"
```

#### Modifier un contact

```bash
go run main.go --action modifier --nom "Youssef" --tel "0600000000"
```

####  Supprimer un contact

```bash
go run main.go --action supprimer --nom "Alice"
```

####  Lister tous les contacts

```bash
go run main.go --action lister
```

####  Lancer l’interface Web

```bash
go run main.go --web
```

Puis, ouvre [http://localhost:8080](http://localhost:8080) dans ton navigateur.

---

##  Exécuter les tests unitaires

```bash
go test ./annuaire
```

---

##  Structure du projet

```
tp2_annuaire/
├── annuaire/
│   ├── annuaire.go      
│   └── annuaire_test.go   
├── web/
│   └── server.go          
├── main.go                
├── go.mod               
└── README.md             
```

---

##  Membres du projet

- **Youssef Alaoui El Mrani**

---

##  Bonus possible (non implémenté Pour le moment)

- Export / import des contacts en **JSON**
- Persistance en base de données ou fichier `.json`
- Interface web plus poussée

---

> Ce projet mini projet en Go a été réalisé dans le cadre du TP2 de programmation Go – Efrei

