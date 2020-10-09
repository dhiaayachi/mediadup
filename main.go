package main

import (
	"fmt"
	"github.com/kenshaw/imdb"
	"log"
	"os"
	"path/filepath"
)

func main() {


	library := make(map[string][]*imdb.MovieResult)
	mfChan := make(chan *MediaInfo)
	mImdbChan := make(chan *imdb.MovieResult)
	go GetMediaInfo(os.Args[1], mfChan)

	go getMediaIMDBID(mfChan,mImdbChan)
	for i := range mImdbChan {
		library[i.ImdbID] = append(library[i.ImdbID],i)
	}
	for k,l:=range library{
		fmt.Printf("%s:%d\n",k,len(l))
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
			log.Printf("%s\n", result.ImdbID)
			mImdbChan <- result
		}

	}
}
