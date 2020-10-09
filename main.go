package main

import (
	"fmt"
	"github.com/kenshaw/imdb"
	"log"
	"os"
	"path/filepath"
)

type LibraryItem struct {
	Imdb *imdb.MovieResult
	MediaInfo *MediaInfo
}

func main() {


	library := make(map[string][]*LibraryItem)
	mfChan := make(chan *MediaInfo)
	mImdbChan := make(chan *LibraryItem)
	go GetMediaInfo(os.Args[1], mfChan)

	go getMediaIMDBID(mfChan,mImdbChan)
	for i := range mImdbChan {
		library[i.Imdb.ImdbID] = append(library[i.Imdb.ImdbID],i)
	}
	for k,l:=range library{
		fmt.Printf("%s:%d",k,len(l))
		if len(l) > 0{
			fmt.Printf("[")
			for _,i:= range l{
				fmt.Printf("%s,",i.MediaInfo.Media.Ref)
			}
			fmt.Printf("]")
		}
		fmt.Printf("\n")

	}
}

func getMediaIMDBID(mfChan chan *MediaInfo, mImdbChan chan *LibraryItem) {
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
			movie = filepath.Base(m.Media.Ref)
			result, err = md.SearchMovie(movie)
			if err != nil{
				log.Println(err.Error())
			}
		}
		if result != nil {
			log.Printf("%s\n", result.ImdbID)
			l := LibraryItem{result,m}
			mImdbChan <- &l
		}

	}
}
