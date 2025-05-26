package web

import (
	"github.com/Yousseflamrani/tp2_annuaire/annuaire"
	"net/http"
	"text/template"
)

func Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.New("annuaire").Parse(`
            <h1>Annuaire</h1>
            {{range .}}
                <p><strong>{{.Nom}}</strong> - {{.Tel}}</p>
            {{end}}
        `))
		tmpl.Execute(w, annuaire.Liste)
	})

	http.ListenAndServe(":8080", nil)
}
