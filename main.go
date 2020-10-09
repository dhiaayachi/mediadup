package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	mfChan := make(chan *MediaInfo)
	go GetMediaInfo(os.Args[1], mfChan)

	for m:= range mfChan {
		t, err := m.GetMovieTrackID()
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		movie := t.Movie
		if !m.IsMedia() {
			continue
		}
		if movie == "" {
			movie = filepath.Base(m.Media.Ref)
		}
		md := NewMovieId()
		result, err := md.SearchMovie(movie)
		if err != nil {
			log.Println(err.Error())
			os.Exit(1)
		}
		if result !=nil {
			fmt.Printf("%s\n", result.ImdbID)
		}
	}
}
