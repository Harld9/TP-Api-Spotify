package functions

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type ApiData struct {
	TokenType   string `json:"token_type"`
	AccessToken string `json:"access_token"`
}

func RefreshToken() {
	// URL de L'API
	urlApi := "https://accounts.spotify.com/api/token"
	urlEnd := "client_id=c235f728d69640e5b067e81dc4af6c87&client_secret=38dc2a7cc1b642bdaf81f6fdbcf2ebca&grant_type=client_credentials"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodPost, urlApi, strings.NewReader(urlEnd))
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
		return
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	fmt.Println(decodeData)
}

func GetDamsosAlbums() {

	type ApiData struct {
		Items []struct {
			Images []struct {
				Url string `json:"url"`
			} `json:"images"`
			Release_Date string `json:"release_date"`
			Name         string `json:"name"`
			Total_Tracks int    `json:"total_tracks"`
		} `json:"items"`
	}

	// URL de L'API
	urlApi := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
		return
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("Authorization", "Bearer BQC6bPxdRHfFYit3lpWBI-E9OVXy64nHXCe-Hk6Ru-zgfFheF1P4TJmeUwvdPkcVLJ-cNogjCIB1MNhiJOmxNQjecFUs3zB1oKTA4ZwIvz8Z_Ptob0SOo2NLbxBavBT-u0NCg12tCS0")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	// Affichage des données
	fmt.Println("Nom de l'album :", decodeData.Items[0].Name)
	fmt.Println("Nombre de musique dans l'album :", decodeData.Items[0].Total_Tracks)
	fmt.Println("Date de sortie de l'album :", decodeData.Items[0].Release_Date)
	fmt.Println("Lien de la cover de l'album :", decodeData.Items[0].Images[0])
}

func GetLaylowsTrack() {

	type ApiData struct {
		Name  string `json:"name"`
		Album struct {
			Artists []struct {
				Name string `json:"name"`
			} `json:"artists"`
			Name        string `json:"name"`
			ReleaseDate string `json:"release_date"`
			Images      []struct {
				Url string `json:"url"`
			} `json:"images"`
		} `json:"album"`
	}

	// URL de L'API
	urlApi := "https://api.spotify.com/v1/tracks/67Pf31pl0PfjBfUmvYNDCL"

	// Initialisation du client HTTP qui va émettre/demander les requêtes
	httpClient := http.Client{
		Timeout: time.Second * 2, // Timeout apres 2sec
	}

	// Création de la requête HTTP vers L'API avec initialisation de la methode HTTP, la route et le corps de la requête
	req, errReq := http.NewRequest(http.MethodGet, urlApi, nil)
	if errReq != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errReq.Error())
		return
	}

	// Ajout d'une métadonnée dans le header, User_Agent permet d'identifier l'application, système ....
	req.Header.Add("Authorization", "Bearer BQC6bPxdRHfFYit3lpWBI-E9OVXy64nHXCe-Hk6Ru-zgfFheF1P4TJmeUwvdPkcVLJ-cNogjCIB1MNhiJOmxNQjecFUs3zB1oKTA4ZwIvz8Z_Ptob0SOo2NLbxBavBT-u0NCg12tCS0")

	// Execution de la requête HTTP vars L'API
	res, errResp := httpClient.Do(req)
	if errResp != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errResp.Error())
		return
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	// Lecture et récupération du corps de la requête HTTP
	body, errBody := io.ReadAll(res.Body)
	if errBody != nil {
		fmt.Println("Oupss, une erreur est survenue : ", errBody.Error())
	}

	// Déclaration de la variable qui va contenir les données
	var decodeData ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	//Affichage des données
	fmt.Println("Nom de l'artiste :", decodeData.Album.Artists[0].Name)
	fmt.Println("Titre :", decodeData.Name)
	fmt.Println("Album :", decodeData.Album.Name)
	fmt.Println("Lien de la cover de l'album :", decodeData.Album.Images[0].Url)
	fmt.Println("Date de Sortie :", decodeData.Album.ReleaseDate)

}
