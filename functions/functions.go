package functions

import (
	structure "TpSpotify/struct"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var ListeAlbums []structure.AlbumData

// Variables Token
var AccessToken string
var TokenType string

// Variables Laylow
var ArtistName string
var Title string
var Album string
var AlbumCover string
var ReleaseDate string

// Variables Damso
var NomAlbumDamso string
var NbMusiqueAlbumDamso int
var DateSortieAlbumDamso string

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
	var decodeData structure.ApiData

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)

	AccessToken = decodeData.AccessToken
	TokenType = decodeData.TokenType
}

func GetDamsosAlbums() {

	type ApiData struct {
		Items []struct {
			Images []struct {
				Url string `json:"url"`
			} `json:"images"`
			ReleaseDate string `json:"release_date"`
			Name        string `json:"name"`
			TotalTracks int    `json:"total_tracks"`
		} `json:"items"`
	}

	// URL de L'API
	urlApi := "https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums?include_groups=album"

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
	req.Header.Add("Authorization", TokenType+" "+AccessToken)

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

	ListeAlbums = []structure.AlbumData{}

	// Decodage des données en format JSON et ajout des donnée à la variable: decodeData
	json.Unmarshal(body, &decodeData)
	for i := 0; i <= len(decodeData.Items)-1; i++ {
		var albumInfos structure.AlbumData
		// Affichage des données
		albumInfos.Name = decodeData.Items[i].Name
		albumInfos.TotalTracks = decodeData.Items[i].TotalTracks
		albumInfos.ReleaseDate = decodeData.Items[i].ReleaseDate

		if len(decodeData.Items[i].Images) > 0 {
			albumInfos.ImageUrl = decodeData.Items[i].Images[0].Url
		}
		ListeAlbums = append(ListeAlbums, albumInfos)
	}
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
	req.Header.Add("Authorization", TokenType+" "+AccessToken)

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

	ArtistName = decodeData.Album.Artists[0].Name
	Title = decodeData.Name
	Album = decodeData.Album.Name
	AlbumCover = decodeData.Album.Images[0].Url
	ReleaseDate = decodeData.Album.ReleaseDate

}
