package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
	"gopkg.in/yaml.v2"
	"gospotify/internal/dataUtils"
	"gospotify/internal/logger"
	"io/ioutil"
	"path/filepath"
	"strings"
)

// setup logging
var logto = logger.NewLogger()

type Track struct {
	Name string
	Artists []spotify.SimpleArtist

}

type SpotifyClient struct {
	Client *spotify.Client
	SpotifyCredentials *SpotifyCredentials
	UserID string
	Token *oauth2.Token
	ValidToken bool

}

type SpotifyCredentials struct {
	UserId string `yaml:"username"`
	ClientId string `yaml:"id"`
	ClientSecret string `yaml:"secret"`
}

//type SpotifyUsers struct {
//	UserIds *SpotifyUser `yaml:"userIds"`
//
//}
//
//type SpotifyUser struct {
//	ClientId string `yaml:"id"`
//	ClientSecret string `yaml:"secret"`
//}

type Genre struct {
	Name string
	Tracks []Track
}

// set data structure
type StringSet struct {
	set map[string]bool
}

// Artist ID Set
type ArtistIdSet struct {
	set map[spotify.ID]bool
}

// initialize map
func NewArtistIdSet() *ArtistIdSet {
	return &ArtistIdSet{make(map[spotify.ID]bool)}
}

// add if value is not true (key not in map)
func (set *StringSet) Add(i string) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found
}

func (set *StringSet) Get(i string) bool {
	_, found := set.set[i]
	return found	//true if it existed already
}

func (set *StringSet) Remove(i string) {
	delete(set.set, i)
}

func (set *ArtistIdSet) Add(i spotify.ID) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found
}

func (set *ArtistIdSet) Get(i spotify.ID) bool {
	_, found := set.set[i]
	return found	//true if it existed already
}

func (set *ArtistIdSet) Remove(i spotify.ID) {
	delete(set.set, i)
}

func (c *SpotifyClient) Init() {

	var spotifyCreds SpotifyCredentials

	// read app config file
	configFile, err := ReadAppConfigFile()

	if err != nil {
		logto.AppConfigFileError(err.Error())
	}

	err = yaml.Unmarshal(configFile, &spotifyCreds)

	if err != nil {
		panic(err)
	}

	//c.UserID = spotifyCreds.UserId
	c.UserID = spotifyCreds.UserId
	//validToken := false

	// provide API credentials
	config := &clientcredentials.Config{
		ClientID:     spotifyCreds.ClientId,
		ClientSecret: spotifyCreds.ClientSecret,
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	c.Token = token

	if err != nil {
		serr := logger.FormatSpotifyErrorMessage(err)[0]
		logto.SpotifyError(serr)

	} else {
		c.ValidToken = true
	}
}

func ReadAppConfigFile() (configFile []uint8, err error) {
	filename, _ := filepath.Abs("appConfig.yaml")
	configFile, err = ioutil.ReadFile(filename)
	return
}

//func test(client spotify.Client, chunkedArtistIds []string) {
//
//	cids := strings.Join(chunkedArtistIds, ",")
//
//	artistDetails, err := client.GetArtists(spotify.ID(cids))
//
//	if err != nil {
//		spotifyError := logger.FormatSpotifyErrorMessage(err)[0]
//		logto.SpotifyError(spotifyError)
//	}
//
//	for _, ad := range artistDetails {
//		//fmt.Println(ad.Name, ad.Genres, ad.Popularity)
//		for _, g := range ad.Genres {
//			genreCounts[strings.TrimSpace(g)] += 1
//		}
//
//	}
//
//
//}

//func handler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello World!")
//}

func main() {

	//// start http server
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	// read app config file
	//configFile, err := ReadAppConfigFile()
	//
	//if err != nil {
	//	logto.AppConfigFileError(err.Error())
	//}
	//
	//var spotifyCreds SpotifyCredentials
	//err = yaml.Unmarshal(configFile, &spotifyCreds)
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//userID := spotifyCreds.UserId
	//validToken := false
	//
	//// provide API credentials
	//config := &clientcredentials.Config{
	//	ClientID:     spotifyCreds.ClientId,
	//	ClientSecret: spotifyCreds.ClientSecret,
	//	TokenURL:     spotify.TokenURL,
	//}
	//
	//token, err := config.Token(context.Background())
	//
	//if err != nil {
	//	serr := logger.FormatSpotifyErrorMessage(err)[0]
	//	logto.SpotifyError(serr)
	//
	//} else {
	//	validToken = true
	//}

	// initialize spotify client
	c := SpotifyClient{}
	c.Init()
	
	if c.ValidToken {

		client := spotify.Authenticator{}.NewClient(c.Token)

		// get playlist
		userPlaylists, err := client.GetPlaylistsForUser(c.UserID)
		if err != nil {
			spotifyError := logger.FormatSpotifyErrorMessage(err)[0]
			logto.SpotifyError(spotifyError)
		}

		artists := NewArtistIdSet()
		tracks := map[string][]string{}
		//genres := map[string][]string{}

		for _, p := range userPlaylists.Playlists {
			//fmt.Println(p.Name, p.Tracks.Total)
			playListTracks, err := client.GetPlaylistTracks(p.ID)
			if err != nil {
				spotifyError := logger.FormatSpotifyErrorMessage(err)[0]
				logto.SpotifyError(spotifyError)
			}

			for _, t := range playListTracks.Tracks {
				//fmt.Println(t)
				//genres[t.Track.]
				//fmt.Println(client.G())

				for _, a := range t.Track.Artists {
					//fmt.Println(client.GetArtist(a.ID))
					//artists.Add(a.Name + "|" + a.ID.String())
					artists.Add(a.ID)
					tracks[a.ID.String()] = append(tracks[a.ID.String()], t.Track.Name)
				}
			}

			//break
		}

		// string buffer for performance
		buf := bytes.Buffer{}
		//var commaSepArtistIdsList []string
		idx := 0
		//idxa := 0

		// add all artist IDs in one request
		for key := range artists.set {
			if idx != 0 {
				buf.WriteString("," + key.String())
			} else {
				buf.WriteString(key.String())
			}

			idx += 1
		}

		// get artists details
		commaSepArtistIds := buf.String()
		artistIds := strings.Split(commaSepArtistIds, ",")
		chunkedArtistIds := dataUtils.GetChunksFromStringArray(artistIds, 50)

		// keep track of genre counts
		var genreCounts = make(map[string]int)

		for _, ids := range chunkedArtistIds {

			cids := strings.Join(ids, ",")
			artistDetails, err := client.GetArtists(spotify.ID(cids))

			if err != nil {
				spotifyError := logger.FormatSpotifyErrorMessage(err)[0]
				logto.SpotifyError(spotifyError)
			}

			for _, ad := range artistDetails {
				//fmt.Println(ad.Name, ad.Genres, ad.Popularity)
				for _, g := range ad.Genres {
					genreCounts[strings.TrimSpace(g)] += 1
				}

			}

		}

		for key, value := range genreCounts {
			if value > 100 {
				fmt.Println(key, value)
			}

		}

	}
}