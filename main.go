package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"gospotify/internal/dataUtils"
	"gospotify/internal/logger"
	"strings"
)

type Track struct {
	Name string
	Artists []spotify.SimpleArtist

}

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

func main() {
	// setup logging
	logto := logger.NewLogger()
	userID := "theillego"
	validToken := false

	// provide API credentials
	config := &clientcredentials.Config{
		ClientID:     "75c5079571ec441dbc4878efd335363f",
		ClientSecret: "0c05eaf460ae43c9a7c611b57bec4737",
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())

	if err != nil {
		serr := logger.FormatSpotifyErrorMessage(err)[0]
		logto.SpotifyError(serr)

	} else {
		validToken = true
	}

	if validToken {

		client := spotify.Authenticator{}.NewClient(token)

		// get playlist
		userPlaylists, err := client.GetPlaylistsForUser(userID)
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
		genreCounts := make(map[string]int)

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
			if value > 19 {
				fmt.Println(key, value)
			}

		}






		//for key, val := range tracks {
		//	if len(val) > 4 {
		//		fmt.Println(len(val), key, val)
		//	}
		//
		//}
	}
}