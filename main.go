package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	mediainfo, err := GetMediaInfo(os.Args[1])
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	for _,m:= range mediainfo {
		t, err := m.GetMovieTrackID()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		movie := t.Movie
		if movie == "" {
			movie = filepath.Base(m.Media.Ref)
		}
		md := NewMovieId()
		result, err := md.SearchMovie(movie)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fmt.Printf("%s\n", result.ImdbID)
	}
}
