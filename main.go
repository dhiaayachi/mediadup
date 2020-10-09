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

	t,err := mediainfo.GetMovieTrackID()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	movie := t.Movie
	if movie == "" {
		movie = filepath.Base(os.Args[1])
	}
	md := NewMovieId()
	result, err := md.SearchMovie(movie)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%v", result)
}
