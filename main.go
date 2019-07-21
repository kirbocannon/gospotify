package main

import (
	"context"
	"fmt"
	"github.com/theillego/gospotify/internal/logger"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
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

// initialize map
func NewStringSet() *StringSet {
	return &StringSet{make(map[string]bool)}
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

func FormatSpotifyErrorMessage(err error) (parsedError []string) {
	parsedError = strings.Split(err.Error(), "Response:")

	// strip any new lines found in messages
	// This is currently NOT working
	for i := 0; i < 2; i++ {
		parsedError = append(parsedError, strings.TrimSuffix(parsedError[i], "\n"))
	}

	return
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
		serr := FormatSpotifyErrorMessage(err)[0]
		logto.SpotifyError(serr)

	} else {
		validToken = true
	}

	if validToken {

		client := spotify.Authenticator{}.NewClient(token)

		// get playlist
		userPlaylists, err := client.GetPlaylistsForUser(userID)
		if err != nil {
			spotifyError := FormatSpotifyErrorMessage(err)[0]
			logto.SpotifyError(spotifyError)
		}

		artists := NewStringSet()
		tracks := map[string][]string{}
		//genres := map[string][]string{}

		for _, p := range userPlaylists.Playlists {
			//fmt.Println(p.Name, p.Tracks.Total)
			playListTracks, err := client.GetPlaylistTracks(p.ID)
			if err != nil {
				spotifyError := FormatSpotifyErrorMessage(err)[0]
				logto.SpotifyError(spotifyError)
			}
			for _, t := range playListTracks.Tracks {
				fmt.Println(t)
				//genres[t.Track.]
				fmt.Println(client.Get())

				for _, a := range t.Track.Artists {
					artists.Add(a.Name + "|" + a.ID.String())
					tracks[a.ID.String()] = append(tracks[a.ID.String()], t.Track.Name)
				}
			}

			break
		}

		//for key, val := range tracks {
		//	if len(val) > 4 {
		//		fmt.Println(len(val), key, val)
		//	}
		//
		//}
	}
}