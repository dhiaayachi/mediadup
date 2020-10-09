package main

import (
	"fmt"
	"github.com/kenshaw/imdb"
	"log"
	"os"
	"path/filepath"
)

func main() {
	mfChan := make(chan *MediaInfo)
	mImdbChan := make(chan *imdb.MovieResult)
	go GetMediaInfo(os.Args[1], mfChan)

	go getMediaIMDBID(mfChan,mImdbChan)
	for i := range mImdbChan {
		fmt.Printf("%s\n", i.ImdbID)
	}
}

func getMediaIMDBID(mfChan chan *MediaInfo,mImdbChan chan *imdb.MovieResult) {
	defer close(mImdbChan)
	for m := range mfChan {
		t, err := m.GetMovieTrackID()
		if err != nil {
			log.Println(err.Error())
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
		}
		if result != nil {
			mImdbChan <- result
		}

	}
}
