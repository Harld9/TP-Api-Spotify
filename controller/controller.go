package controller

import (
	"TpSpotify/functions"
	"html/template"
	"net/http"
)

// renderTemplate est une fonction utilitaire pour afficher un template HTML avec des données dynamiques
func renderTemplate(w http.ResponseWriter, filename string, data map[string]string) {
	tmpl := template.Must(template.ParseFiles("template/" + filename)) // Charge le fichier template depuis le dossier "template"
	tmpl.Execute(w, data)                                              // Exécute le template et écrit le résultat dans la réponse HTTP
}

type PageData struct {
	ArtistName  string
	Title       string
	Album       string
	AlbumCover  string
	ReleaseDate string
}

func Home(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, data)
}

func Damso(w http.ResponseWriter, r *http.Request) {
	data := PageData{}
	tmpl := template.Must(template.ParseFiles("template/damso.html"))
	tmpl.Execute(w, data)
}

func Laylow(w http.ResponseWriter, r *http.Request) {
	functions.GetLaylowsTrack()
	data := PageData{
		ArtistName:  functions.ArtistName,
		Title:       functions.Title,
		Album:       functions.Album,
		AlbumCover:  functions.AlbumCover,
		ReleaseDate: functions.ReleaseDate,
	}
	tmpl := template.Must(template.ParseFiles("template/laylow.html"))
	tmpl.Execute(w, data)
}
