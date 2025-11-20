package main

import (
	"TpSpotify/functions"
	"TpSpotify/router"
	"fmt"
	"net/http"
)

func main() {

	mux := router.New()
	functions.RefreshToken()
	functions.GetLaylowsTrack()
	fmt.Println("🚀 Token généré pour 1h :", functions.AccessToken)
	fmt.Println("🚀 Serveur démarré sur http://localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Erreur serveur :", err)
	}
}
