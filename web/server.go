package web

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/Yousseflamrani/tp2_annuaire/annuaire"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("annuaire").Parse(`
			<!DOCTYPE html>
			<html lang="fr">
			<head>
				<meta charset="UTF-8">
				<title>Annuaire</title>
			</head>
			<body>
				<h1>Annuaire</h1>
				{{if .}}
					{{range .}}
						<p><strong>{{.Nom}}</strong> - {{.Tel}}</p>
					{{end}}
				{{else}}
					<p>Aucun contact enregistré.</p>
				{{end}}
			</body>
			</html>
		`))

		if err := tmpl.Execute(w, annuaire.Liste); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("🌐 Serveur démarré sur http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
