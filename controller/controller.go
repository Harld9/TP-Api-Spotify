package controller

import (
	"html/template"
	"net/http"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

type PageData struct {
	Title   string
	Message string
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/damso.html"))
	tmpl.Execute(w, data)
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title:   "Accueil",
		Message: "Bienvenue Au Jeu Puissance 4 🎉",
	}
	tmpl := template.Must(template.ParseFiles("template/laylow.html"))
	tmpl.Execute(w, data)
}
