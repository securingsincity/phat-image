package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/golang/freetype/truetype"
	"github.com/securingsincity/phat-image/imagebuilder"
	"github.com/zmb3/spotify"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/font/gofont/gobolditalic"
)

// redirectURI is the OAuth redirect URI for the application.
// You must register an application at Spotify's developer portal
// and enter this value.

var (
	ch    = make(chan *spotify.Client)
	state = "abc123ab3"
)

func main() {
	const S = 400
	redirectURI := GetString("SPOTIFY_REDIRECT_URL", "http://localhost:8080/callback")
	var auth = spotify.NewAuthenticator(redirectURI, spotify.ScopeUserReadCurrentlyPlaying)
	var writeOutImage = true
	var title = ""
	var artist = ""
	var style = 1
	flag.StringVar(&title, "title", "", "name of the song")
	flag.StringVar(&artist, "artist", "", "name of the song")
	flag.IntVar(&style, "style", 1, "name of the song")
	flag.BoolVar(&writeOutImage, "outto", true, "output images")
	flag.Parse()

	boldFont, err := truetype.Parse(gobold.TTF)
	if err != nil {
		panic("")
	}
	boldFace := truetype.NewFace(boldFont, &truetype.Options{
		Size: 18,
	})

	boldItalicfont, _ := truetype.Parse(gobolditalic.TTF)

	boldItalicFace := truetype.NewFace(boldItalicfont, &truetype.Options{
		Size: 16,
	})

	// first start an HTTP server
	http.HandleFunc("/callback", completeAuth(auth))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
	})
	go http.ListenAndServe(":8080", nil)

	url := auth.AuthURL(state)
	fmt.Println("Please log in to Spotify by visiting the following page in your browser:", url)

	// wait for auth to complete
	client := <-ch

	// use the client to make calls that require authorization
	t := time.Tick(time.Second * 20)

	for range t {
		currentPlaying, err := client.PlayerCurrentlyPlaying()
		if err != nil {
			log.Fatalf("FATAL line 46:%s", err)
		}

		newTitle := currentPlaying.Item.Name
		newArtist := currentPlaying.Item.Artists[0].Name
		if newTitle != title || newArtist != artist {
			title = newTitle
			artist = newArtist
			log.Printf("Currently Playing: %v by %v", currentPlaying.Item.Name, currentPlaying.Item.Artists[0].Name)

			rand.Seed(time.Now().UnixNano())
			style = rand.Intn(4) + 1
			dc := imagebuilder.GenerateImage(style, boldFace, boldItalicFace, title, artist)
			if writeOutImage == true {
				imagebuilder.WriteToEink(dc)
			} else {
				dc.SavePNG("out.png")
			}

		}

	}
}

type Handler func(w http.ResponseWriter, r *http.Request)

func completeAuth(auth spotify.Authenticator) Handler {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Got request for:", r.URL.String())
		tok, err := auth.Token(state, r)
		if err != nil {
			http.Error(w, "Couldn't get token", http.StatusForbidden)
			log.Fatal(err)
		}
		if st := r.FormValue("state"); st != state {
			http.NotFound(w, r)
			log.Fatalf("State mismatch: %s != %s\n", st, state)
		}
		// use the token to get an authenticated client
		client := auth.NewClient(tok)
		fmt.Fprintf(w, "Login Completed!")
		ch <- &client
	}
}

func GetString(key string, defaultVal string) string {
	val := os.Getenv(key)
	if len(val) == 0 {
		return defaultVal
	}

	return val
}
